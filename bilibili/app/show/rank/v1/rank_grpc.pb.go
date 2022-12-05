// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/app/show/rank/v1/rank.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankClient interface {
	// 全站排行榜
	RankAll(ctx context.Context, in *RankAllResultReq, opts ...grpc.CallOption) (*RankListReply, error)
	// 分区排行榜
	RankRegion(ctx context.Context, in *RankRegionResultReq, opts ...grpc.CallOption) (*RankListReply, error)
}

type rankClient struct {
	cc grpc.ClientConnInterface
}

func NewRankClient(cc grpc.ClientConnInterface) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) RankAll(ctx context.Context, in *RankAllResultReq, opts ...grpc.CallOption) (*RankListReply, error) {
	out := new(RankListReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.show.v1.Rank/RankAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankClient) RankRegion(ctx context.Context, in *RankRegionResultReq, opts ...grpc.CallOption) (*RankListReply, error) {
	out := new(RankListReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.show.v1.Rank/RankRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
// All implementations must embed UnimplementedRankServer
// for forward compatibility
type RankServer interface {
	// 全站排行榜
	RankAll(context.Context, *RankAllResultReq) (*RankListReply, error)
	// 分区排行榜
	RankRegion(context.Context, *RankRegionResultReq) (*RankListReply, error)
	mustEmbedUnimplementedRankServer()
}

// UnimplementedRankServer must be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (UnimplementedRankServer) RankAll(context.Context, *RankAllResultReq) (*RankListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankAll not implemented")
}
func (UnimplementedRankServer) RankRegion(context.Context, *RankRegionResultReq) (*RankListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankRegion not implemented")
}
func (UnimplementedRankServer) mustEmbedUnimplementedRankServer() {}

// UnsafeRankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServer will
// result in compilation errors.
type UnsafeRankServer interface {
	mustEmbedUnimplementedRankServer()
}

func RegisterRankServer(s grpc.ServiceRegistrar, srv RankServer) {
	s.RegisterService(&Rank_ServiceDesc, srv)
}

func _Rank_RankAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankAllResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).RankAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.show.v1.Rank/RankAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).RankAll(ctx, req.(*RankAllResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rank_RankRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankRegionResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).RankRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.show.v1.Rank/RankRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).RankRegion(ctx, req.(*RankRegionResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rank_ServiceDesc is the grpc.ServiceDesc for Rank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.show.v1.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RankAll",
			Handler:    _Rank_RankAll_Handler,
		},
		{
			MethodName: "RankRegion",
			Handler:    _Rank_RankRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/show/rank/v1/rank.proto",
}
