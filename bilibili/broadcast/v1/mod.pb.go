// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/broadcast/v1/mod.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ModResourceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atcion        int32  `protobuf:"varint,1,opt,name=atcion,proto3" json:"atcion,omitempty"`
	AppKey        string `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	PoolName      string `protobuf:"bytes,3,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	ModuleName    string `protobuf:"bytes,4,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ModuleVersion int64  `protobuf:"varint,5,opt,name=module_version,json=moduleVersion,proto3" json:"module_version,omitempty"`
	ListVersion   int64  `protobuf:"varint,6,opt,name=list_version,json=listVersion,proto3" json:"list_version,omitempty"`
}

func (x *ModResourceResp) Reset() {
	*x = ModResourceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_mod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModResourceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModResourceResp) ProtoMessage() {}

func (x *ModResourceResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_mod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModResourceResp.ProtoReflect.Descriptor instead.
func (*ModResourceResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_mod_proto_rawDescGZIP(), []int{0}
}

func (x *ModResourceResp) GetAtcion() int32 {
	if x != nil {
		return x.Atcion
	}
	return 0
}

func (x *ModResourceResp) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *ModResourceResp) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

func (x *ModResourceResp) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *ModResourceResp) GetModuleVersion() int64 {
	if x != nil {
		return x.ModuleVersion
	}
	return 0
}

func (x *ModResourceResp) GetListVersion() int64 {
	if x != nil {
		return x.ListVersion
	}
	return 0
}

var File_bilibili_broadcast_v1_mod_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_v1_mod_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x63,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x74, 0x63, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0x5f, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_v1_mod_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_v1_mod_proto_rawDescData = file_bilibili_broadcast_v1_mod_proto_rawDesc
)

func file_bilibili_broadcast_v1_mod_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_v1_mod_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_v1_mod_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_v1_mod_proto_rawDescData)
	})
	return file_bilibili_broadcast_v1_mod_proto_rawDescData
}

var file_bilibili_broadcast_v1_mod_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_broadcast_v1_mod_proto_goTypes = []interface{}{
	(*ModResourceResp)(nil), // 0: bilibili.broadcast.v1.ModResourceResp
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
}
var file_bilibili_broadcast_v1_mod_proto_depIdxs = []int32{
	1, // 0: bilibili.broadcast.v1.ModManager.WatchResource:input_type -> google.protobuf.Empty
	0, // 1: bilibili.broadcast.v1.ModManager.WatchResource:output_type -> bilibili.broadcast.v1.ModResourceResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_v1_mod_proto_init() }
func file_bilibili_broadcast_v1_mod_proto_init() {
	if File_bilibili_broadcast_v1_mod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_v1_mod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModResourceResp); i {
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
			RawDescriptor: file_bilibili_broadcast_v1_mod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_v1_mod_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_v1_mod_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_v1_mod_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_v1_mod_proto = out.File
	file_bilibili_broadcast_v1_mod_proto_rawDesc = nil
	file_bilibili_broadcast_v1_mod_proto_goTypes = nil
	file_bilibili_broadcast_v1_mod_proto_depIdxs = nil
}
