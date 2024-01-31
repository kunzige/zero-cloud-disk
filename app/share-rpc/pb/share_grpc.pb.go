// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: share.proto

package pb

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

const (
	Sharecenter_FileShare_FullMethodName       = "/pb.sharecenter/FileShare"
	Sharecenter_CancelFileShare_FullMethodName = "/pb.sharecenter/CancelFileShare"
	Sharecenter_GetFileShare_FullMethodName    = "/pb.sharecenter/GetFileShare"
)

// SharecenterClient is the client API for Sharecenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharecenterClient interface {
	FileShare(ctx context.Context, in *FileShareReq, opts ...grpc.CallOption) (*FileShareRes, error)
	CancelFileShare(ctx context.Context, in *CancelFileShareReq, opts ...grpc.CallOption) (*CancelFileShareRes, error)
	GetFileShare(ctx context.Context, in *GetShareReq, opts ...grpc.CallOption) (*SharedFiles, error)
}

type sharecenterClient struct {
	cc grpc.ClientConnInterface
}

func NewSharecenterClient(cc grpc.ClientConnInterface) SharecenterClient {
	return &sharecenterClient{cc}
}

func (c *sharecenterClient) FileShare(ctx context.Context, in *FileShareReq, opts ...grpc.CallOption) (*FileShareRes, error) {
	out := new(FileShareRes)
	err := c.cc.Invoke(ctx, Sharecenter_FileShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharecenterClient) CancelFileShare(ctx context.Context, in *CancelFileShareReq, opts ...grpc.CallOption) (*CancelFileShareRes, error) {
	out := new(CancelFileShareRes)
	err := c.cc.Invoke(ctx, Sharecenter_CancelFileShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharecenterClient) GetFileShare(ctx context.Context, in *GetShareReq, opts ...grpc.CallOption) (*SharedFiles, error) {
	out := new(SharedFiles)
	err := c.cc.Invoke(ctx, Sharecenter_GetFileShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharecenterServer is the server API for Sharecenter service.
// All implementations must embed UnimplementedSharecenterServer
// for forward compatibility
type SharecenterServer interface {
	FileShare(context.Context, *FileShareReq) (*FileShareRes, error)
	CancelFileShare(context.Context, *CancelFileShareReq) (*CancelFileShareRes, error)
	GetFileShare(context.Context, *GetShareReq) (*SharedFiles, error)
	mustEmbedUnimplementedSharecenterServer()
}

// UnimplementedSharecenterServer must be embedded to have forward compatible implementations.
type UnimplementedSharecenterServer struct {
}

func (UnimplementedSharecenterServer) FileShare(context.Context, *FileShareReq) (*FileShareRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileShare not implemented")
}
func (UnimplementedSharecenterServer) CancelFileShare(context.Context, *CancelFileShareReq) (*CancelFileShareRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelFileShare not implemented")
}
func (UnimplementedSharecenterServer) GetFileShare(context.Context, *GetShareReq) (*SharedFiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileShare not implemented")
}
func (UnimplementedSharecenterServer) mustEmbedUnimplementedSharecenterServer() {}

// UnsafeSharecenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharecenterServer will
// result in compilation errors.
type UnsafeSharecenterServer interface {
	mustEmbedUnimplementedSharecenterServer()
}

func RegisterSharecenterServer(s grpc.ServiceRegistrar, srv SharecenterServer) {
	s.RegisterService(&Sharecenter_ServiceDesc, srv)
}

func _Sharecenter_FileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharecenterServer).FileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sharecenter_FileShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharecenterServer).FileShare(ctx, req.(*FileShareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sharecenter_CancelFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelFileShareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharecenterServer).CancelFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sharecenter_CancelFileShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharecenterServer).CancelFileShare(ctx, req.(*CancelFileShareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sharecenter_GetFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharecenterServer).GetFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sharecenter_GetFileShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharecenterServer).GetFileShare(ctx, req.(*GetShareReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sharecenter_ServiceDesc is the grpc.ServiceDesc for Sharecenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sharecenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.sharecenter",
	HandlerType: (*SharecenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileShare",
			Handler:    _Sharecenter_FileShare_Handler,
		},
		{
			MethodName: "CancelFileShare",
			Handler:    _Sharecenter_CancelFileShare_Handler,
		},
		{
			MethodName: "GetFileShare",
			Handler:    _Sharecenter_GetFileShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "share.proto",
}
