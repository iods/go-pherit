// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package beta

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

// BetaServiceClient is the client API for BetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BetaServiceClient interface {
	SayHello(ctx context.Context, in *Beta, opts ...grpc.CallOption) (*Beta, error)
}

type betaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBetaServiceClient(cc grpc.ClientConnInterface) BetaServiceClient {
	return &betaServiceClient{cc}
}

func (c *betaServiceClient) SayHello(ctx context.Context, in *Beta, opts ...grpc.CallOption) (*Beta, error) {
	out := new(Beta)
	err := c.cc.Invoke(ctx, "/BetaService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BetaServiceServer is the server API for BetaService service.
// All implementations must embed UnimplementedBetaServiceServer
// for forward compatibility
type BetaServiceServer interface {
	SayHello(context.Context, *Beta) (*Beta, error)
	mustEmbedUnimplementedBetaServiceServer()
}

// UnimplementedBetaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBetaServiceServer struct {
}

func (UnimplementedBetaServiceServer) SayHello(context.Context, *Beta) (*Beta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBetaServiceServer) mustEmbedUnimplementedBetaServiceServer() {}

// UnsafeBetaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BetaServiceServer will
// result in compilation errors.
type UnsafeBetaServiceServer interface {
	mustEmbedUnimplementedBetaServiceServer()
}

func RegisterBetaServiceServer(s grpc.ServiceRegistrar, srv BetaServiceServer) {
	s.RegisterService(&BetaService_ServiceDesc, srv)
}

func _BetaService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Beta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetaServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BetaService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetaServiceServer).SayHello(ctx, req.(*Beta))
	}
	return interceptor(ctx, in, info, handler)
}

// BetaService_ServiceDesc is the grpc.ServiceDesc for BetaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BetaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BetaService",
	HandlerType: (*BetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BetaService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/serviceBeta.proto",
}
