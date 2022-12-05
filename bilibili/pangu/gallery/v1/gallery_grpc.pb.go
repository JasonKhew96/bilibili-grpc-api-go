// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/pangu/gallery/v1/gallery.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GalleryInterfaceClient is the client API for GalleryInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GalleryInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoReply, error)
	ListNFTByMid(ctx context.Context, in *ListNFTByMidReq, opts ...grpc.CallOption) (*ListNFTByMidReply, error)
	ListOrderByMid(ctx context.Context, in *ListOrderByMidReq, opts ...grpc.CallOption) (*ListOrderByMidReply, error)
	BasicInfo(ctx context.Context, in *BasicInfoReq, opts ...grpc.CallOption) (*BasicInfoReply, error)
	UserCheck(ctx context.Context, in *UserCheckReq, opts ...grpc.CallOption) (*UserCheckReply, error)
	AgreePolicy(ctx context.Context, in *AgreePolicyReq, opts ...grpc.CallOption) (*AgreePolicyReply, error)
	GetLastPolicy(ctx context.Context, in *GetLastPolicyReq, opts ...grpc.CallOption) (*GetLastPolicyReply, error)
}

type galleryInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewGalleryInterfaceClient(cc grpc.ClientConnInterface) GalleryInterfaceClient {
	return &galleryInterfaceClient{cc}
}

func (c *galleryInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) UserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) ListNFTByMid(ctx context.Context, in *ListNFTByMidReq, opts ...grpc.CallOption) (*ListNFTByMidReply, error) {
	out := new(ListNFTByMidReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/ListNFTByMid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) ListOrderByMid(ctx context.Context, in *ListOrderByMidReq, opts ...grpc.CallOption) (*ListOrderByMidReply, error) {
	out := new(ListOrderByMidReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/ListOrderByMid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) BasicInfo(ctx context.Context, in *BasicInfoReq, opts ...grpc.CallOption) (*BasicInfoReply, error) {
	out := new(BasicInfoReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/BasicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) UserCheck(ctx context.Context, in *UserCheckReq, opts ...grpc.CallOption) (*UserCheckReply, error) {
	out := new(UserCheckReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/UserCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) AgreePolicy(ctx context.Context, in *AgreePolicyReq, opts ...grpc.CallOption) (*AgreePolicyReply, error) {
	out := new(AgreePolicyReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/AgreePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryInterfaceClient) GetLastPolicy(ctx context.Context, in *GetLastPolicyReq, opts ...grpc.CallOption) (*GetLastPolicyReply, error) {
	out := new(GetLastPolicyReply)
	err := c.cc.Invoke(ctx, "/bilibili.pangu.gallery.v1.GalleryInterface/GetLastPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GalleryInterfaceServer is the server API for GalleryInterface service.
// All implementations must embed UnimplementedGalleryInterfaceServer
// for forward compatibility
type GalleryInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoReply, error)
	ListNFTByMid(context.Context, *ListNFTByMidReq) (*ListNFTByMidReply, error)
	ListOrderByMid(context.Context, *ListOrderByMidReq) (*ListOrderByMidReply, error)
	BasicInfo(context.Context, *BasicInfoReq) (*BasicInfoReply, error)
	UserCheck(context.Context, *UserCheckReq) (*UserCheckReply, error)
	AgreePolicy(context.Context, *AgreePolicyReq) (*AgreePolicyReply, error)
	GetLastPolicy(context.Context, *GetLastPolicyReq) (*GetLastPolicyReply, error)
	mustEmbedUnimplementedGalleryInterfaceServer()
}

// UnimplementedGalleryInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedGalleryInterfaceServer struct {
}

func (UnimplementedGalleryInterfaceServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGalleryInterfaceServer) UserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedGalleryInterfaceServer) ListNFTByMid(context.Context, *ListNFTByMidReq) (*ListNFTByMidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNFTByMid not implemented")
}
func (UnimplementedGalleryInterfaceServer) ListOrderByMid(context.Context, *ListOrderByMidReq) (*ListOrderByMidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrderByMid not implemented")
}
func (UnimplementedGalleryInterfaceServer) BasicInfo(context.Context, *BasicInfoReq) (*BasicInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicInfo not implemented")
}
func (UnimplementedGalleryInterfaceServer) UserCheck(context.Context, *UserCheckReq) (*UserCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCheck not implemented")
}
func (UnimplementedGalleryInterfaceServer) AgreePolicy(context.Context, *AgreePolicyReq) (*AgreePolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreePolicy not implemented")
}
func (UnimplementedGalleryInterfaceServer) GetLastPolicy(context.Context, *GetLastPolicyReq) (*GetLastPolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastPolicy not implemented")
}
func (UnimplementedGalleryInterfaceServer) mustEmbedUnimplementedGalleryInterfaceServer() {}

// UnsafeGalleryInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GalleryInterfaceServer will
// result in compilation errors.
type UnsafeGalleryInterfaceServer interface {
	mustEmbedUnimplementedGalleryInterfaceServer()
}

func RegisterGalleryInterfaceServer(s grpc.ServiceRegistrar, srv GalleryInterfaceServer) {
	s.RegisterService(&GalleryInterface_ServiceDesc, srv)
}

func _GalleryInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).UserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_ListNFTByMid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNFTByMidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).ListNFTByMid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/ListNFTByMid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).ListNFTByMid(ctx, req.(*ListNFTByMidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_ListOrderByMid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderByMidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).ListOrderByMid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/ListOrderByMid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).ListOrderByMid(ctx, req.(*ListOrderByMidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_BasicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).BasicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/BasicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).BasicInfo(ctx, req.(*BasicInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_UserCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).UserCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/UserCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).UserCheck(ctx, req.(*UserCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_AgreePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).AgreePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/AgreePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).AgreePolicy(ctx, req.(*AgreePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryInterface_GetLastPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryInterfaceServer).GetLastPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.pangu.gallery.v1.GalleryInterface/GetLastPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryInterfaceServer).GetLastPolicy(ctx, req.(*GetLastPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GalleryInterface_ServiceDesc is the grpc.ServiceDesc for GalleryInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GalleryInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.pangu.gallery.v1.GalleryInterface",
	HandlerType: (*GalleryInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GalleryInterface_Ping_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _GalleryInterface_UserInfo_Handler,
		},
		{
			MethodName: "ListNFTByMid",
			Handler:    _GalleryInterface_ListNFTByMid_Handler,
		},
		{
			MethodName: "ListOrderByMid",
			Handler:    _GalleryInterface_ListOrderByMid_Handler,
		},
		{
			MethodName: "BasicInfo",
			Handler:    _GalleryInterface_BasicInfo_Handler,
		},
		{
			MethodName: "UserCheck",
			Handler:    _GalleryInterface_UserCheck_Handler,
		},
		{
			MethodName: "AgreePolicy",
			Handler:    _GalleryInterface_AgreePolicy_Handler,
		},
		{
			MethodName: "GetLastPolicy",
			Handler:    _GalleryInterface_GetLastPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/pangu/gallery/v1/gallery.proto",
}
