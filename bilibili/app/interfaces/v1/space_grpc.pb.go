// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/app/interfaces/v1/space.proto

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

// SpaceClient is the client API for Space service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceClient interface {
	SearchTab(ctx context.Context, in *SearchTabReq, opts ...grpc.CallOption) (*SearchTabReply, error)
	SearchArchive(ctx context.Context, in *SearchArchiveReq, opts ...grpc.CallOption) (*SearchArchiveReply, error)
	SearchDynamic(ctx context.Context, in *SearchDynamicReq, opts ...grpc.CallOption) (*SearchDynamicReply, error)
}

type spaceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceClient(cc grpc.ClientConnInterface) SpaceClient {
	return &spaceClient{cc}
}

func (c *spaceClient) SearchTab(ctx context.Context, in *SearchTabReq, opts ...grpc.CallOption) (*SearchTabReply, error) {
	out := new(SearchTabReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.Space/SearchTab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceClient) SearchArchive(ctx context.Context, in *SearchArchiveReq, opts ...grpc.CallOption) (*SearchArchiveReply, error) {
	out := new(SearchArchiveReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.Space/SearchArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceClient) SearchDynamic(ctx context.Context, in *SearchDynamicReq, opts ...grpc.CallOption) (*SearchDynamicReply, error) {
	out := new(SearchDynamicReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.Space/SearchDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServer is the server API for Space service.
// All implementations must embed UnimplementedSpaceServer
// for forward compatibility
type SpaceServer interface {
	SearchTab(context.Context, *SearchTabReq) (*SearchTabReply, error)
	SearchArchive(context.Context, *SearchArchiveReq) (*SearchArchiveReply, error)
	SearchDynamic(context.Context, *SearchDynamicReq) (*SearchDynamicReply, error)
	mustEmbedUnimplementedSpaceServer()
}

// UnimplementedSpaceServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceServer struct {
}

func (UnimplementedSpaceServer) SearchTab(context.Context, *SearchTabReq) (*SearchTabReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTab not implemented")
}
func (UnimplementedSpaceServer) SearchArchive(context.Context, *SearchArchiveReq) (*SearchArchiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchive not implemented")
}
func (UnimplementedSpaceServer) SearchDynamic(context.Context, *SearchDynamicReq) (*SearchDynamicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDynamic not implemented")
}
func (UnimplementedSpaceServer) mustEmbedUnimplementedSpaceServer() {}

// UnsafeSpaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServer will
// result in compilation errors.
type UnsafeSpaceServer interface {
	mustEmbedUnimplementedSpaceServer()
}

func RegisterSpaceServer(s grpc.ServiceRegistrar, srv SpaceServer) {
	s.RegisterService(&Space_ServiceDesc, srv)
}

func _Space_SearchTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTabReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServer).SearchTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.Space/SearchTab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServer).SearchTab(ctx, req.(*SearchTabReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Space_SearchArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServer).SearchArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.Space/SearchArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServer).SearchArchive(ctx, req.(*SearchArchiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Space_SearchDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDynamicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServer).SearchDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.Space/SearchDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServer).SearchDynamic(ctx, req.(*SearchDynamicReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Space_ServiceDesc is the grpc.ServiceDesc for Space service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Space_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interface.v1.Space",
	HandlerType: (*SpaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchTab",
			Handler:    _Space_SearchTab_Handler,
		},
		{
			MethodName: "SearchArchive",
			Handler:    _Space_SearchArchive_Handler,
		},
		{
			MethodName: "SearchDynamic",
			Handler:    _Space_SearchDynamic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/interfaces/v1/space.proto",
}