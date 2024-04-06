// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: stream.proto

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

const (
	Greeter_GetStream_FullMethodName = "/Greeter/GetStream"
	Greeter_PutStream_FullMethodName = "/Greeter/PutStream"
	Greeter_AllStream_FullMethodName = "/Greeter/AllStream"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// 服务端流模式
	GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error)
	// 客户端流模式
	PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error)
	// 双向流模式
	AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], Greeter_GetStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GetStreamClient interface {
	Recv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterGetStreamClient struct {
	grpc.ClientStream
}

func (x *greeterGetStreamClient) Recv() (*StreamResData, error) {
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], Greeter_PutStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterPutStreamClient{stream}
	return x, nil
}

type Greeter_PutStreamClient interface {
	Send(*StreamReqData) error
	CloseAndRecv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterPutStreamClient struct {
	grpc.ClientStream
}

func (x *greeterPutStreamClient) Send(m *StreamReqData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterPutStreamClient) CloseAndRecv() (*StreamResData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], Greeter_AllStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterAllStreamClient{stream}
	return x, nil
}

type Greeter_AllStreamClient interface {
	Send(*StreamReqData) error
	Recv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterAllStreamClient struct {
	grpc.ClientStream
}

func (x *greeterAllStreamClient) Send(m *StreamReqData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterAllStreamClient) Recv() (*StreamResData, error) {
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// 服务端流模式
	GetStream(*StreamReqData, Greeter_GetStreamServer) error
	// 客户端流模式
	PutStream(Greeter_PutStreamServer) error
	// 双向流模式
	AllStream(Greeter_AllStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) GetStream(*StreamReqData, Greeter_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedGreeterServer) PutStream(Greeter_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}
func (UnimplementedGreeterServer) AllStream(Greeter_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReqData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetStream(m, &greeterGetStreamServer{stream})
}

type Greeter_GetStreamServer interface {
	Send(*StreamResData) error
	grpc.ServerStream
}

type greeterGetStreamServer struct {
	grpc.ServerStream
}

func (x *greeterGetStreamServer) Send(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).PutStream(&greeterPutStreamServer{stream})
}

type Greeter_PutStreamServer interface {
	SendAndClose(*StreamResData) error
	Recv() (*StreamReqData, error)
	grpc.ServerStream
}

type greeterPutStreamServer struct {
	grpc.ServerStream
}

func (x *greeterPutStreamServer) SendAndClose(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterPutStreamServer) Recv() (*StreamReqData, error) {
	m := new(StreamReqData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).AllStream(&greeterAllStreamServer{stream})
}

type Greeter_AllStreamServer interface {
	Send(*StreamResData) error
	Recv() (*StreamReqData, error)
	grpc.ServerStream
}

type greeterAllStreamServer struct {
	grpc.ServerStream
}

func (x *greeterAllStreamServer) Send(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterAllStreamServer) Recv() (*StreamReqData, error) {
	m := new(StreamReqData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Greeter_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutStream",
			Handler:       _Greeter_PutStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _Greeter_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
