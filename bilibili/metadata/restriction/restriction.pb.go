// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/metadata/restriction/restriction.proto

package restriction

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

// 模式类型
type ModeType int32

const (
	ModeType_NORMAL    ModeType = 0 // 正常模式
	ModeType_TEENAGERS ModeType = 1 // 青少年模式
	ModeType_LESSONS   ModeType = 2 // 课堂模式
)

// Enum value maps for ModeType.
var (
	ModeType_name = map[int32]string{
		0: "NORMAL",
		1: "TEENAGERS",
		2: "LESSONS",
	}
	ModeType_value = map[string]int32{
		"NORMAL":    0,
		"TEENAGERS": 1,
		"LESSONS":   2,
	}
)

func (x ModeType) Enum() *ModeType {
	p := new(ModeType)
	*p = x
	return p
}

func (x ModeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModeType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_metadata_restriction_restriction_proto_enumTypes[0].Descriptor()
}

func (ModeType) Type() protoreflect.EnumType {
	return &file_bilibili_metadata_restriction_restriction_proto_enumTypes[0]
}

func (x ModeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModeType.Descriptor instead.
func (ModeType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_metadata_restriction_restriction_proto_rawDescGZIP(), []int{0}
}

// 限制条件
type Restriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 青少年模式开关状态
	TeenagersMode bool `protobuf:"varint,1,opt,name=teenagers_mode,json=teenagersMode,proto3" json:"teenagers_mode,omitempty"`
	// 课堂模式开关状态
	LessonsMode bool `protobuf:"varint,2,opt,name=lessons_mode,json=lessonsMode,proto3" json:"lessons_mode,omitempty"`
	// 模式类型(旧版)
	Mode ModeType `protobuf:"varint,3,opt,name=mode,proto3,enum=bilibili.metadata.restriction.ModeType" json:"mode,omitempty"`
	// app 审核review状态
	Review      bool `protobuf:"varint,4,opt,name=review,proto3" json:"review,omitempty"`
	DisableRcmd bool `protobuf:"varint,5,opt,name=disable_rcmd,json=disableRcmd,proto3" json:"disable_rcmd,omitempty"`
}

func (x *Restriction) Reset() {
	*x = Restriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_metadata_restriction_restriction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restriction) ProtoMessage() {}

func (x *Restriction) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_metadata_restriction_restriction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restriction.ProtoReflect.Descriptor instead.
func (*Restriction) Descriptor() ([]byte, []int) {
	return file_bilibili_metadata_restriction_restriction_proto_rawDescGZIP(), []int{0}
}

func (x *Restriction) GetTeenagersMode() bool {
	if x != nil {
		return x.TeenagersMode
	}
	return false
}

func (x *Restriction) GetLessonsMode() bool {
	if x != nil {
		return x.LessonsMode
	}
	return false
}

func (x *Restriction) GetMode() ModeType {
	if x != nil {
		return x.Mode
	}
	return ModeType_NORMAL
}

func (x *Restriction) GetReview() bool {
	if x != nil {
		return x.Review
	}
	return false
}

func (x *Restriction) GetDisableRcmd() bool {
	if x != nil {
		return x.DisableRcmd
	}
	return false
}

var File_bilibili_metadata_restriction_restriction_proto protoreflect.FileDescriptor

var file_bilibili_metadata_restriction_restriction_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x65, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x65, 0x65, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x63, 0x6d, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x63,
	0x6d, 0x64, 0x2a, 0x32, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45,
	0x45, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x45, 0x53,
	0x53, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_metadata_restriction_restriction_proto_rawDescOnce sync.Once
	file_bilibili_metadata_restriction_restriction_proto_rawDescData = file_bilibili_metadata_restriction_restriction_proto_rawDesc
)

func file_bilibili_metadata_restriction_restriction_proto_rawDescGZIP() []byte {
	file_bilibili_metadata_restriction_restriction_proto_rawDescOnce.Do(func() {
		file_bilibili_metadata_restriction_restriction_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_metadata_restriction_restriction_proto_rawDescData)
	})
	return file_bilibili_metadata_restriction_restriction_proto_rawDescData
}

var file_bilibili_metadata_restriction_restriction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bilibili_metadata_restriction_restriction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_metadata_restriction_restriction_proto_goTypes = []interface{}{
	(ModeType)(0),       // 0: bilibili.metadata.restriction.ModeType
	(*Restriction)(nil), // 1: bilibili.metadata.restriction.Restriction
}
var file_bilibili_metadata_restriction_restriction_proto_depIdxs = []int32{
	0, // 0: bilibili.metadata.restriction.Restriction.mode:type_name -> bilibili.metadata.restriction.ModeType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_metadata_restriction_restriction_proto_init() }
func file_bilibili_metadata_restriction_restriction_proto_init() {
	if File_bilibili_metadata_restriction_restriction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_metadata_restriction_restriction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restriction); i {
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
			RawDescriptor: file_bilibili_metadata_restriction_restriction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_metadata_restriction_restriction_proto_goTypes,
		DependencyIndexes: file_bilibili_metadata_restriction_restriction_proto_depIdxs,
		EnumInfos:         file_bilibili_metadata_restriction_restriction_proto_enumTypes,
		MessageInfos:      file_bilibili_metadata_restriction_restriction_proto_msgTypes,
	}.Build()
	File_bilibili_metadata_restriction_restriction_proto = out.File
	file_bilibili_metadata_restriction_restriction_proto_rawDesc = nil
	file_bilibili_metadata_restriction_restriction_proto_goTypes = nil
	file_bilibili_metadata_restriction_restriction_proto_depIdxs = nil
}
