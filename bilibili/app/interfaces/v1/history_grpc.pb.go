// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/app/interfaces/v1/history.proto

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

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryClient interface {
	// 获取历史记录tab
	HistoryTab(ctx context.Context, in *HistoryTabReq, opts ...grpc.CallOption) (*HistoryTabReply, error)
	// 获取历史记录列表(旧版)
	Cursor(ctx context.Context, in *CursorReq, opts ...grpc.CallOption) (*CursorReply, error)
	// 获取历史记录列表
	CursorV2(ctx context.Context, in *CursorV2Req, opts ...grpc.CallOption) (*CursorV2Reply, error)
	// 删除历史记录
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*NoReply, error)
	// 搜索历史记录
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchReply, error)
	// 清空历史记录
	Clear(ctx context.Context, in *ClearReq, opts ...grpc.CallOption) (*NoReply, error)
	// 获取最新的历史记录
	LatestHistory(ctx context.Context, in *LatestHistoryReq, opts ...grpc.CallOption) (*LatestHistoryReply, error)
}

type historyClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryClient(cc grpc.ClientConnInterface) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) HistoryTab(ctx context.Context, in *HistoryTabReq, opts ...grpc.CallOption) (*HistoryTabReply, error) {
	out := new(HistoryTabReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/HistoryTab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) Cursor(ctx context.Context, in *CursorReq, opts ...grpc.CallOption) (*CursorReply, error) {
	out := new(CursorReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/Cursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) CursorV2(ctx context.Context, in *CursorV2Req, opts ...grpc.CallOption) (*CursorV2Reply, error) {
	out := new(CursorV2Reply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/CursorV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) Clear(ctx context.Context, in *ClearReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) LatestHistory(ctx context.Context, in *LatestHistoryReq, opts ...grpc.CallOption) (*LatestHistoryReply, error) {
	out := new(LatestHistoryReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.History/LatestHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
// All implementations must embed UnimplementedHistoryServer
// for forward compatibility
type HistoryServer interface {
	// 获取历史记录tab
	HistoryTab(context.Context, *HistoryTabReq) (*HistoryTabReply, error)
	// 获取历史记录列表(旧版)
	Cursor(context.Context, *CursorReq) (*CursorReply, error)
	// 获取历史记录列表
	CursorV2(context.Context, *CursorV2Req) (*CursorV2Reply, error)
	// 删除历史记录
	Delete(context.Context, *DeleteReq) (*NoReply, error)
	// 搜索历史记录
	Search(context.Context, *SearchReq) (*SearchReply, error)
	// 清空历史记录
	Clear(context.Context, *ClearReq) (*NoReply, error)
	// 获取最新的历史记录
	LatestHistory(context.Context, *LatestHistoryReq) (*LatestHistoryReply, error)
	mustEmbedUnimplementedHistoryServer()
}

// UnimplementedHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServer struct {
}

func (UnimplementedHistoryServer) HistoryTab(context.Context, *HistoryTabReq) (*HistoryTabReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryTab not implemented")
}
func (UnimplementedHistoryServer) Cursor(context.Context, *CursorReq) (*CursorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cursor not implemented")
}
func (UnimplementedHistoryServer) CursorV2(context.Context, *CursorV2Req) (*CursorV2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CursorV2 not implemented")
}
func (UnimplementedHistoryServer) Delete(context.Context, *DeleteReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHistoryServer) Search(context.Context, *SearchReq) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedHistoryServer) Clear(context.Context, *ClearReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedHistoryServer) LatestHistory(context.Context, *LatestHistoryReq) (*LatestHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestHistory not implemented")
}
func (UnimplementedHistoryServer) mustEmbedUnimplementedHistoryServer() {}

// UnsafeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServer will
// result in compilation errors.
type UnsafeHistoryServer interface {
	mustEmbedUnimplementedHistoryServer()
}

func RegisterHistoryServer(s grpc.ServiceRegistrar, srv HistoryServer) {
	s.RegisterService(&History_ServiceDesc, srv)
}

func _History_HistoryTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryTabReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).HistoryTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/HistoryTab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).HistoryTab(ctx, req.(*HistoryTabReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_Cursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CursorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Cursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/Cursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Cursor(ctx, req.(*CursorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_CursorV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CursorV2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).CursorV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/CursorV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).CursorV2(ctx, req.(*CursorV2Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Clear(ctx, req.(*ClearReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_LatestHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).LatestHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.History/LatestHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).LatestHistory(ctx, req.(*LatestHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// History_ServiceDesc is the grpc.ServiceDesc for History service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var History_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interface.v1.History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HistoryTab",
			Handler:    _History_HistoryTab_Handler,
		},
		{
			MethodName: "Cursor",
			Handler:    _History_Cursor_Handler,
		},
		{
			MethodName: "CursorV2",
			Handler:    _History_CursorV2_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _History_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _History_Search_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _History_Clear_Handler,
		},
		{
			MethodName: "LatestHistory",
			Handler:    _History_LatestHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/interfaces/v1/history.proto",
}
