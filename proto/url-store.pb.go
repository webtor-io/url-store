// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: url-store.proto

package url_store

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

// The push response message
type PushReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *PushReply) Reset() {
	*x = PushReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushReply) ProtoMessage() {}

func (x *PushReply) ProtoReflect() protoreflect.Message {
	mi := &file_url_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushReply.ProtoReflect.Descriptor instead.
func (*PushReply) Descriptor() ([]byte, []int) {
	return file_url_store_proto_rawDescGZIP(), []int{0}
}

func (x *PushReply) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// The push request message
type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_url_store_proto_rawDescGZIP(), []int{1}
}

func (x *PushRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// The check request message containing the infoHash
type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_url_store_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// The check response message containing existance flag
type CheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *CheckReply) Reset() {
	*x = CheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReply) ProtoMessage() {}

func (x *CheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_url_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReply.ProtoReflect.Descriptor instead.
func (*CheckReply) Descriptor() ([]byte, []int) {
	return file_url_store_proto_rawDescGZIP(), []int{3}
}

func (x *CheckReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_url_store_proto protoreflect.FileDescriptor

var file_url_store_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x6f,
	0x72, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x24, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x95,
	0x01, 0x0a, 0x08, 0x55, 0x72, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x50,
	0x75, 0x73, 0x68, 0x12, 0x1c, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x2e, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x6f,
	0x72, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x72,
	0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x75,
	0x72, 0x6c, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_store_proto_rawDescOnce sync.Once
	file_url_store_proto_rawDescData = file_url_store_proto_rawDesc
)

func file_url_store_proto_rawDescGZIP() []byte {
	file_url_store_proto_rawDescOnce.Do(func() {
		file_url_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_store_proto_rawDescData)
	})
	return file_url_store_proto_rawDescData
}

var file_url_store_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_url_store_proto_goTypes = []interface{}{
	(*PushReply)(nil),    // 0: webtor.url_sore.PushReply
	(*PushRequest)(nil),  // 1: webtor.url_sore.PushRequest
	(*CheckRequest)(nil), // 2: webtor.url_sore.CheckRequest
	(*CheckReply)(nil),   // 3: webtor.url_sore.CheckReply
}
var file_url_store_proto_depIdxs = []int32{
	1, // 0: webtor.url_sore.UrlStore.Push:input_type -> webtor.url_sore.PushRequest
	2, // 1: webtor.url_sore.UrlStore.Check:input_type -> webtor.url_sore.CheckRequest
	0, // 2: webtor.url_sore.UrlStore.Push:output_type -> webtor.url_sore.PushReply
	3, // 3: webtor.url_sore.UrlStore.Check:output_type -> webtor.url_sore.CheckReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_store_proto_init() }
func file_url_store_proto_init() {
	if File_url_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushReply); i {
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
		file_url_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_url_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_url_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckReply); i {
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
			RawDescriptor: file_url_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_store_proto_goTypes,
		DependencyIndexes: file_url_store_proto_depIdxs,
		MessageInfos:      file_url_store_proto_msgTypes,
	}.Build()
	File_url_store_proto = out.File
	file_url_store_proto_rawDesc = nil
	file_url_store_proto_goTypes = nil
	file_url_store_proto_depIdxs = nil
}