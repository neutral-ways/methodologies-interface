// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: pb/methodology.proto

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

// MethodologyServiceClient is the client API for MethodologyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MethodologyServiceClient interface {
	Calculate(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	GetDetail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type methodologyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMethodologyServiceClient(cc grpc.ClientConnInterface) MethodologyServiceClient {
	return &methodologyServiceClient{cc}
}

func (c *methodologyServiceClient) Calculate(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/methodology.MethodologyService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *methodologyServiceClient) GetDetail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, "/methodology.MethodologyService/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MethodologyServiceServer is the server API for MethodologyService service.
// All implementations must embed UnimplementedMethodologyServiceServer
// for forward compatibility
type MethodologyServiceServer interface {
	Calculate(context.Context, *DataRequest) (*CalculateResponse, error)
	GetDetail(context.Context, *DetailRequest) (*DetailResponse, error)
	mustEmbedUnimplementedMethodologyServiceServer()
}

// UnimplementedMethodologyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMethodologyServiceServer struct {
}

func (UnimplementedMethodologyServiceServer) Calculate(context.Context, *DataRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedMethodologyServiceServer) GetDetail(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedMethodologyServiceServer) mustEmbedUnimplementedMethodologyServiceServer() {}

// UnsafeMethodologyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MethodologyServiceServer will
// result in compilation errors.
type UnsafeMethodologyServiceServer interface {
	mustEmbedUnimplementedMethodologyServiceServer()
}

func RegisterMethodologyServiceServer(s grpc.ServiceRegistrar, srv MethodologyServiceServer) {
	s.RegisterService(&MethodologyService_ServiceDesc, srv)
}

func _MethodologyService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodologyServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/methodology.MethodologyService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodologyServiceServer).Calculate(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MethodologyService_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodologyServiceServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/methodology.MethodologyService/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodologyServiceServer).GetDetail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MethodologyService_ServiceDesc is the grpc.ServiceDesc for MethodologyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MethodologyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "methodology.MethodologyService",
	HandlerType: (*MethodologyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _MethodologyService_Calculate_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _MethodologyService_GetDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/methodology.proto",
}
