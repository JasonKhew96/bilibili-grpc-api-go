// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/app/interfaces/v1/space.proto

package v1

import (
	v11 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/archive/middleware/v1"
	v1 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/archive/v1"
	v2 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/dynamic/v2"
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

type From int32

const (
	From_ArchiveTab From = 0 //
	From_DynamicTab From = 1 //
)

// Enum value maps for From.
var (
	From_name = map[int32]string{
		0: "ArchiveTab",
		1: "DynamicTab",
	}
	From_value = map[string]int32{
		"ArchiveTab": 0,
		"DynamicTab": 1,
	}
)

func (x From) Enum() *From {
	p := new(From)
	*p = x
	return p
}

func (x From) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (From) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_interfaces_v1_space_proto_enumTypes[0].Descriptor()
}

func (From) Type() protoreflect.EnumType {
	return &file_bilibili_app_interfaces_v1_space_proto_enumTypes[0]
}

func (x From) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use From.Descriptor instead.
func (From) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{0}
}

type Arc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Archive *v1.Arc `protobuf:"bytes,1,opt,name=archive,proto3" json:"archive,omitempty"`
	Uri     string  `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Arc) Reset() {
	*x = Arc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arc) ProtoMessage() {}

func (x *Arc) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Arc.ProtoReflect.Descriptor instead.
func (*Arc) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{0}
}

func (x *Arc) GetArchive() *v1.Arc {
	if x != nil {
		return x.Archive
	}
	return nil
}

func (x *Arc) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type Dynamic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dynamic *v2.DynamicItem `protobuf:"bytes,1,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
}

func (x *Dynamic) Reset() {
	*x = Dynamic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dynamic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dynamic) ProtoMessage() {}

func (x *Dynamic) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dynamic.ProtoReflect.Descriptor instead.
func (*Dynamic) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{1}
}

func (x *Dynamic) GetDynamic() *v2.DynamicItem {
	if x != nil {
		return x.Dynamic
	}
	return nil
}

type SearchTabReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Focus int64  `protobuf:"varint,1,opt,name=focus,proto3" json:"focus,omitempty"`
	Tabs  []*Tab `protobuf:"bytes,2,rep,name=tabs,proto3" json:"tabs,omitempty"`
}

func (x *SearchTabReply) Reset() {
	*x = SearchTabReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTabReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTabReply) ProtoMessage() {}

func (x *SearchTabReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTabReply.ProtoReflect.Descriptor instead.
func (*SearchTabReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{2}
}

func (x *SearchTabReply) GetFocus() int64 {
	if x != nil {
		return x.Focus
	}
	return 0
}

func (x *SearchTabReply) GetTabs() []*Tab {
	if x != nil {
		return x.Tabs
	}
	return nil
}

type SearchTabReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Mid     int64  `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	From    int32  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *SearchTabReq) Reset() {
	*x = SearchTabReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTabReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTabReq) ProtoMessage() {}

func (x *SearchTabReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTabReq.ProtoReflect.Descriptor instead.
func (*SearchTabReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{3}
}

func (x *SearchTabReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchTabReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *SearchTabReq) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

type SearchArchiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Archives []*Arc `protobuf:"bytes,1,rep,name=archives,proto3" json:"archives,omitempty"`
	Total    int64  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SearchArchiveReply) Reset() {
	*x = SearchArchiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArchiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArchiveReply) ProtoMessage() {}

func (x *SearchArchiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArchiveReply.ProtoReflect.Descriptor instead.
func (*SearchArchiveReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{4}
}

func (x *SearchArchiveReply) GetArchives() []*Arc {
	if x != nil {
		return x.Archives
	}
	return nil
}

func (x *SearchArchiveReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SearchArchiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string          `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Mid        int64           `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Pn         int64           `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps         int64           `protobuf:"varint,4,opt,name=ps,proto3" json:"ps,omitempty"`
	PlayerArgs *v11.PlayerArgs `protobuf:"bytes,5,opt,name=player_args,json=playerArgs,proto3" json:"player_args,omitempty"`
}

func (x *SearchArchiveReq) Reset() {
	*x = SearchArchiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArchiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArchiveReq) ProtoMessage() {}

func (x *SearchArchiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArchiveReq.ProtoReflect.Descriptor instead.
func (*SearchArchiveReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{5}
}

func (x *SearchArchiveReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchArchiveReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *SearchArchiveReq) GetPn() int64 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *SearchArchiveReq) GetPs() int64 {
	if x != nil {
		return x.Ps
	}
	return 0
}

func (x *SearchArchiveReq) GetPlayerArgs() *v11.PlayerArgs {
	if x != nil {
		return x.PlayerArgs
	}
	return nil
}

type SearchDynamicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dynamics []*Dynamic `protobuf:"bytes,1,rep,name=dynamics,proto3" json:"dynamics,omitempty"`
	Total    int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SearchDynamicReply) Reset() {
	*x = SearchDynamicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDynamicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDynamicReply) ProtoMessage() {}

func (x *SearchDynamicReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDynamicReply.ProtoReflect.Descriptor instead.
func (*SearchDynamicReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{6}
}

func (x *SearchDynamicReply) GetDynamics() []*Dynamic {
	if x != nil {
		return x.Dynamics
	}
	return nil
}

func (x *SearchDynamicReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SearchDynamicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string          `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Mid        int64           `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Pn         int64           `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps         int64           `protobuf:"varint,4,opt,name=ps,proto3" json:"ps,omitempty"`
	PlayerArgs *v11.PlayerArgs `protobuf:"bytes,5,opt,name=player_args,json=playerArgs,proto3" json:"player_args,omitempty"`
}

func (x *SearchDynamicReq) Reset() {
	*x = SearchDynamicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDynamicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDynamicReq) ProtoMessage() {}

func (x *SearchDynamicReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDynamicReq.ProtoReflect.Descriptor instead.
func (*SearchDynamicReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{7}
}

func (x *SearchDynamicReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchDynamicReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *SearchDynamicReq) GetPn() int64 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *SearchDynamicReq) GetPs() int64 {
	if x != nil {
		return x.Ps
	}
	return 0
}

func (x *SearchDynamicReq) GetPlayerArgs() *v11.PlayerArgs {
	if x != nil {
		return x.PlayerArgs
	}
	return nil
}

type Tab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Uri   string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Tab) Reset() {
	*x = Tab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tab) ProtoMessage() {}

func (x *Tab) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_interfaces_v1_space_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tab.ProtoReflect.Descriptor instead.
func (*Tab) Descriptor() ([]byte, []int) {
	return file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP(), []int{8}
}

func (x *Tab) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Tab) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_bilibili_app_interfaces_v1_space_proto protoreflect.FileDescriptor

var file_bilibili_app_interfaces_v1_space_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x30, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x03, 0x41, 0x72, 0x63, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63, 0x52, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x22, 0x49, 0x0a, 0x07, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12,
	0x3e, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x22,
	0x5a, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61, 0x62, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x62, 0x52, 0x04, 0x74, 0x61, 0x62, 0x73, 0x22, 0x4e, 0x0a, 0x0c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x62, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x66, 0x0a, 0x12, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x63, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x70, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x70, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x22, 0x6a, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x52, 0x08, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x70, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x70, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41,
	0x72, 0x67, 0x73, 0x22, 0x2d, 0x0a, 0x03, 0x54, 0x61, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x2a, 0x26, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x61, 0x62, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x54, 0x61, 0x62, 0x10, 0x01, 0x32, 0xc2, 0x02, 0x0a, 0x05, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61,
	0x62, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6b, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x6b, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x1a, 0x2d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69,
	0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_interfaces_v1_space_proto_rawDescOnce sync.Once
	file_bilibili_app_interfaces_v1_space_proto_rawDescData = file_bilibili_app_interfaces_v1_space_proto_rawDesc
)

func file_bilibili_app_interfaces_v1_space_proto_rawDescGZIP() []byte {
	file_bilibili_app_interfaces_v1_space_proto_rawDescOnce.Do(func() {
		file_bilibili_app_interfaces_v1_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_interfaces_v1_space_proto_rawDescData)
	})
	return file_bilibili_app_interfaces_v1_space_proto_rawDescData
}

var file_bilibili_app_interfaces_v1_space_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bilibili_app_interfaces_v1_space_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bilibili_app_interfaces_v1_space_proto_goTypes = []interface{}{
	(From)(0),                  // 0: bilibili.app.interface.v1.From
	(*Arc)(nil),                // 1: bilibili.app.interface.v1.Arc
	(*Dynamic)(nil),            // 2: bilibili.app.interface.v1.Dynamic
	(*SearchTabReply)(nil),     // 3: bilibili.app.interface.v1.SearchTabReply
	(*SearchTabReq)(nil),       // 4: bilibili.app.interface.v1.SearchTabReq
	(*SearchArchiveReply)(nil), // 5: bilibili.app.interface.v1.SearchArchiveReply
	(*SearchArchiveReq)(nil),   // 6: bilibili.app.interface.v1.SearchArchiveReq
	(*SearchDynamicReply)(nil), // 7: bilibili.app.interface.v1.SearchDynamicReply
	(*SearchDynamicReq)(nil),   // 8: bilibili.app.interface.v1.SearchDynamicReq
	(*Tab)(nil),                // 9: bilibili.app.interface.v1.Tab
	(*v1.Arc)(nil),             // 10: bilibili.app.archive.v1.Arc
	(*v2.DynamicItem)(nil),     // 11: bilibili.app.dynamic.v2.DynamicItem
	(*v11.PlayerArgs)(nil),     // 12: bilibili.app.archive.middleware.v1.PlayerArgs
}
var file_bilibili_app_interfaces_v1_space_proto_depIdxs = []int32{
	10, // 0: bilibili.app.interface.v1.Arc.archive:type_name -> bilibili.app.archive.v1.Arc
	11, // 1: bilibili.app.interface.v1.Dynamic.dynamic:type_name -> bilibili.app.dynamic.v2.DynamicItem
	9,  // 2: bilibili.app.interface.v1.SearchTabReply.tabs:type_name -> bilibili.app.interface.v1.Tab
	1,  // 3: bilibili.app.interface.v1.SearchArchiveReply.archives:type_name -> bilibili.app.interface.v1.Arc
	12, // 4: bilibili.app.interface.v1.SearchArchiveReq.player_args:type_name -> bilibili.app.archive.middleware.v1.PlayerArgs
	2,  // 5: bilibili.app.interface.v1.SearchDynamicReply.dynamics:type_name -> bilibili.app.interface.v1.Dynamic
	12, // 6: bilibili.app.interface.v1.SearchDynamicReq.player_args:type_name -> bilibili.app.archive.middleware.v1.PlayerArgs
	4,  // 7: bilibili.app.interface.v1.Space.SearchTab:input_type -> bilibili.app.interface.v1.SearchTabReq
	6,  // 8: bilibili.app.interface.v1.Space.SearchArchive:input_type -> bilibili.app.interface.v1.SearchArchiveReq
	8,  // 9: bilibili.app.interface.v1.Space.SearchDynamic:input_type -> bilibili.app.interface.v1.SearchDynamicReq
	3,  // 10: bilibili.app.interface.v1.Space.SearchTab:output_type -> bilibili.app.interface.v1.SearchTabReply
	5,  // 11: bilibili.app.interface.v1.Space.SearchArchive:output_type -> bilibili.app.interface.v1.SearchArchiveReply
	7,  // 12: bilibili.app.interface.v1.Space.SearchDynamic:output_type -> bilibili.app.interface.v1.SearchDynamicReply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_bilibili_app_interfaces_v1_space_proto_init() }
func file_bilibili_app_interfaces_v1_space_proto_init() {
	if File_bilibili_app_interfaces_v1_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Arc); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dynamic); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTabReply); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTabReq); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArchiveReply); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArchiveReq); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDynamicReply); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDynamicReq); i {
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
		file_bilibili_app_interfaces_v1_space_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tab); i {
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
			RawDescriptor: file_bilibili_app_interfaces_v1_space_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_interfaces_v1_space_proto_goTypes,
		DependencyIndexes: file_bilibili_app_interfaces_v1_space_proto_depIdxs,
		EnumInfos:         file_bilibili_app_interfaces_v1_space_proto_enumTypes,
		MessageInfos:      file_bilibili_app_interfaces_v1_space_proto_msgTypes,
	}.Build()
	File_bilibili_app_interfaces_v1_space_proto = out.File
	file_bilibili_app_interfaces_v1_space_proto_rawDesc = nil
	file_bilibili_app_interfaces_v1_space_proto_goTypes = nil
	file_bilibili_app_interfaces_v1_space_proto_depIdxs = nil
}
