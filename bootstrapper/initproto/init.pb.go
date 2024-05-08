// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.22.1
// source: bootstrapper/initproto/init.proto

package initproto

import (
	context "context"
	components "github.com/edgelesssys/constellation/v2/internal/versions/components"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KmsUri               string                  `protobuf:"bytes,1,opt,name=kms_uri,json=kmsUri,proto3" json:"kms_uri,omitempty"`
	StorageUri           string                  `protobuf:"bytes,2,opt,name=storage_uri,json=storageUri,proto3" json:"storage_uri,omitempty"`
	MeasurementSalt      []byte                  `protobuf:"bytes,3,opt,name=measurement_salt,json=measurementSalt,proto3" json:"measurement_salt,omitempty"`
	KubernetesVersion    string                  `protobuf:"bytes,5,opt,name=kubernetes_version,json=kubernetesVersion,proto3" json:"kubernetes_version,omitempty"`
	ConformanceMode      bool                    `protobuf:"varint,6,opt,name=conformance_mode,json=conformanceMode,proto3" json:"conformance_mode,omitempty"`
	KubernetesComponents []*components.Component `protobuf:"bytes,7,rep,name=kubernetes_components,json=kubernetesComponents,proto3" json:"kubernetes_components,omitempty"`
	InitSecret           []byte                  `protobuf:"bytes,8,opt,name=init_secret,json=initSecret,proto3" json:"init_secret,omitempty"`
	ClusterName          string                  `protobuf:"bytes,9,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	ApiserverCertSans    []string                `protobuf:"bytes,10,rep,name=apiserver_cert_sans,json=apiserverCertSans,proto3" json:"apiserver_cert_sans,omitempty"`
	ServiceCidr          string                  `protobuf:"bytes,11,opt,name=service_cidr,json=serviceCidr,proto3" json:"service_cidr,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetKmsUri() string {
	if x != nil {
		return x.KmsUri
	}
	return ""
}

func (x *InitRequest) GetStorageUri() string {
	if x != nil {
		return x.StorageUri
	}
	return ""
}

func (x *InitRequest) GetMeasurementSalt() []byte {
	if x != nil {
		return x.MeasurementSalt
	}
	return nil
}

func (x *InitRequest) GetKubernetesVersion() string {
	if x != nil {
		return x.KubernetesVersion
	}
	return ""
}

func (x *InitRequest) GetConformanceMode() bool {
	if x != nil {
		return x.ConformanceMode
	}
	return false
}

func (x *InitRequest) GetKubernetesComponents() []*components.Component {
	if x != nil {
		return x.KubernetesComponents
	}
	return nil
}

func (x *InitRequest) GetInitSecret() []byte {
	if x != nil {
		return x.InitSecret
	}
	return nil
}

func (x *InitRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *InitRequest) GetApiserverCertSans() []string {
	if x != nil {
		return x.ApiserverCertSans
	}
	return nil
}

func (x *InitRequest) GetServiceCidr() string {
	if x != nil {
		return x.ServiceCidr
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*InitResponse_InitSuccess
	//	*InitResponse_InitFailure
	//	*InitResponse_Log
	Kind isInitResponse_Kind `protobuf_oneof:"kind"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{1}
}

func (m *InitResponse) GetKind() isInitResponse_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *InitResponse) GetInitSuccess() *InitSuccessResponse {
	if x, ok := x.GetKind().(*InitResponse_InitSuccess); ok {
		return x.InitSuccess
	}
	return nil
}

func (x *InitResponse) GetInitFailure() *InitFailureResponse {
	if x, ok := x.GetKind().(*InitResponse_InitFailure); ok {
		return x.InitFailure
	}
	return nil
}

func (x *InitResponse) GetLog() *LogResponseType {
	if x, ok := x.GetKind().(*InitResponse_Log); ok {
		return x.Log
	}
	return nil
}

type isInitResponse_Kind interface {
	isInitResponse_Kind()
}

type InitResponse_InitSuccess struct {
	InitSuccess *InitSuccessResponse `protobuf:"bytes,1,opt,name=init_success,json=initSuccess,proto3,oneof"`
}

type InitResponse_InitFailure struct {
	InitFailure *InitFailureResponse `protobuf:"bytes,2,opt,name=init_failure,json=initFailure,proto3,oneof"`
}

type InitResponse_Log struct {
	Log *LogResponseType `protobuf:"bytes,3,opt,name=log,proto3,oneof"`
}

func (*InitResponse_InitSuccess) isInitResponse_Kind() {}

func (*InitResponse_InitFailure) isInitResponse_Kind() {}

func (*InitResponse_Log) isInitResponse_Kind() {}

type InitSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubeconfig []byte `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	OwnerId    []byte `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ClusterId  []byte `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *InitSuccessResponse) Reset() {
	*x = InitSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSuccessResponse) ProtoMessage() {}

func (x *InitSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSuccessResponse.ProtoReflect.Descriptor instead.
func (*InitSuccessResponse) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{2}
}

func (x *InitSuccessResponse) GetKubeconfig() []byte {
	if x != nil {
		return x.Kubeconfig
	}
	return nil
}

func (x *InitSuccessResponse) GetOwnerId() []byte {
	if x != nil {
		return x.OwnerId
	}
	return nil
}

func (x *InitSuccessResponse) GetClusterId() []byte {
	if x != nil {
		return x.ClusterId
	}
	return nil
}

type InitFailureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *InitFailureResponse) Reset() {
	*x = InitFailureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitFailureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitFailureResponse) ProtoMessage() {}

func (x *InitFailureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitFailureResponse.ProtoReflect.Descriptor instead.
func (*InitFailureResponse) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{3}
}

func (x *InitFailureResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LogResponseType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log []byte `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *LogResponseType) Reset() {
	*x = LogResponseType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponseType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponseType) ProtoMessage() {}

func (x *LogResponseType) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponseType.ProtoReflect.Descriptor instead.
func (*LogResponseType) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{4}
}

func (x *LogResponseType) GetLog() []byte {
	if x != nil {
		return x.Log
	}
	return nil
}

type KubernetesComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Hash        string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	InstallPath string `protobuf:"bytes,3,opt,name=install_path,json=installPath,proto3" json:"install_path,omitempty"`
	Extract     bool   `protobuf:"varint,4,opt,name=extract,proto3" json:"extract,omitempty"`
}

func (x *KubernetesComponent) Reset() {
	*x = KubernetesComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrapper_initproto_init_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesComponent) ProtoMessage() {}

func (x *KubernetesComponent) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrapper_initproto_init_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesComponent.ProtoReflect.Descriptor instead.
func (*KubernetesComponent) Descriptor() ([]byte, []int) {
	return file_bootstrapper_initproto_init_proto_rawDescGZIP(), []int{5}
}

func (x *KubernetesComponent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *KubernetesComponent) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *KubernetesComponent) GetInstallPath() string {
	if x != nil {
		return x.InstallPath
	}
	return ""
}

func (x *KubernetesComponent) GetExtract() bool {
	if x != nil {
		return x.Extract
	}
	return false
}

var File_bootstrapper_initproto_init_proto protoreflect.FileDescriptor

var file_bootstrapper_initproto_init_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x69, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x6d, 0x73, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x6d, 0x73, 0x55, 0x72,
	0x69, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x69, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x2d, 0x0a,
	0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x4a, 0x0a, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x14, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x61, 0x6e, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x65, 0x72, 0x74, 0x53, 0x61, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x69, 0x64, 0x72, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x52, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x22, 0xc1, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x69, 0x6e, 0x69, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x69, 0x6e, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x69, 0x74,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x48, 0x00, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22,
	0x6f, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2b, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x23, 0x0a,
	0x0f, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6c,
	0x6f, 0x67, 0x22, 0x78, 0x0a, 0x13, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x32, 0x36, 0x0a, 0x03,
	0x41, 0x50, 0x49, 0x12, 0x2f, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x69, 0x6e,
	0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x73, 0x79, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x69,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootstrapper_initproto_init_proto_rawDescOnce sync.Once
	file_bootstrapper_initproto_init_proto_rawDescData = file_bootstrapper_initproto_init_proto_rawDesc
)

func file_bootstrapper_initproto_init_proto_rawDescGZIP() []byte {
	file_bootstrapper_initproto_init_proto_rawDescOnce.Do(func() {
		file_bootstrapper_initproto_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootstrapper_initproto_init_proto_rawDescData)
	})
	return file_bootstrapper_initproto_init_proto_rawDescData
}

var file_bootstrapper_initproto_init_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bootstrapper_initproto_init_proto_goTypes = []interface{}{
	(*InitRequest)(nil),          // 0: init.InitRequest
	(*InitResponse)(nil),         // 1: init.InitResponse
	(*InitSuccessResponse)(nil),  // 2: init.InitSuccessResponse
	(*InitFailureResponse)(nil),  // 3: init.InitFailureResponse
	(*LogResponseType)(nil),      // 4: init.LogResponseType
	(*KubernetesComponent)(nil),  // 5: init.KubernetesComponent
	(*components.Component)(nil), // 6: components.Component
}
var file_bootstrapper_initproto_init_proto_depIdxs = []int32{
	6, // 0: init.InitRequest.kubernetes_components:type_name -> components.Component
	2, // 1: init.InitResponse.init_success:type_name -> init.InitSuccessResponse
	3, // 2: init.InitResponse.init_failure:type_name -> init.InitFailureResponse
	4, // 3: init.InitResponse.log:type_name -> init.LogResponseType
	0, // 4: init.API.Init:input_type -> init.InitRequest
	1, // 5: init.API.Init:output_type -> init.InitResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bootstrapper_initproto_init_proto_init() }
func file_bootstrapper_initproto_init_proto_init() {
	if File_bootstrapper_initproto_init_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bootstrapper_initproto_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bootstrapper_initproto_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bootstrapper_initproto_init_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSuccessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bootstrapper_initproto_init_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitFailureResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bootstrapper_initproto_init_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponseType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bootstrapper_initproto_init_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_bootstrapper_initproto_init_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InitResponse_InitSuccess)(nil),
		(*InitResponse_InitFailure)(nil),
		(*InitResponse_Log)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bootstrapper_initproto_init_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bootstrapper_initproto_init_proto_goTypes,
		DependencyIndexes: file_bootstrapper_initproto_init_proto_depIdxs,
		MessageInfos:      file_bootstrapper_initproto_init_proto_msgTypes,
	}.Build()
	File_bootstrapper_initproto_init_proto = out.File
	file_bootstrapper_initproto_init_proto_rawDesc = nil
	file_bootstrapper_initproto_init_proto_goTypes = nil
	file_bootstrapper_initproto_init_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (API_InitClient, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (API_InitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/init.API/Init", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIInitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_InitClient interface {
	Recv() (*InitResponse, error)
	grpc.ClientStream
}

type aPIInitClient struct {
	grpc.ClientStream
}

func (x *aPIInitClient) Recv() (*InitResponse, error) {
	m := new(InitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Init(*InitRequest, API_InitServer) error
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Init(*InitRequest, API_InitServer) error {
	return status.Errorf(codes.Unimplemented, "method Init not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Init_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Init(m, &aPIInitServer{stream})
}

type API_InitServer interface {
	Send(*InitResponse) error
	grpc.ServerStream
}

type aPIInitServer struct {
	grpc.ServerStream
}

func (x *aPIInitServer) Send(m *InitResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "init.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Init",
			Handler:       _API_Init_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bootstrapper/initproto/init.proto",
}
