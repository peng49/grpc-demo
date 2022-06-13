// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/Hello.proto

package services

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	//简单模式
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	//服务端流模式
	SayHelloForServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayHelloForServerStreamClient, error)
	//客户端流模式
	SayHelloForClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloForClientStreamClient, error)
	//双向流模式
	SayHelloForStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloForStreamClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayHelloForServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayHelloForServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/HelloService/SayHelloForServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloForServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_SayHelloForServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloForServerStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloForServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloForClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloForClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], "/HelloService/SayHelloForClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloForClientStreamClient{stream}
	return x, nil
}

type HelloService_SayHelloForClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloForClientStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloForClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloForClientStreamClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloForStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloForStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], "/HelloService/SayHelloForStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloForStreamClient{stream}
	return x, nil
}

type HelloService_SayHelloForStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloForStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloForStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloForStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	//简单模式
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	//服务端流模式
	SayHelloForServerStream(*HelloRequest, HelloService_SayHelloForServerStreamServer) error
	//客户端流模式
	SayHelloForClientStream(HelloService_SayHelloForClientStreamServer) error
	//双向流模式
	SayHelloForStream(HelloService_SayHelloForStreamServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloForServerStream(*HelloRequest, HelloService_SayHelloForServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloForServerStream not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloForClientStream(HelloService_SayHelloForClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloForClientStream not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloForStream(HelloService_SayHelloForStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloForStream not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayHelloForServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).SayHelloForServerStream(m, &helloServiceSayHelloForServerStreamServer{stream})
}

type HelloService_SayHelloForServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceSayHelloForServerStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloForServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_SayHelloForClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloForClientStream(&helloServiceSayHelloForClientStreamServer{stream})
}

type HelloService_SayHelloForClientStreamServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloForClientStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloForClientStreamServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloForClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_SayHelloForStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloForStream(&helloServiceSayHelloForStreamServer{stream})
}

type HelloService_SayHelloForStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloForStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloForStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloForStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloForServerStream",
			Handler:       _HelloService_SayHelloForServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloForClientStream",
			Handler:       _HelloService_SayHelloForClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloForStream",
			Handler:       _HelloService_SayHelloForStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/Hello.proto",
}
