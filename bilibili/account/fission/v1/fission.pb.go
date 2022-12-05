// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/account/fission/v1/fission.proto

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

// 动画效果
type AnimateIcon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// icon文件
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// 动效json文件
	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *AnimateIcon) Reset() {
	*x = AnimateIcon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimateIcon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimateIcon) ProtoMessage() {}

func (x *AnimateIcon) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimateIcon.ProtoReflect.Descriptor instead.
func (*AnimateIcon) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{0}
}

func (x *AnimateIcon) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AnimateIcon) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

// 活动入口-响应
type EntranceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 展示图标
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// 活动名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 活动跳转链接
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// 动画效果
	AnimateIcon *AnimateIcon `protobuf:"bytes,4,opt,name=animate_icon,json=animateIcon,proto3" json:"animate_icon,omitempty"`
}

func (x *EntranceReply) Reset() {
	*x = EntranceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntranceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntranceReply) ProtoMessage() {}

func (x *EntranceReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntranceReply.ProtoReflect.Descriptor instead.
func (*EntranceReply) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{1}
}

func (x *EntranceReply) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *EntranceReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EntranceReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *EntranceReply) GetAnimateIcon() *AnimateIcon {
	if x != nil {
		return x.AnimateIcon
	}
	return nil
}

// 活动入口-请求
type EntranceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntranceReq) Reset() {
	*x = EntranceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntranceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntranceReq) ProtoMessage() {}

func (x *EntranceReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntranceReq.ProtoReflect.Descriptor instead.
func (*EntranceReq) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{2}
}

type PrivacyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PrivacyReply) Reset() {
	*x = PrivacyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyReply) ProtoMessage() {}

func (x *PrivacyReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyReply.ProtoReflect.Descriptor instead.
func (*PrivacyReply) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{3}
}

func (x *PrivacyReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PrivacyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityUid string `protobuf:"bytes,1,opt,name=activity_uid,json=activityUid,proto3" json:"activity_uid,omitempty"`
}

func (x *PrivacyReq) Reset() {
	*x = PrivacyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyReq) ProtoMessage() {}

func (x *PrivacyReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyReq.ProtoReflect.Descriptor instead.
func (*PrivacyReq) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{4}
}

func (x *PrivacyReq) GetActivityUid() string {
	if x != nil {
		return x.ActivityUid
	}
	return ""
}

// 首页弹窗-响应
type WindowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 弹窗类型
	// 0:弹窗 1:普通页面
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 跳转链接
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// 上报数据字段
	ReportData string `protobuf:"bytes,3,opt,name=report_data,json=reportData,proto3" json:"report_data,omitempty"`
}

func (x *WindowReply) Reset() {
	*x = WindowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowReply) ProtoMessage() {}

func (x *WindowReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowReply.ProtoReflect.Descriptor instead.
func (*WindowReply) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{5}
}

func (x *WindowReply) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *WindowReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WindowReply) GetReportData() string {
	if x != nil {
		return x.ReportData
	}
	return ""
}

// 首页弹窗-请求
type WindowReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WindowReq) Reset() {
	*x = WindowReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowReq) ProtoMessage() {}

func (x *WindowReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_account_fission_v1_fission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowReq.ProtoReflect.Descriptor instead.
func (*WindowReq) Descriptor() ([]byte, []int) {
	return file_bilibili_account_fission_v1_fission_proto_rawDescGZIP(), []int{6}
}

var File_bilibili_account_fission_v1_fission_proto protoreflect.FileDescriptor

var file_bilibili_account_fission_v1_fission_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x0b, 0x41, 0x6e, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22,
	0x96, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x4b, 0x0a, 0x0c, 0x61,
	0x6e, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2f, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x55,
	0x69, 0x64, 0x22, 0x54, 0x0a, 0x0b, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0b, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x32, 0xa6, 0x02, 0x0a, 0x07, 0x46, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x60, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x26, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x5d, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x66, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x48,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61,
	0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x66, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_account_fission_v1_fission_proto_rawDescOnce sync.Once
	file_bilibili_account_fission_v1_fission_proto_rawDescData = file_bilibili_account_fission_v1_fission_proto_rawDesc
)

func file_bilibili_account_fission_v1_fission_proto_rawDescGZIP() []byte {
	file_bilibili_account_fission_v1_fission_proto_rawDescOnce.Do(func() {
		file_bilibili_account_fission_v1_fission_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_account_fission_v1_fission_proto_rawDescData)
	})
	return file_bilibili_account_fission_v1_fission_proto_rawDescData
}

var file_bilibili_account_fission_v1_fission_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bilibili_account_fission_v1_fission_proto_goTypes = []interface{}{
	(*AnimateIcon)(nil),   // 0: bilibili.account.fission.v1.AnimateIcon
	(*EntranceReply)(nil), // 1: bilibili.account.fission.v1.EntranceReply
	(*EntranceReq)(nil),   // 2: bilibili.account.fission.v1.EntranceReq
	(*PrivacyReply)(nil),  // 3: bilibili.account.fission.v1.PrivacyReply
	(*PrivacyReq)(nil),    // 4: bilibili.account.fission.v1.PrivacyReq
	(*WindowReply)(nil),   // 5: bilibili.account.fission.v1.WindowReply
	(*WindowReq)(nil),     // 6: bilibili.account.fission.v1.WindowReq
}
var file_bilibili_account_fission_v1_fission_proto_depIdxs = []int32{
	0, // 0: bilibili.account.fission.v1.EntranceReply.animate_icon:type_name -> bilibili.account.fission.v1.AnimateIcon
	2, // 1: bilibili.account.fission.v1.Fission.Entrance:input_type -> bilibili.account.fission.v1.EntranceReq
	6, // 2: bilibili.account.fission.v1.Fission.Window:input_type -> bilibili.account.fission.v1.WindowReq
	4, // 3: bilibili.account.fission.v1.Fission.Privacy:input_type -> bilibili.account.fission.v1.PrivacyReq
	1, // 4: bilibili.account.fission.v1.Fission.Entrance:output_type -> bilibili.account.fission.v1.EntranceReply
	5, // 5: bilibili.account.fission.v1.Fission.Window:output_type -> bilibili.account.fission.v1.WindowReply
	3, // 6: bilibili.account.fission.v1.Fission.Privacy:output_type -> bilibili.account.fission.v1.PrivacyReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_account_fission_v1_fission_proto_init() }
func file_bilibili_account_fission_v1_fission_proto_init() {
	if File_bilibili_account_fission_v1_fission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_account_fission_v1_fission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimateIcon); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntranceReply); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntranceReq); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyReply); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyReq); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowReply); i {
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
		file_bilibili_account_fission_v1_fission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowReq); i {
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
			RawDescriptor: file_bilibili_account_fission_v1_fission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_account_fission_v1_fission_proto_goTypes,
		DependencyIndexes: file_bilibili_account_fission_v1_fission_proto_depIdxs,
		MessageInfos:      file_bilibili_account_fission_v1_fission_proto_msgTypes,
	}.Build()
	File_bilibili_account_fission_v1_fission_proto = out.File
	file_bilibili_account_fission_v1_fission_proto_rawDesc = nil
	file_bilibili_account_fission_v1_fission_proto_goTypes = nil
	file_bilibili_account_fission_v1_fission_proto_depIdxs = nil
}
