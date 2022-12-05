// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/broadcast/message/editor/notify.proto

package editor

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

// OperationNotifyClient is the client API for OperationNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationNotifyClient interface {
	OperationNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OperationNotify_OperationNotifyClient, error)
}

type operationNotifyClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationNotifyClient(cc grpc.ClientConnInterface) OperationNotifyClient {
	return &operationNotifyClient{cc}
}

func (c *operationNotifyClient) OperationNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OperationNotify_OperationNotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &OperationNotify_ServiceDesc.Streams[0], "/bilibili.broadcast.message.editor.OperationNotify/OperationNotify", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationNotifyOperationNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationNotify_OperationNotifyClient interface {
	Recv() (*Notify, error)
	grpc.ClientStream
}

type operationNotifyOperationNotifyClient struct {
	grpc.ClientStream
}

func (x *operationNotifyOperationNotifyClient) Recv() (*Notify, error) {
	m := new(Notify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperationNotifyServer is the server API for OperationNotify service.
// All implementations must embed UnimplementedOperationNotifyServer
// for forward compatibility
type OperationNotifyServer interface {
	OperationNotify(*emptypb.Empty, OperationNotify_OperationNotifyServer) error
	mustEmbedUnimplementedOperationNotifyServer()
}

// UnimplementedOperationNotifyServer must be embedded to have forward compatible implementations.
type UnimplementedOperationNotifyServer struct {
}

func (UnimplementedOperationNotifyServer) OperationNotify(*emptypb.Empty, OperationNotify_OperationNotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method OperationNotify not implemented")
}
func (UnimplementedOperationNotifyServer) mustEmbedUnimplementedOperationNotifyServer() {}

// UnsafeOperationNotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationNotifyServer will
// result in compilation errors.
type UnsafeOperationNotifyServer interface {
	mustEmbedUnimplementedOperationNotifyServer()
}

func RegisterOperationNotifyServer(s grpc.ServiceRegistrar, srv OperationNotifyServer) {
	s.RegisterService(&OperationNotify_ServiceDesc, srv)
}

func _OperationNotify_OperationNotify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationNotifyServer).OperationNotify(m, &operationNotifyOperationNotifyServer{stream})
}

type OperationNotify_OperationNotifyServer interface {
	Send(*Notify) error
	grpc.ServerStream
}

type operationNotifyOperationNotifyServer struct {
	grpc.ServerStream
}

func (x *operationNotifyOperationNotifyServer) Send(m *Notify) error {
	return x.ServerStream.SendMsg(m)
}

// OperationNotify_ServiceDesc is the grpc.ServiceDesc for OperationNotify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationNotify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.editor.OperationNotify",
	HandlerType: (*OperationNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OperationNotify",
			Handler:       _OperationNotify_OperationNotify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/message/editor/notify.proto",
}
