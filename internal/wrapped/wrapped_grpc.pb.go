// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wrapped

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WrappedServiceClient is the client API for WrappedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WrappedServiceClient interface {
	GetMessage(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	GetBytesMessage(ctx context.Context, in *wrappers.BytesValue, opts ...grpc.CallOption) (*wrappers.BytesValue, error)
}

type wrappedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWrappedServiceClient(cc grpc.ClientConnInterface) WrappedServiceClient {
	return &wrappedServiceClient{cc}
}

func (c *wrappedServiceClient) GetMessage(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/wrapped.WrappedService/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappedServiceClient) GetBytesMessage(ctx context.Context, in *wrappers.BytesValue, opts ...grpc.CallOption) (*wrappers.BytesValue, error) {
	out := new(wrappers.BytesValue)
	err := c.cc.Invoke(ctx, "/wrapped.WrappedService/GetBytesMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WrappedServiceServer is the server API for WrappedService service.
// All implementations must embed UnimplementedWrappedServiceServer
// for forward compatibility
type WrappedServiceServer interface {
	GetMessage(context.Context, *wrappers.StringValue) (*wrappers.StringValue, error)
	GetBytesMessage(context.Context, *wrappers.BytesValue) (*wrappers.BytesValue, error)
	mustEmbedUnimplementedWrappedServiceServer()
}

// UnimplementedWrappedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWrappedServiceServer struct {
}

func (UnimplementedWrappedServiceServer) GetMessage(context.Context, *wrappers.StringValue) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedWrappedServiceServer) GetBytesMessage(context.Context, *wrappers.BytesValue) (*wrappers.BytesValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBytesMessage not implemented")
}
func (UnimplementedWrappedServiceServer) mustEmbedUnimplementedWrappedServiceServer() {}

// UnsafeWrappedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WrappedServiceServer will
// result in compilation errors.
type UnsafeWrappedServiceServer interface {
	mustEmbedUnimplementedWrappedServiceServer()
}

func RegisterWrappedServiceServer(s grpc.ServiceRegistrar, srv WrappedServiceServer) {
	s.RegisterService(&WrappedService_ServiceDesc, srv)
}

func _WrappedService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappedServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wrapped.WrappedService/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappedServiceServer).GetMessage(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _WrappedService_GetBytesMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.BytesValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappedServiceServer).GetBytesMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wrapped.WrappedService/GetBytesMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappedServiceServer).GetBytesMessage(ctx, req.(*wrappers.BytesValue))
	}
	return interceptor(ctx, in, info, handler)
}

// WrappedService_ServiceDesc is the grpc.ServiceDesc for WrappedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WrappedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wrapped.WrappedService",
	HandlerType: (*WrappedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _WrappedService_GetMessage_Handler,
		},
		{
			MethodName: "GetBytesMessage",
			Handler:    _WrappedService_GetBytesMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wrapped.proto",
}
