/*
Package constants contains the constants used by Constellation.
Constants should never be overwritable by command line flags or configuration files.
*/
package constants

import "time"

const (
	//
	// Constellation.
	//

	// ConstellationNameLength is the maximum length of a Constellation's name.
	ConstellationNameLength = 37
	// ConstellationMasterSecretStoreName is the name for the Constellation secrets in Kubernetes.
	ConstellationMasterSecretStoreName = "constellation-mastersecret"
	// ConstellationMasterSecretKey is the name of the key for master secret in the master secret store secret.
	ConstellationMasterSecretKey = "mastersecret"

	//
	// Ports.
	//

	CoordinatorPort = 9000
	EnclaveSSHPort  = 2222
	SSHPort         = 22
	WireguardPort   = 51820
	NVMEOverTCPPort = 8009
	// Default NodePort Range
	// https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	NodePortFrom = 30000
	NodePortTo   = 32767

	//
	// Filenames.
	//

	StateFilename         = "constellation-state.json"
	ConfigFilename        = "constellation-conf.yaml"
	DebugdConfigFilename  = "cdbg-conf.yaml"
	AdminConfFilename     = "constellation-admin.conf"
	MasterSecretFilename  = "constellation-mastersecret.base64"
	WGQuickConfigFilename = "wg0.conf"

	//
	// Cryptographic constants.
	//

	StateDiskKeyLength = 32
	// DerivedKeyLengthDefault is the default length in bytes for KMS derived keys.
	DerivedKeyLengthDefault = 32
	// MasterSecretLengthDefault is the default length in bytes for CLI generated master secrets.
	MasterSecretLengthDefault = 32
	// MasterSecretLengthMin is the minimal length in bytes for user provided master secrets.
	MasterSecretLengthMin = 16

	//
	// CLI.
	//

	MinControllerCount = 1
	MinWorkerCount     = 1

	//
	// Kubernetes.
	//

	// KubernetesVersion installed by kubeadm.
	KubernetesVersion      = "stable-1.23"
	KubernetesJoinTokenTTL = 15 * time.Minute

	//
	// VPN.
	//

	// WireguardAdminMTU is the MTU designated for the admin's WireGuard interface.
	// WireGuard doesn't support Path MTU Discovery. Thus, its default MTU can be too high on some networks.
	WireguardAdminMTU = 1300
)

// CliVersion is the version of the CLI. Left as a separate variable to allow override during build.
var CliVersion = "0.0.0"
