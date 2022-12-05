// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/app/topic/v1/topic.proto

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

// TopicClient is the client API for Topic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicClient interface {
	TopicDetailsAll(ctx context.Context, in *TopicDetailsAllReq, opts ...grpc.CallOption) (*TopicDetailsAllReply, error)
	TopicDetailsFold(ctx context.Context, in *TopicDetailsFoldReq, opts ...grpc.CallOption) (*TopicDetailsFoldReply, error)
	TopicSetDetails(ctx context.Context, in *TopicSetDetailsReq, opts ...grpc.CallOption) (*TopicSetDetailsReply, error)
}

type topicClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicClient(cc grpc.ClientConnInterface) TopicClient {
	return &topicClient{cc}
}

func (c *topicClient) TopicDetailsAll(ctx context.Context, in *TopicDetailsAllReq, opts ...grpc.CallOption) (*TopicDetailsAllReply, error) {
	out := new(TopicDetailsAllReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.topic.v1.Topic/TopicDetailsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) TopicDetailsFold(ctx context.Context, in *TopicDetailsFoldReq, opts ...grpc.CallOption) (*TopicDetailsFoldReply, error) {
	out := new(TopicDetailsFoldReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.topic.v1.Topic/TopicDetailsFold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) TopicSetDetails(ctx context.Context, in *TopicSetDetailsReq, opts ...grpc.CallOption) (*TopicSetDetailsReply, error) {
	out := new(TopicSetDetailsReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.topic.v1.Topic/TopicSetDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServer is the server API for Topic service.
// All implementations must embed UnimplementedTopicServer
// for forward compatibility
type TopicServer interface {
	TopicDetailsAll(context.Context, *TopicDetailsAllReq) (*TopicDetailsAllReply, error)
	TopicDetailsFold(context.Context, *TopicDetailsFoldReq) (*TopicDetailsFoldReply, error)
	TopicSetDetails(context.Context, *TopicSetDetailsReq) (*TopicSetDetailsReply, error)
	mustEmbedUnimplementedTopicServer()
}

// UnimplementedTopicServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServer struct {
}

func (UnimplementedTopicServer) TopicDetailsAll(context.Context, *TopicDetailsAllReq) (*TopicDetailsAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicDetailsAll not implemented")
}
func (UnimplementedTopicServer) TopicDetailsFold(context.Context, *TopicDetailsFoldReq) (*TopicDetailsFoldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicDetailsFold not implemented")
}
func (UnimplementedTopicServer) TopicSetDetails(context.Context, *TopicSetDetailsReq) (*TopicSetDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicSetDetails not implemented")
}
func (UnimplementedTopicServer) mustEmbedUnimplementedTopicServer() {}

// UnsafeTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServer will
// result in compilation errors.
type UnsafeTopicServer interface {
	mustEmbedUnimplementedTopicServer()
}

func RegisterTopicServer(s grpc.ServiceRegistrar, srv TopicServer) {
	s.RegisterService(&Topic_ServiceDesc, srv)
}

func _Topic_TopicDetailsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicDetailsAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).TopicDetailsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.topic.v1.Topic/TopicDetailsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).TopicDetailsAll(ctx, req.(*TopicDetailsAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_TopicDetailsFold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicDetailsFoldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).TopicDetailsFold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.topic.v1.Topic/TopicDetailsFold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).TopicDetailsFold(ctx, req.(*TopicDetailsFoldReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_TopicSetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicSetDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).TopicSetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.topic.v1.Topic/TopicSetDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).TopicSetDetails(ctx, req.(*TopicSetDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Topic_ServiceDesc is the grpc.ServiceDesc for Topic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Topic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.topic.v1.Topic",
	HandlerType: (*TopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TopicDetailsAll",
			Handler:    _Topic_TopicDetailsAll_Handler,
		},
		{
			MethodName: "TopicDetailsFold",
			Handler:    _Topic_TopicDetailsFold_Handler,
		},
		{
			MethodName: "TopicSetDetails",
			Handler:    _Topic_TopicSetDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/topic/v1/topic.proto",
}
