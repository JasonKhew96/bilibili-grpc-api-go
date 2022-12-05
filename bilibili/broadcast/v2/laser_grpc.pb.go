// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/broadcast/v2/laser.proto

package v2

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

// LaserClient is the client API for Laser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LaserClient interface {
	// 监听Laser事件
	WatchEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Laser_WatchEventClient, error)
}

type laserClient struct {
	cc grpc.ClientConnInterface
}

func NewLaserClient(cc grpc.ClientConnInterface) LaserClient {
	return &laserClient{cc}
}

func (c *laserClient) WatchEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Laser_WatchEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Laser_ServiceDesc.Streams[0], "/bilibili.broadcast.v2.Laser/WatchEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &laserWatchEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Laser_WatchEventClient interface {
	Recv() (*LaserEventResp, error)
	grpc.ClientStream
}

type laserWatchEventClient struct {
	grpc.ClientStream
}

func (x *laserWatchEventClient) Recv() (*LaserEventResp, error) {
	m := new(LaserEventResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaserServer is the server API for Laser service.
// All implementations must embed UnimplementedLaserServer
// for forward compatibility
type LaserServer interface {
	// 监听Laser事件
	WatchEvent(*emptypb.Empty, Laser_WatchEventServer) error
	mustEmbedUnimplementedLaserServer()
}

// UnimplementedLaserServer must be embedded to have forward compatible implementations.
type UnimplementedLaserServer struct {
}

func (UnimplementedLaserServer) WatchEvent(*emptypb.Empty, Laser_WatchEventServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchEvent not implemented")
}
func (UnimplementedLaserServer) mustEmbedUnimplementedLaserServer() {}

// UnsafeLaserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LaserServer will
// result in compilation errors.
type UnsafeLaserServer interface {
	mustEmbedUnimplementedLaserServer()
}

func RegisterLaserServer(s grpc.ServiceRegistrar, srv LaserServer) {
	s.RegisterService(&Laser_ServiceDesc, srv)
}

func _Laser_WatchEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaserServer).WatchEvent(m, &laserWatchEventServer{stream})
}

type Laser_WatchEventServer interface {
	Send(*LaserEventResp) error
	grpc.ServerStream
}

type laserWatchEventServer struct {
	grpc.ServerStream
}

func (x *laserWatchEventServer) Send(m *LaserEventResp) error {
	return x.ServerStream.SendMsg(m)
}

// Laser_ServiceDesc is the grpc.ServiceDesc for Laser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Laser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.v2.Laser",
	HandlerType: (*LaserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchEvent",
			Handler:       _Laser_WatchEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/v2/laser.proto",
}
