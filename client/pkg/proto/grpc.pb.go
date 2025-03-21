// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.27.3
// source: proto/grpc.proto

package uploadpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileUploadRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AgnosticFilePath string                 `protobuf:"bytes,1,opt,name=agnosticFilePath,proto3" json:"agnosticFilePath,omitempty"`
	Chunk            []byte                 `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	mi := &file_proto_grpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadRequest) GetAgnosticFilePath() string {
	if x != nil {
		return x.AgnosticFilePath
	}
	return ""
}

func (x *FileUploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FileUploadResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AgnosticFilePath string                 `protobuf:"bytes,1,opt,name=agnosticFilePath,proto3" json:"agnosticFilePath,omitempty"`
	Size             uint32                 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	mi := &file_proto_grpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadResponse) GetAgnosticFilePath() string {
	if x != nil {
		return x.AgnosticFilePath
	}
	return ""
}

func (x *FileUploadResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FileDownloadRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AgnosticFilePath string                 `protobuf:"bytes,1,opt,name=agnosticFilePath,proto3" json:"agnosticFilePath,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *FileDownloadRequest) Reset() {
	*x = FileDownloadRequest{}
	mi := &file_proto_grpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadRequest) ProtoMessage() {}

func (x *FileDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadRequest.ProtoReflect.Descriptor instead.
func (*FileDownloadRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *FileDownloadRequest) GetAgnosticFilePath() string {
	if x != nil {
		return x.AgnosticFilePath
	}
	return ""
}

type FileDownloadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileDownloadResponse) Reset() {
	*x = FileDownloadResponse{}
	mi := &file_proto_grpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadResponse) ProtoMessage() {}

func (x *FileDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadResponse.ProtoReflect.Descriptor instead.
func (*FileDownloadResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *FileDownloadResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type GetMetaDataMapRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetaDataMapRequest) Reset() {
	*x = GetMetaDataMapRequest{}
	mi := &file_proto_grpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetaDataMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetaDataMapRequest) ProtoMessage() {}

func (x *GetMetaDataMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetaDataMapRequest.ProtoReflect.Descriptor instead.
func (*GetMetaDataMapRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{4}
}

type FileMetaData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileCheckSum  string                 `protobuf:"bytes,1,opt,name=fileCheckSum,proto3" json:"fileCheckSum,omitempty"`
	LastModified  string                 `protobuf:"bytes,2,opt,name=LastModified,proto3" json:"LastModified,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileMetaData) Reset() {
	*x = FileMetaData{}
	mi := &file_proto_grpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetaData) ProtoMessage() {}

func (x *FileMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetaData.ProtoReflect.Descriptor instead.
func (*FileMetaData) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *FileMetaData) GetFileCheckSum() string {
	if x != nil {
		return x.FileCheckSum
	}
	return ""
}

func (x *FileMetaData) GetLastModified() string {
	if x != nil {
		return x.LastModified
	}
	return ""
}

type GetMetaDataMapResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	MetaDataMap   map[string]*FileMetaData `protobuf:"bytes,1,rep,name=metaDataMap,proto3" json:"metaDataMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetaDataMapResponse) Reset() {
	*x = GetMetaDataMapResponse{}
	mi := &file_proto_grpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetaDataMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetaDataMapResponse) ProtoMessage() {}

func (x *GetMetaDataMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetaDataMapResponse.ProtoReflect.Descriptor instead.
func (*GetMetaDataMapResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetMetaDataMapResponse) GetMetaDataMap() map[string]*FileMetaData {
	if x != nil {
		return x.MetaDataMap
	}
	return nil
}

var File_proto_grpc_proto protoreflect.FileDescriptor

var file_proto_grpc_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x22, 0x54, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x41, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x0a, 0x14, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x56, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x4d, 0x61, 0x70, 0x1a, 0x53, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe4, 0x01, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x08, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x61, 0x70, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_grpc_proto_rawDescOnce sync.Once
	file_proto_grpc_proto_rawDescData []byte
)

func file_proto_grpc_proto_rawDescGZIP() []byte {
	file_proto_grpc_proto_rawDescOnce.Do(func() {
		file_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_grpc_proto_rawDesc), len(file_proto_grpc_proto_rawDesc)))
	})
	return file_proto_grpc_proto_rawDescData
}

var file_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_grpc_proto_goTypes = []any{
	(*FileUploadRequest)(nil),      // 0: proto.FileUploadRequest
	(*FileUploadResponse)(nil),     // 1: proto.FileUploadResponse
	(*FileDownloadRequest)(nil),    // 2: proto.FileDownloadRequest
	(*FileDownloadResponse)(nil),   // 3: proto.FileDownloadResponse
	(*GetMetaDataMapRequest)(nil),  // 4: proto.GetMetaDataMapRequest
	(*FileMetaData)(nil),           // 5: proto.FileMetaData
	(*GetMetaDataMapResponse)(nil), // 6: proto.GetMetaDataMapResponse
	nil,                            // 7: proto.GetMetaDataMapResponse.MetaDataMapEntry
}
var file_proto_grpc_proto_depIdxs = []int32{
	7, // 0: proto.GetMetaDataMapResponse.metaDataMap:type_name -> proto.GetMetaDataMapResponse.MetaDataMapEntry
	5, // 1: proto.GetMetaDataMapResponse.MetaDataMapEntry.value:type_name -> proto.FileMetaData
	0, // 2: proto.FileService.Upload:input_type -> proto.FileUploadRequest
	2, // 3: proto.FileService.Download:input_type -> proto.FileDownloadRequest
	4, // 4: proto.FileService.GetMetaDataMap:input_type -> proto.GetMetaDataMapRequest
	1, // 5: proto.FileService.Upload:output_type -> proto.FileUploadResponse
	3, // 6: proto.FileService.Download:output_type -> proto.FileDownloadResponse
	6, // 7: proto.FileService.GetMetaDataMap:output_type -> proto.GetMetaDataMapResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_grpc_proto_init() }
func file_proto_grpc_proto_init() {
	if File_proto_grpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_grpc_proto_rawDesc), len(file_proto_grpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_proto_goTypes,
		DependencyIndexes: file_proto_grpc_proto_depIdxs,
		MessageInfos:      file_proto_grpc_proto_msgTypes,
	}.Build()
	File_proto_grpc_proto = out.File
	file_proto_grpc_proto_goTypes = nil
	file_proto_grpc_proto_depIdxs = nil
}
