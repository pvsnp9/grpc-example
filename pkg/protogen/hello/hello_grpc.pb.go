// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/proto/hello/hello.proto

package hello

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
	HelloService_SayHello_FullMethodName           = "/hello.HelloService/SayHello"
	HelloService_SayManyHellos_FullMethodName      = "/hello.HelloService/sayManyHellos"
	HelloService_SayHelloToEveryOne_FullMethodName = "/hello.HelloService/SayHelloToEveryOne"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayManyHellos(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayManyHellosClient, error)
	SayHelloToEveryOne(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloToEveryOneClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayManyHellos(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayManyHellosClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_SayManyHellos_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayManyHellosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_SayManyHellosClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayManyHellosClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayManyHellosClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloToEveryOne(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloToEveryOneClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], HelloService_SayHelloToEveryOne_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloToEveryOneClient{stream}
	return x, nil
}

type HelloService_SayHelloToEveryOneClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloToEveryOneClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloToEveryOneClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloToEveryOneClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
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
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayManyHellos(*HelloRequest, HelloService_SayManyHellosServer) error
	SayHelloToEveryOne(HelloService_SayHelloToEveryOneServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) SayManyHellos(*HelloRequest, HelloService_SayManyHellosServer) error {
	return status.Errorf(codes.Unimplemented, "method SayManyHellos not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloToEveryOne(HelloService_SayHelloToEveryOneServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloToEveryOne not implemented")
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
		FullMethod: HelloService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayManyHellos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).SayManyHellos(m, &helloServiceSayManyHellosServer{stream})
}

type HelloService_SayManyHellosServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceSayManyHellosServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayManyHellosServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_SayHelloToEveryOne_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloToEveryOne(&helloServiceSayHelloToEveryOneServer{stream})
}

type HelloService_SayHelloToEveryOneServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloToEveryOneServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloToEveryOneServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloToEveryOneServer) Recv() (*HelloRequest, error) {
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
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sayManyHellos",
			Handler:       _HelloService_SayManyHellos_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloToEveryOne",
			Handler:       _HelloService_SayHelloToEveryOne_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/hello/hello.proto",
}
