// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: definition.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PipelineDefinitionProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline yml source
	PipelineSource string `protobuf:"bytes,1,opt,name=pipelineSource,proto3" json:"pipelineSource,omitempty"`
	// pipeline yml name
	PipelineYmlName string `protobuf:"bytes,2,opt,name=pipelineYmlName,proto3" json:"pipelineYmlName,omitempty"`
	// pipeline yml content
	PipelineYml string `protobuf:"bytes,3,opt,name=pipelineYml,proto3" json:"pipelineYml,omitempty"`
	// snippetConfig means definition can be quoted by snippet
	SnippetConfig *structpb.Value `protobuf:"bytes,4,opt,name=snippetConfig,proto3" json:"snippetConfig,omitempty"`
	// versionLock means db optimistic lock
	VersionLock uint64 `protobuf:"varint,5,opt,name=versionLock,proto3" json:"versionLock,omitempty"`
	// isDelete
	// - if = true send With pipelineSource and pipelineYmlName to request
	IsDelete bool `protobuf:"varint,6,opt,name=isDelete,proto3" json:"isDelete,omitempty"`
	// pipelineCreateRequest means definition can autoRun
	PipelineCreateRequest *structpb.Value `protobuf:"bytes,7,opt,name=pipelineCreateRequest,proto3" json:"pipelineCreateRequest,omitempty"`
}

func (x *PipelineDefinitionProcessRequest) Reset() {
	*x = PipelineDefinitionProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineDefinitionProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineDefinitionProcessRequest) ProtoMessage() {}

func (x *PipelineDefinitionProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineDefinitionProcessRequest.ProtoReflect.Descriptor instead.
func (*PipelineDefinitionProcessRequest) Descriptor() ([]byte, []int) {
	return file_definition_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineDefinitionProcessRequest) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *PipelineDefinitionProcessRequest) GetPipelineYmlName() string {
	if x != nil {
		return x.PipelineYmlName
	}
	return ""
}

func (x *PipelineDefinitionProcessRequest) GetPipelineYml() string {
	if x != nil {
		return x.PipelineYml
	}
	return ""
}

func (x *PipelineDefinitionProcessRequest) GetSnippetConfig() *structpb.Value {
	if x != nil {
		return x.SnippetConfig
	}
	return nil
}

func (x *PipelineDefinitionProcessRequest) GetVersionLock() uint64 {
	if x != nil {
		return x.VersionLock
	}
	return 0
}

func (x *PipelineDefinitionProcessRequest) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

func (x *PipelineDefinitionProcessRequest) GetPipelineCreateRequest() *structpb.Value {
	if x != nil {
		return x.PipelineCreateRequest
	}
	return nil
}

type PipelineDefinitionProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PipelineSource  string `protobuf:"bytes,2,opt,name=pipelineSource,proto3" json:"pipelineSource,omitempty"`
	PipelineYmlName string `protobuf:"bytes,3,opt,name=pipelineYmlName,proto3" json:"pipelineYmlName,omitempty"`
	PipelineYml     string `protobuf:"bytes,4,opt,name=pipelineYml,proto3" json:"pipelineYml,omitempty"`
	VersionLock     uint64 `protobuf:"varint,5,opt,name=versionLock,proto3" json:"versionLock,omitempty"`
}

func (x *PipelineDefinitionProcessResponse) Reset() {
	*x = PipelineDefinitionProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineDefinitionProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineDefinitionProcessResponse) ProtoMessage() {}

func (x *PipelineDefinitionProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineDefinitionProcessResponse.ProtoReflect.Descriptor instead.
func (*PipelineDefinitionProcessResponse) Descriptor() ([]byte, []int) {
	return file_definition_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineDefinitionProcessResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PipelineDefinitionProcessResponse) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *PipelineDefinitionProcessResponse) GetPipelineYmlName() string {
	if x != nil {
		return x.PipelineYmlName
	}
	return ""
}

func (x *PipelineDefinitionProcessResponse) GetPipelineYml() string {
	if x != nil {
		return x.PipelineYml
	}
	return ""
}

func (x *PipelineDefinitionProcessResponse) GetVersionLock() uint64 {
	if x != nil {
		return x.VersionLock
	}
	return 0
}

type PipelineDefinitionProcessVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline yml source
	PipelineSource string `protobuf:"bytes,1,opt,name=pipelineSource,proto3" json:"pipelineSource,omitempty"`
	// pipeline yml name
	PipelineYmlName string `protobuf:"bytes,2,opt,name=pipelineYmlName,proto3" json:"pipelineYmlName,omitempty"`
}

func (x *PipelineDefinitionProcessVersionRequest) Reset() {
	*x = PipelineDefinitionProcessVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineDefinitionProcessVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineDefinitionProcessVersionRequest) ProtoMessage() {}

func (x *PipelineDefinitionProcessVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineDefinitionProcessVersionRequest.ProtoReflect.Descriptor instead.
func (*PipelineDefinitionProcessVersionRequest) Descriptor() ([]byte, []int) {
	return file_definition_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineDefinitionProcessVersionRequest) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *PipelineDefinitionProcessVersionRequest) GetPipelineYmlName() string {
	if x != nil {
		return x.PipelineYmlName
	}
	return ""
}

type PipelineDefinitionProcessVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionLock uint64 `protobuf:"varint,1,opt,name=versionLock,proto3" json:"versionLock,omitempty"`
}

func (x *PipelineDefinitionProcessVersionResponse) Reset() {
	*x = PipelineDefinitionProcessVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_definition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineDefinitionProcessVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineDefinitionProcessVersionResponse) ProtoMessage() {}

func (x *PipelineDefinitionProcessVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_definition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineDefinitionProcessVersionResponse.ProtoReflect.Descriptor instead.
func (*PipelineDefinitionProcessVersionResponse) Descriptor() ([]byte, []int) {
	return file_definition_proto_rawDescGZIP(), []int{3}
}

func (x *PipelineDefinitionProcessVersionResponse) GetVersionLock() uint64 {
	if x != nil {
		return x.VersionLock
	}
	return 0
}

var File_definition_proto protoreflect.FileDescriptor

var file_definition_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02,
	0x0a, 0x20, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x59, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x4c, 0x0a, 0x15, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xc9, 0x01, 0x0a, 0x21, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x22, 0x7b, 0x0a, 0x27,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x28, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x32, 0xa5, 0x03, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12,
	0xcd, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72,
	0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_definition_proto_rawDescOnce sync.Once
	file_definition_proto_rawDescData = file_definition_proto_rawDesc
)

func file_definition_proto_rawDescGZIP() []byte {
	file_definition_proto_rawDescOnce.Do(func() {
		file_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_definition_proto_rawDescData)
	})
	return file_definition_proto_rawDescData
}

var file_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_definition_proto_goTypes = []interface{}{
	(*PipelineDefinitionProcessRequest)(nil),         // 0: erda.core.pipeline.definition.PipelineDefinitionProcessRequest
	(*PipelineDefinitionProcessResponse)(nil),        // 1: erda.core.pipeline.definition.PipelineDefinitionProcessResponse
	(*PipelineDefinitionProcessVersionRequest)(nil),  // 2: erda.core.pipeline.definition.PipelineDefinitionProcessVersionRequest
	(*PipelineDefinitionProcessVersionResponse)(nil), // 3: erda.core.pipeline.definition.PipelineDefinitionProcessVersionResponse
	(*structpb.Value)(nil),                           // 4: google.protobuf.Value
}
var file_definition_proto_depIdxs = []int32{
	4, // 0: erda.core.pipeline.definition.PipelineDefinitionProcessRequest.snippetConfig:type_name -> google.protobuf.Value
	4, // 1: erda.core.pipeline.definition.PipelineDefinitionProcessRequest.pipelineCreateRequest:type_name -> google.protobuf.Value
	0, // 2: erda.core.pipeline.definition.DefinitionService.Process:input_type -> erda.core.pipeline.definition.PipelineDefinitionProcessRequest
	2, // 3: erda.core.pipeline.definition.DefinitionService.Version:input_type -> erda.core.pipeline.definition.PipelineDefinitionProcessVersionRequest
	1, // 4: erda.core.pipeline.definition.DefinitionService.Process:output_type -> erda.core.pipeline.definition.PipelineDefinitionProcessResponse
	3, // 5: erda.core.pipeline.definition.DefinitionService.Version:output_type -> erda.core.pipeline.definition.PipelineDefinitionProcessVersionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_definition_proto_init() }
func file_definition_proto_init() {
	if File_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineDefinitionProcessRequest); i {
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
		file_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineDefinitionProcessResponse); i {
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
		file_definition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineDefinitionProcessVersionRequest); i {
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
		file_definition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineDefinitionProcessVersionResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_definition_proto_goTypes,
		DependencyIndexes: file_definition_proto_depIdxs,
		MessageInfos:      file_definition_proto_msgTypes,
	}.Build()
	File_definition_proto = out.File
	file_definition_proto_rawDesc = nil
	file_definition_proto_goTypes = nil
	file_definition_proto_depIdxs = nil
}
