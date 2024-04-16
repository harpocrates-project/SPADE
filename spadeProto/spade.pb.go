// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: spade.proto

package __

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

// Define empty message for functions returning nil
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[0]
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
	return file_spade_proto_rawDescGZIP(), []int{0}
}

type PublicParamsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublicParamsReq) Reset() {
	*x = PublicParamsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicParamsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicParamsReq) ProtoMessage() {}

func (x *PublicParamsReq) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicParamsReq.ProtoReflect.Descriptor instead.
func (*PublicParamsReq) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{1}
}

type PublicParamsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q   []byte   `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`     // Store q as bytes
	G   []byte   `protobuf:"bytes,2,opt,name=g,proto3" json:"g,omitempty"`     // Store g as bytes
	Mpk [][]byte `protobuf:"bytes,3,rep,name=mpk,proto3" json:"mpk,omitempty"` // Store mpk as bytes
}

func (x *PublicParamsRes) Reset() {
	*x = PublicParamsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicParamsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicParamsRes) ProtoMessage() {}

func (x *PublicParamsRes) ProtoReflect() protoreflect.Message {
	mi := &file_spade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicParamsRes.ProtoReflect.Descriptor instead.
func (*PublicParamsRes) Descriptor() ([]byte, []int) {
	return file_spade_proto_rawDescGZIP(), []int{2}
}

func (x *PublicParamsRes) GetQ() []byte {
	if x != nil {
		return x.Q
	}
	return nil
}

func (x *PublicParamsRes) GetG() []byte {
	if x != nil {
		return x.G
	}
	return nil
}

func (x *PublicParamsRes) GetMpk() [][]byte {
	if x != nil {
		return x.Mpk
	}
	return nil
}

var File_spade_proto protoreflect.FileDescriptor

var file_spade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x3f, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x01, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x01, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x03, 0x6d, 0x70, 0x6b, 0x32, 0x58, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spade_proto_rawDescOnce sync.Once
	file_spade_proto_rawDescData = file_spade_proto_rawDesc
)

func file_spade_proto_rawDescGZIP() []byte {
	file_spade_proto_rawDescOnce.Do(func() {
		file_spade_proto_rawDescData = protoimpl.X.CompressGZIP(file_spade_proto_rawDescData)
	})
	return file_spade_proto_rawDescData
}

var file_spade_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spade_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: spadeproto.Empty
	(*PublicParamsReq)(nil), // 1: spadeproto.PublicParamsReq
	(*PublicParamsRes)(nil), // 2: spadeproto.PublicParamsRes
}
var file_spade_proto_depIdxs = []int32{
	1, // 0: spadeproto.Curator.GetPublicParams:input_type -> spadeproto.PublicParamsReq
	2, // 1: spadeproto.Curator.GetPublicParams:output_type -> spadeproto.PublicParamsRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spade_proto_init() }
func file_spade_proto_init() {
	if File_spade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_spade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicParamsReq); i {
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
		file_spade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicParamsRes); i {
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
			RawDescriptor: file_spade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spade_proto_goTypes,
		DependencyIndexes: file_spade_proto_depIdxs,
		MessageInfos:      file_spade_proto_msgTypes,
	}.Build()
	File_spade_proto = out.File
	file_spade_proto_rawDesc = nil
	file_spade_proto_goTypes = nil
	file_spade_proto_depIdxs = nil
}
