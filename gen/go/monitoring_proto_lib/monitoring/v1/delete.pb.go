// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: monitoring_proto_lib/monitoring/v1/delete.proto

package v1

import (
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

type NotifyDeletionRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	JobName        string                 `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	InstanceNumber int32                  `protobuf:"varint,2,opt,name=instance_number,json=instanceNumber,proto3" json:"instance_number,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *NotifyDeletionRequest) Reset() {
	*x = NotifyDeletionRequest{}
	mi := &file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyDeletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyDeletionRequest) ProtoMessage() {}

func (x *NotifyDeletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyDeletionRequest.ProtoReflect.Descriptor instead.
func (*NotifyDeletionRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyDeletionRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *NotifyDeletionRequest) GetInstanceNumber() int32 {
	if x != nil {
		return x.InstanceNumber
	}
	return 0
}

type NotifyDeletionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Acknowledged  bool                   `protobuf:"varint,1,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotifyDeletionResponse) Reset() {
	*x = NotifyDeletionResponse{}
	mi := &file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyDeletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyDeletionResponse) ProtoMessage() {}

func (x *NotifyDeletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyDeletionResponse.ProtoReflect.Descriptor instead.
func (*NotifyDeletionResponse) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyDeletionResponse) GetAcknowledged() bool {
	if x != nil {
		return x.Acknowledged
	}
	return false
}

func (x *NotifyDeletionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_monitoring_proto_lib_monitoring_v1_delete_proto protoreflect.FileDescriptor

var file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x22, 0x5b, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x56, 0x0a,
	0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x6e, 0x7a, 0x6c, 0x6e, 0x73, 0x6b, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6c, 0x69,
	0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescOnce sync.Once
	file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescData = file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDesc
)

func file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescGZIP() []byte {
	file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescOnce.Do(func() {
		file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescData)
	})
	return file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDescData
}

var file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_monitoring_proto_lib_monitoring_v1_delete_proto_goTypes = []any{
	(*NotifyDeletionRequest)(nil),  // 0: monitoring.v1.NotifyDeletionRequest
	(*NotifyDeletionResponse)(nil), // 1: monitoring.v1.NotifyDeletionResponse
}
var file_monitoring_proto_lib_monitoring_v1_delete_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitoring_proto_lib_monitoring_v1_delete_proto_init() }
func file_monitoring_proto_lib_monitoring_v1_delete_proto_init() {
	if File_monitoring_proto_lib_monitoring_v1_delete_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_monitoring_proto_lib_monitoring_v1_delete_proto_goTypes,
		DependencyIndexes: file_monitoring_proto_lib_monitoring_v1_delete_proto_depIdxs,
		MessageInfos:      file_monitoring_proto_lib_monitoring_v1_delete_proto_msgTypes,
	}.Build()
	File_monitoring_proto_lib_monitoring_v1_delete_proto = out.File
	file_monitoring_proto_lib_monitoring_v1_delete_proto_rawDesc = nil
	file_monitoring_proto_lib_monitoring_v1_delete_proto_goTypes = nil
	file_monitoring_proto_lib_monitoring_v1_delete_proto_depIdxs = nil
}
