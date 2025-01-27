// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/app/card/v1/ad.proto

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

type AdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreativeId      int64            `protobuf:"varint,1,opt,name=creative_id,json=creativeId,proto3" json:"creative_id,omitempty"`
	CreativeType    int32            `protobuf:"varint,2,opt,name=creative_type,json=creativeType,proto3" json:"creative_type,omitempty"`
	CardType        int32            `protobuf:"varint,3,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	CreativeContent *CreativeContent `protobuf:"bytes,4,opt,name=creative_content,json=creativeContent,proto3" json:"creative_content,omitempty"`
	AdCb            string           `protobuf:"bytes,5,opt,name=ad_cb,json=adCb,proto3" json:"ad_cb,omitempty"`
	Resource        int64            `protobuf:"varint,6,opt,name=resource,proto3" json:"resource,omitempty"`
	Source          int32            `protobuf:"varint,7,opt,name=source,proto3" json:"source,omitempty"`
	RequestId       string           `protobuf:"bytes,8,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	IsAd            bool             `protobuf:"varint,9,opt,name=is_ad,json=isAd,proto3" json:"is_ad,omitempty"`
	CmMark          int64            `protobuf:"varint,10,opt,name=cm_mark,json=cmMark,proto3" json:"cm_mark,omitempty"`
	Index           int32            `protobuf:"varint,11,opt,name=index,proto3" json:"index,omitempty"`
	IsAdLoc         bool             `protobuf:"varint,12,opt,name=is_ad_loc,json=isAdLoc,proto3" json:"is_ad_loc,omitempty"`
	CardIndex       int32            `protobuf:"varint,13,opt,name=card_index,json=cardIndex,proto3" json:"card_index,omitempty"`
	ClientIp        string           `protobuf:"bytes,14,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Extra           []byte           `protobuf:"bytes,15,opt,name=extra,proto3" json:"extra,omitempty"`
	CreativeStyle   int32            `protobuf:"varint,16,opt,name=creative_style,json=creativeStyle,proto3" json:"creative_style,omitempty"`
}

func (x *AdInfo) Reset() {
	*x = AdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_card_v1_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfo) ProtoMessage() {}

func (x *AdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_card_v1_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfo.ProtoReflect.Descriptor instead.
func (*AdInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_app_card_v1_ad_proto_rawDescGZIP(), []int{0}
}

func (x *AdInfo) GetCreativeId() int64 {
	if x != nil {
		return x.CreativeId
	}
	return 0
}

func (x *AdInfo) GetCreativeType() int32 {
	if x != nil {
		return x.CreativeType
	}
	return 0
}

func (x *AdInfo) GetCardType() int32 {
	if x != nil {
		return x.CardType
	}
	return 0
}

func (x *AdInfo) GetCreativeContent() *CreativeContent {
	if x != nil {
		return x.CreativeContent
	}
	return nil
}

func (x *AdInfo) GetAdCb() string {
	if x != nil {
		return x.AdCb
	}
	return ""
}

func (x *AdInfo) GetResource() int64 {
	if x != nil {
		return x.Resource
	}
	return 0
}

func (x *AdInfo) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *AdInfo) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AdInfo) GetIsAd() bool {
	if x != nil {
		return x.IsAd
	}
	return false
}

func (x *AdInfo) GetCmMark() int64 {
	if x != nil {
		return x.CmMark
	}
	return 0
}

func (x *AdInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *AdInfo) GetIsAdLoc() bool {
	if x != nil {
		return x.IsAdLoc
	}
	return false
}

func (x *AdInfo) GetCardIndex() int32 {
	if x != nil {
		return x.CardIndex
	}
	return 0
}

func (x *AdInfo) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *AdInfo) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *AdInfo) GetCreativeStyle() int32 {
	if x != nil {
		return x.CreativeStyle
	}
	return 0
}

type CreativeContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	VideoId     int64  `protobuf:"varint,3,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	ImageUrl    string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	ImageMd5    string `protobuf:"bytes,6,opt,name=image_md5,json=imageMd5,proto3" json:"image_md5,omitempty"`
	LogUrl      string `protobuf:"bytes,7,opt,name=log_url,json=logUrl,proto3" json:"log_url,omitempty"`
	LogMd5      string `protobuf:"bytes,8,opt,name=log_md5,json=logMd5,proto3" json:"log_md5,omitempty"`
	Url         string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	ClickUrl    string `protobuf:"bytes,10,opt,name=click_url,json=clickUrl,proto3" json:"click_url,omitempty"`
	ShowUrl     string `protobuf:"bytes,11,opt,name=show_url,json=showUrl,proto3" json:"show_url,omitempty"`
}

func (x *CreativeContent) Reset() {
	*x = CreativeContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_card_v1_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreativeContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreativeContent) ProtoMessage() {}

func (x *CreativeContent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_card_v1_ad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreativeContent.ProtoReflect.Descriptor instead.
func (*CreativeContent) Descriptor() ([]byte, []int) {
	return file_bilibili_app_card_v1_ad_proto_rawDescGZIP(), []int{1}
}

func (x *CreativeContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreativeContent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreativeContent) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CreativeContent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreativeContent) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreativeContent) GetImageMd5() string {
	if x != nil {
		return x.ImageMd5
	}
	return ""
}

func (x *CreativeContent) GetLogUrl() string {
	if x != nil {
		return x.LogUrl
	}
	return ""
}

func (x *CreativeContent) GetLogMd5() string {
	if x != nil {
		return x.LogMd5
	}
	return ""
}

func (x *CreativeContent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreativeContent) GetClickUrl() string {
	if x != nil {
		return x.ClickUrl
	}
	return ""
}

func (x *CreativeContent) GetShowUrl() string {
	if x != nil {
		return x.ShowUrl
	}
	return ""
}

var File_bilibili_app_card_v1_ad_proto protoreflect.FileDescriptor

var file_bilibili_app_card_v1_ad_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63,
	0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xfe, 0x03, 0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x63, 0x62, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x43, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x13, 0x0a,
	0x05, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73,
	0x41, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6d, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x4c, 0x6f, 0x63, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x22, 0xb6, 0x02, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4d, 0x64, 0x35, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x6f, 0x67, 0x4d, 0x64, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x55, 0x72, 0x6c, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69,
	0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_card_v1_ad_proto_rawDescOnce sync.Once
	file_bilibili_app_card_v1_ad_proto_rawDescData = file_bilibili_app_card_v1_ad_proto_rawDesc
)

func file_bilibili_app_card_v1_ad_proto_rawDescGZIP() []byte {
	file_bilibili_app_card_v1_ad_proto_rawDescOnce.Do(func() {
		file_bilibili_app_card_v1_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_card_v1_ad_proto_rawDescData)
	})
	return file_bilibili_app_card_v1_ad_proto_rawDescData
}

var file_bilibili_app_card_v1_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_app_card_v1_ad_proto_goTypes = []interface{}{
	(*AdInfo)(nil),          // 0: bilibili.app.card.v1.AdInfo
	(*CreativeContent)(nil), // 1: bilibili.app.card.v1.CreativeContent
}
var file_bilibili_app_card_v1_ad_proto_depIdxs = []int32{
	1, // 0: bilibili.app.card.v1.AdInfo.creative_content:type_name -> bilibili.app.card.v1.CreativeContent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_app_card_v1_ad_proto_init() }
func file_bilibili_app_card_v1_ad_proto_init() {
	if File_bilibili_app_card_v1_ad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_card_v1_ad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfo); i {
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
		file_bilibili_app_card_v1_ad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreativeContent); i {
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
			RawDescriptor: file_bilibili_app_card_v1_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_card_v1_ad_proto_goTypes,
		DependencyIndexes: file_bilibili_app_card_v1_ad_proto_depIdxs,
		MessageInfos:      file_bilibili_app_card_v1_ad_proto_msgTypes,
	}.Build()
	File_bilibili_app_card_v1_ad_proto = out.File
	file_bilibili_app_card_v1_ad_proto_rawDesc = nil
	file_bilibili_app_card_v1_ad_proto_goTypes = nil
	file_bilibili_app_card_v1_ad_proto_depIdxs = nil
}
