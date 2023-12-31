// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	Get(ctx context.Context, in *TransportRequest, opts ...grpc.CallOption) (*FindTransportResult, error)
	Find(ctx context.Context, in *FindTransportRequest, opts ...grpc.CallOption) (*FindTransportResponse, error)
	Create(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*TransportResponse, error)
	Update(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*TransportResponse, error)
	Delete(ctx context.Context, in *TransportRequest, opts ...grpc.CallOption) (*TransportResponse, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) Get(ctx context.Context, in *TransportRequest, opts ...grpc.CallOption) (*FindTransportResult, error) {
	out := new(FindTransportResult)
	err := c.cc.Invoke(ctx, "/api.TransportService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) Find(ctx context.Context, in *FindTransportRequest, opts ...grpc.CallOption) (*FindTransportResponse, error) {
	out := new(FindTransportResponse)
	err := c.cc.Invoke(ctx, "/api.TransportService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) Create(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*TransportResponse, error) {
	out := new(TransportResponse)
	err := c.cc.Invoke(ctx, "/api.TransportService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) Update(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*TransportResponse, error) {
	out := new(TransportResponse)
	err := c.cc.Invoke(ctx, "/api.TransportService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) Delete(ctx context.Context, in *TransportRequest, opts ...grpc.CallOption) (*TransportResponse, error) {
	out := new(TransportResponse)
	err := c.cc.Invoke(ctx, "/api.TransportService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility
type TransportServiceServer interface {
	Get(context.Context, *TransportRequest) (*FindTransportResult, error)
	Find(context.Context, *FindTransportRequest) (*FindTransportResponse, error)
	Create(context.Context, *Transport) (*TransportResponse, error)
	Update(context.Context, *Transport) (*TransportResponse, error)
	Delete(context.Context, *TransportRequest) (*TransportResponse, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (UnimplementedTransportServiceServer) Get(context.Context, *TransportRequest) (*FindTransportResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTransportServiceServer) Find(context.Context, *FindTransportRequest) (*FindTransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedTransportServiceServer) Create(context.Context, *Transport) (*TransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransportServiceServer) Update(context.Context, *Transport) (*TransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransportServiceServer) Delete(context.Context, *TransportRequest) (*TransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransportService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Get(ctx, req.(*TransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransportService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Find(ctx, req.(*FindTransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransportService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Create(ctx, req.(*Transport))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransportService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Update(ctx, req.(*Transport))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransportService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Delete(ctx, req.(*TransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TransportService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _TransportService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TransportService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransportService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransportService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
