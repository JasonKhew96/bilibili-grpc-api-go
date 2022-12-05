// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/community/service/govern/v1/govern.proto

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

type QoeReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Scene        int64           `protobuf:"varint,2,opt,name=scene,proto3" json:"scene,omitempty"`
	Type         int32           `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Cancel       bool            `protobuf:"varint,4,opt,name=cancel,proto3" json:"cancel,omitempty"`
	BusinessType string          `protobuf:"bytes,5,opt,name=business_type,json=businessType,proto3" json:"business_type,omitempty"`
	Oid          int64           `protobuf:"varint,6,opt,name=oid,proto3" json:"oid,omitempty"`
	ScoreResult  *QoeScoreResult `protobuf:"bytes,7,opt,name=score_result,json=scoreResult,proto3" json:"score_result,omitempty"`
	BusinessData string          `protobuf:"bytes,8,opt,name=business_data,json=businessData,proto3" json:"business_data,omitempty"`
}

func (x *QoeReportReq) Reset() {
	*x = QoeReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_community_service_govern_v1_govern_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QoeReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QoeReportReq) ProtoMessage() {}

func (x *QoeReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_community_service_govern_v1_govern_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QoeReportReq.ProtoReflect.Descriptor instead.
func (*QoeReportReq) Descriptor() ([]byte, []int) {
	return file_bilibili_community_service_govern_v1_govern_proto_rawDescGZIP(), []int{0}
}

func (x *QoeReportReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QoeReportReq) GetScene() int64 {
	if x != nil {
		return x.Scene
	}
	return 0
}

func (x *QoeReportReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *QoeReportReq) GetCancel() bool {
	if x != nil {
		return x.Cancel
	}
	return false
}

func (x *QoeReportReq) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *QoeReportReq) GetOid() int64 {
	if x != nil {
		return x.Oid
	}
	return 0
}

func (x *QoeReportReq) GetScoreResult() *QoeScoreResult {
	if x != nil {
		return x.ScoreResult
	}
	return nil
}

func (x *QoeReportReq) GetBusinessData() string {
	if x != nil {
		return x.BusinessData
	}
	return ""
}

type QoeScoreResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *QoeScoreResult) Reset() {
	*x = QoeScoreResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_community_service_govern_v1_govern_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QoeScoreResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QoeScoreResult) ProtoMessage() {}

func (x *QoeScoreResult) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_community_service_govern_v1_govern_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QoeScoreResult.ProtoReflect.Descriptor instead.
func (*QoeScoreResult) Descriptor() ([]byte, []int) {
	return file_bilibili_community_service_govern_v1_govern_proto_rawDescGZIP(), []int{1}
}

func (x *QoeScoreResult) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_bilibili_community_service_govern_v1_govern_proto protoreflect.FileDescriptor

var file_bilibili_community_service_govern_v1_govern_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x76,
	0x65, 0x72, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0c, 0x51, 0x6f, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6f, 0x69, 0x64,
	0x12, 0x57, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x6f,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x26,
	0x0a, 0x0e, 0x51, 0x6f, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x5e, 0x0a, 0x03, 0x51, 0x6f, 0x65, 0x12, 0x57, 0x0a,
	0x09, 0x51, 0x6f, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x6f, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bilibili_community_service_govern_v1_govern_proto_rawDescOnce sync.Once
	file_bilibili_community_service_govern_v1_govern_proto_rawDescData = file_bilibili_community_service_govern_v1_govern_proto_rawDesc
)

func file_bilibili_community_service_govern_v1_govern_proto_rawDescGZIP() []byte {
	file_bilibili_community_service_govern_v1_govern_proto_rawDescOnce.Do(func() {
		file_bilibili_community_service_govern_v1_govern_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_community_service_govern_v1_govern_proto_rawDescData)
	})
	return file_bilibili_community_service_govern_v1_govern_proto_rawDescData
}

var file_bilibili_community_service_govern_v1_govern_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_community_service_govern_v1_govern_proto_goTypes = []interface{}{
	(*QoeReportReq)(nil),   // 0: bilibili.community.service.govern.v1.QoeReportReq
	(*QoeScoreResult)(nil), // 1: bilibili.community.service.govern.v1.QoeScoreResult
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_bilibili_community_service_govern_v1_govern_proto_depIdxs = []int32{
	1, // 0: bilibili.community.service.govern.v1.QoeReportReq.score_result:type_name -> bilibili.community.service.govern.v1.QoeScoreResult
	0, // 1: bilibili.community.service.govern.v1.Qoe.QoeReport:input_type -> bilibili.community.service.govern.v1.QoeReportReq
	2, // 2: bilibili.community.service.govern.v1.Qoe.QoeReport:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_community_service_govern_v1_govern_proto_init() }
func file_bilibili_community_service_govern_v1_govern_proto_init() {
	if File_bilibili_community_service_govern_v1_govern_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_community_service_govern_v1_govern_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QoeReportReq); i {
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
		file_bilibili_community_service_govern_v1_govern_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QoeScoreResult); i {
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
			RawDescriptor: file_bilibili_community_service_govern_v1_govern_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_community_service_govern_v1_govern_proto_goTypes,
		DependencyIndexes: file_bilibili_community_service_govern_v1_govern_proto_depIdxs,
		MessageInfos:      file_bilibili_community_service_govern_v1_govern_proto_msgTypes,
	}.Build()
	File_bilibili_community_service_govern_v1_govern_proto = out.File
	file_bilibili_community_service_govern_v1_govern_proto_rawDesc = nil
	file_bilibili_community_service_govern_v1_govern_proto_goTypes = nil
	file_bilibili_community_service_govern_v1_govern_proto_depIdxs = nil
}
