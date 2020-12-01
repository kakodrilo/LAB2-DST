// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: Lab_2.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string  `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ChunksDn1 []int32 `protobuf:"varint,2,rep,packed,name=chunksDn1,proto3" json:"chunksDn1,omitempty"`
	ChunksDn2 []int32 `protobuf:"varint,3,rep,packed,name=chunksDn2,proto3" json:"chunksDn2,omitempty"`
	ChunksDn3 []int32 `protobuf:"varint,4,rep,packed,name=chunksDn3,proto3" json:"chunksDn3,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{2}
}

func (x *Proposal) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Proposal) GetChunksDn1() []int32 {
	if x != nil {
		return x.ChunksDn1
	}
	return nil
}

func (x *Proposal) GetChunksDn2() []int32 {
	if x != nil {
		return x.ChunksDn2
	}
	return nil
}

func (x *Proposal) GetChunksDn3() []int32 {
	if x != nil {
		return x.ChunksDn3
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Chunk   []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{3}
}

func (x *Chunk) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type RequestChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
}

func (x *RequestChunk) Reset() {
	*x = RequestChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChunk) ProtoMessage() {}

func (x *RequestChunk) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChunk.ProtoReflect.Descriptor instead.
func (*RequestChunk) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{4}
}

func (x *RequestChunk) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{5}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type ChunkAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip []string `protobuf:"bytes,1,rep,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ChunkAddress) Reset() {
	*x = ChunkAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lab_2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkAddress) ProtoMessage() {}

func (x *ChunkAddress) ProtoReflect() protoreflect.Message {
	mi := &file_Lab_2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkAddress.ProtoReflect.Descriptor instead.
func (*ChunkAddress) Descriptor() ([]byte, []int) {
	return file_Lab_2_proto_rawDescGZIP(), []int{6}
}

func (x *ChunkAddress) GetIp() []string {
	if x != nil {
		return x.Ip
	}
	return nil
}

var File_Lab_2_proto protoreflect.FileDescriptor

var file_Lab_2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x5f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x44, 0x6e, 0x31, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x44, 0x6e, 0x31, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x44, 0x6e, 0x32, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x44, 0x6e, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x44, 0x6e, 0x33, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x44, 0x6e, 0x33, 0x22, 0x38, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22,
	0x28, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a,
	0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x32, 0xba, 0x01,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x2d, 0x0a, 0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9b, 0x02, 0x0a, 0x08, 0x4e,
	0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x13, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69,
	0x74, 0x65, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Lab_2_proto_rawDescOnce sync.Once
	file_Lab_2_proto_rawDescData = file_Lab_2_proto_rawDesc
)

func file_Lab_2_proto_rawDescGZIP() []byte {
	file_Lab_2_proto_rawDescOnce.Do(func() {
		file_Lab_2_proto_rawDescData = protoimpl.X.CompressGZIP(file_Lab_2_proto_rawDescData)
	})
	return file_Lab_2_proto_rawDescData
}

var file_Lab_2_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Lab_2_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: pb.Empty
	(*Response)(nil),     // 1: pb.Response
	(*Proposal)(nil),     // 2: pb.Proposal
	(*Chunk)(nil),        // 3: pb.Chunk
	(*RequestChunk)(nil), // 4: pb.RequestChunk
	(*File)(nil),         // 5: pb.File
	(*ChunkAddress)(nil), // 6: pb.ChunkAddress
}
var file_Lab_2_proto_depIdxs = []int32{
	3,  // 0: pb.DataNode.UploadChunks:input_type -> pb.Chunk
	4,  // 1: pb.DataNode.DownloadChunks:input_type -> pb.RequestChunk
	3,  // 2: pb.DataNode.SaveChunks:input_type -> pb.Chunk
	2,  // 3: pb.DataNode.ProposalRequest:input_type -> pb.Proposal
	0,  // 4: pb.NameNode.RequestWrite:input_type -> pb.Empty
	2,  // 5: pb.NameNode.CentralizedWriteLog:input_type -> pb.Proposal
	2,  // 6: pb.NameNode.DistribuitedWriteLog:input_type -> pb.Proposal
	2,  // 7: pb.NameNode.FinalProposal:input_type -> pb.Proposal
	0,  // 8: pb.NameNode.FileRequest:input_type -> pb.Empty
	5,  // 9: pb.NameNode.AddressRequest:input_type -> pb.File
	1,  // 10: pb.DataNode.UploadChunks:output_type -> pb.Response
	3,  // 11: pb.DataNode.DownloadChunks:output_type -> pb.Chunk
	1,  // 12: pb.DataNode.SaveChunks:output_type -> pb.Response
	1,  // 13: pb.DataNode.ProposalRequest:output_type -> pb.Response
	1,  // 14: pb.NameNode.RequestWrite:output_type -> pb.Response
	1,  // 15: pb.NameNode.CentralizedWriteLog:output_type -> pb.Response
	1,  // 16: pb.NameNode.DistribuitedWriteLog:output_type -> pb.Response
	2,  // 17: pb.NameNode.FinalProposal:output_type -> pb.Proposal
	5,  // 18: pb.NameNode.FileRequest:output_type -> pb.File
	6,  // 19: pb.NameNode.AddressRequest:output_type -> pb.ChunkAddress
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_Lab_2_proto_init() }
func file_Lab_2_proto_init() {
	if File_Lab_2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Lab_2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_Lab_2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_Lab_2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_Lab_2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_Lab_2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChunk); i {
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
		file_Lab_2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_Lab_2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkAddress); i {
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
			RawDescriptor: file_Lab_2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Lab_2_proto_goTypes,
		DependencyIndexes: file_Lab_2_proto_depIdxs,
		MessageInfos:      file_Lab_2_proto_msgTypes,
	}.Build()
	File_Lab_2_proto = out.File
	file_Lab_2_proto_rawDesc = nil
	file_Lab_2_proto_goTypes = nil
	file_Lab_2_proto_depIdxs = nil
}
