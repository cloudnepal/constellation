package kubernetes

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/edgelesssys/constellation/coordinator/kubernetes/k8sapi"
	"github.com/edgelesssys/constellation/coordinator/kubernetes/k8sapi/resources"
	"github.com/edgelesssys/constellation/coordinator/role"
	"github.com/spf13/afero"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3"
)

// k8s pod network cidr. This value was chosen to match the default flannel pod network.
const (
	podNetworkCidr = "10.244.0.0/16"
	serviceCidr    = "10.245.0.0/24"
)

// configReader provides kubeconfig as []byte.
type configReader interface {
	ReadKubeconfig() ([]byte, error)
}

// configurationProvider provides kubeadm init and join configuration.
type configurationProvider interface {
	InitConfiguration(externalCloudProvider bool) k8sapi.KubeadmInitYAML
	JoinConfiguration(externalCloudProvider bool) k8sapi.KubeadmJoinYAML
}

// KubeWrapper implements ClusterWrapper interface.
type KubeWrapper struct {
	clusterUtil      k8sapi.ClusterUtil
	configProvider   configurationProvider
	client           k8sapi.Client
	kubeconfigReader configReader
}

// New creates a new KubeWrapper with real values.
func New(clusterUtil k8sapi.ClusterUtil, configProvider configurationProvider, client k8sapi.Client) *KubeWrapper {
	return &KubeWrapper{
		clusterUtil:      clusterUtil,
		configProvider:   configProvider,
		client:           client,
		kubeconfigReader: &KubeconfigReader{fs: afero.Afero{Fs: afero.NewOsFs()}},
	}
}

// InitCluster initializes a new Kubernetes cluster and applies pod network provider.
func (k *KubeWrapper) InitCluster(in InitClusterInput) error {
	// TODO: k8s version should be user input
	if err := k.clusterUtil.InstallComponents(context.TODO(), "1.23.6"); err != nil {
		return err
	}

	initConfig := k.configProvider.InitConfiguration(in.SupportsCloudControllerManager)
	initConfig.SetApiServerAdvertiseAddress(in.APIServerAdvertiseIP)
	initConfig.SetNodeIP(in.NodeIP)
	initConfig.SetNodeName(in.NodeName)
	initConfig.SetPodNetworkCIDR(podNetworkCidr)
	initConfig.SetServiceCIDR(serviceCidr)
	initConfig.SetProviderID(in.ProviderID)
	initConfigYAML, err := initConfig.Marshal()
	if err != nil {
		return fmt.Errorf("encoding kubeadm init configuration as YAML failed: %w", err)
	}
	if err := k.clusterUtil.InitCluster(initConfigYAML); err != nil {
		return fmt.Errorf("kubeadm init failed: %w", err)
	}
	kubeConfig, err := k.GetKubeconfig()
	if err != nil {
		return fmt.Errorf("reading kubeconfig after cluster initialization failed: %w", err)
	}
	k.client.SetKubeconfig(kubeConfig)
	flannel := resources.NewDefaultFlannelDeployment()
	if err = k.clusterUtil.SetupPodNetwork(k.client, flannel); err != nil {
		return fmt.Errorf("setup of pod network failed: %w", err)
	}

	kms := resources.NewKMSDeployment(in.MasterSecret)
	if err = k.clusterUtil.SetupKMS(k.client, kms); err != nil {
		return fmt.Errorf("setup of kms failed: %w", err)
	}

	if in.SupportsCloudControllerManager {
		cloudControllerManagerConfiguration := resources.NewDefaultCloudControllerManagerDeployment(
			in.CloudControllerManagerName, in.CloudControllerManagerImage, in.CloudControllerManagerPath, in.CloudControllerManagerExtraArgs,
			in.CloudControllerManagerVolumes, in.CloudControllerManagerVolumeMounts, in.CloudControllerManagerEnv,
		)
		if err := k.clusterUtil.SetupCloudControllerManager(k.client, cloudControllerManagerConfiguration, in.CloudControllerManagerConfigMaps, in.CloudControllerManagerSecrets); err != nil {
			return fmt.Errorf("failed to setup cloud-controller-manager: %w", err)
		}
	}

	if in.SupportsCloudNodeManager {
		cloudNodeManagerConfiguration := resources.NewDefaultCloudNodeManagerDeployment(
			in.CloudNodeManagerImage, in.CloudNodeManagerPath, in.CloudNodeManagerExtraArgs,
		)
		if err := k.clusterUtil.SetupCloudNodeManager(k.client, cloudNodeManagerConfiguration); err != nil {
			return fmt.Errorf("failed to setup cloud-node-manager: %w", err)
		}
	}

	if in.SupportClusterAutoscaler {
		clusterAutoscalerConfiguration := resources.NewDefaultAutoscalerDeployment(in.AutoscalingVolumes, in.AutoscalingVolumeMounts, in.AutoscalingEnv)
		clusterAutoscalerConfiguration.SetAutoscalerCommand(in.AutoscalingCloudprovider, in.AutoscalingNodeGroups)
		if err := k.clusterUtil.SetupAutoscaling(k.client, clusterAutoscalerConfiguration, in.AutoscalingSecrets); err != nil {
			return fmt.Errorf("failed to setup cluster-autoscaler: %w", err)
		}
	}

	return nil
}

// JoinCluster joins existing Kubernetes cluster.
func (k *KubeWrapper) JoinCluster(args *kubeadm.BootstrapTokenDiscovery, nodeName, nodeInternalIP, nodeVPNIP, providerID, certKey string, ccmSupported bool, peerRole role.Role) error {
	// TODO: k8s version should be user input
	if err := k.clusterUtil.InstallComponents(context.TODO(), "1.23.6"); err != nil {
		return err
	}

	joinConfig := k.configProvider.JoinConfiguration(ccmSupported)
	joinConfig.SetApiServerEndpoint(args.APIServerEndpoint)
	joinConfig.SetToken(args.Token)
	joinConfig.AppendDiscoveryTokenCaCertHash(args.CACertHashes[0])
	joinConfig.SetNodeIP(nodeInternalIP)
	joinConfig.SetNodeName(nodeName)
	joinConfig.SetProviderID(providerID)
	if peerRole == role.Coordinator {
		joinConfig.SetControlPlane(nodeVPNIP, certKey)
	}
	joinConfigYAML, err := joinConfig.Marshal()
	if err != nil {
		return fmt.Errorf("encoding kubeadm join configuration as YAML failed: %w", err)
	}
	if err := k.clusterUtil.JoinCluster(joinConfigYAML); err != nil {
		return fmt.Errorf("joining cluster failed: %v %w ", string(joinConfigYAML), err)
	}

	return nil
}

// GetKubeconfig returns the current nodes kubeconfig of stored on disk.
func (k *KubeWrapper) GetKubeconfig() ([]byte, error) {
	kubeconf, err := k.kubeconfigReader.ReadKubeconfig()
	if err != nil {
		return nil, err
	}
	// replace the cluster.Server endpoint (127.0.0.1:16443) in admin.conf with the first coordinator endpoint (10.118.0.1:6443)
	// kube-api server listens on 10.118.0.1:6443
	// 127.0.0.1:16443 is the high availability balancer nginx endpoint, runnining localy on all nodes
	// alternatively one could also start a local high availability balancer.
	return []byte(strings.ReplaceAll(string(kubeconf), "127.0.0.1:16443", "10.118.0.1:6443")), nil
}

// GetKubeadmCertificateKey return the key needed to join the Cluster as Control-Plane (has to be executed on a control-plane; errors otherwise).
func (k *KubeWrapper) GetKubeadmCertificateKey() (string, error) {
	return k.clusterUtil.GetControlPlaneJoinCertificateKey()
}

// GetJoinToken returns a bootstrap (join) token.
func (k *KubeWrapper) GetJoinToken(ttl time.Duration) (*kubeadm.BootstrapTokenDiscovery, error) {
	return k.clusterUtil.CreateJoinToken(ttl)
}

// StartKubelet starts the kubelet service.
func (k *KubeWrapper) StartKubelet() error {
	return k.clusterUtil.StartKubelet()
}

type fakeK8SClient struct {
	kubeconfig []byte
}

// NewFakeK8SClient creates a new, fake k8s client where apply always works.
func NewFakeK8SClient([]byte) (k8sapi.Client, error) {
	return &fakeK8SClient{}, nil
}

// Apply fakes applying Kubernetes resources.
func (f *fakeK8SClient) Apply(resources resources.Marshaler, forceConflicts bool) error {
	return nil
}

// SetKubeconfig stores the kubeconfig given to it.
func (f *fakeK8SClient) SetKubeconfig(kubeconfig []byte) {
	f.kubeconfig = kubeconfig
}
