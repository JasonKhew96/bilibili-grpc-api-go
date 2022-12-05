// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/polymer/app/search/v1/search.proto

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

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	SearchAll(ctx context.Context, in *SearchAllRequest, opts ...grpc.CallOption) (*SearchAllResponse, error)
	SearchByType(ctx context.Context, in *SearchByTypeRequest, opts ...grpc.CallOption) (*SearchByTypeResponse, error)
	SearchComic(ctx context.Context, in *SearchComicRequest, opts ...grpc.CallOption) (*SearchComicResponse, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) SearchAll(ctx context.Context, in *SearchAllRequest, opts ...grpc.CallOption) (*SearchAllResponse, error) {
	out := new(SearchAllResponse)
	err := c.cc.Invoke(ctx, "/bilibili.polymer.app.search.v1.Search/SearchAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) SearchByType(ctx context.Context, in *SearchByTypeRequest, opts ...grpc.CallOption) (*SearchByTypeResponse, error) {
	out := new(SearchByTypeResponse)
	err := c.cc.Invoke(ctx, "/bilibili.polymer.app.search.v1.Search/SearchByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) SearchComic(ctx context.Context, in *SearchComicRequest, opts ...grpc.CallOption) (*SearchComicResponse, error) {
	out := new(SearchComicResponse)
	err := c.cc.Invoke(ctx, "/bilibili.polymer.app.search.v1.Search/SearchComic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility
type SearchServer interface {
	SearchAll(context.Context, *SearchAllRequest) (*SearchAllResponse, error)
	SearchByType(context.Context, *SearchByTypeRequest) (*SearchByTypeResponse, error)
	SearchComic(context.Context, *SearchComicRequest) (*SearchComicResponse, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (UnimplementedSearchServer) SearchAll(context.Context, *SearchAllRequest) (*SearchAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAll not implemented")
}
func (UnimplementedSearchServer) SearchByType(context.Context, *SearchByTypeRequest) (*SearchByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByType not implemented")
}
func (UnimplementedSearchServer) SearchComic(context.Context, *SearchComicRequest) (*SearchComicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComic not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_SearchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.polymer.app.search.v1.Search/SearchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchAll(ctx, req.(*SearchAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_SearchByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.polymer.app.search.v1.Search/SearchByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchByType(ctx, req.(*SearchByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_SearchComic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchComicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SearchComic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.polymer.app.search.v1.Search/SearchComic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SearchComic(ctx, req.(*SearchComicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.polymer.app.search.v1.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchAll",
			Handler:    _Search_SearchAll_Handler,
		},
		{
			MethodName: "SearchByType",
			Handler:    _Search_SearchByType_Handler,
		},
		{
			MethodName: "SearchComic",
			Handler:    _Search_SearchComic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/polymer/app/search/v1/search.proto",
}
