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

// TruckServiceClient is the client API for TruckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TruckServiceClient interface {
	Get(ctx context.Context, in *TruckRequest, opts ...grpc.CallOption) (*FindTruckResult, error)
	Find(ctx context.Context, in *FindTruckRequest, opts ...grpc.CallOption) (*FindTruckResponse, error)
	Create(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*TruckResponse, error)
	Update(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*TruckResponse, error)
	Delete(ctx context.Context, in *TruckRequest, opts ...grpc.CallOption) (*TruckResponse, error)
}

type truckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTruckServiceClient(cc grpc.ClientConnInterface) TruckServiceClient {
	return &truckServiceClient{cc}
}

func (c *truckServiceClient) Get(ctx context.Context, in *TruckRequest, opts ...grpc.CallOption) (*FindTruckResult, error) {
	out := new(FindTruckResult)
	err := c.cc.Invoke(ctx, "/api.TruckService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckServiceClient) Find(ctx context.Context, in *FindTruckRequest, opts ...grpc.CallOption) (*FindTruckResponse, error) {
	out := new(FindTruckResponse)
	err := c.cc.Invoke(ctx, "/api.TruckService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckServiceClient) Create(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*TruckResponse, error) {
	out := new(TruckResponse)
	err := c.cc.Invoke(ctx, "/api.TruckService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckServiceClient) Update(ctx context.Context, in *Truck, opts ...grpc.CallOption) (*TruckResponse, error) {
	out := new(TruckResponse)
	err := c.cc.Invoke(ctx, "/api.TruckService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckServiceClient) Delete(ctx context.Context, in *TruckRequest, opts ...grpc.CallOption) (*TruckResponse, error) {
	out := new(TruckResponse)
	err := c.cc.Invoke(ctx, "/api.TruckService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TruckServiceServer is the server API for TruckService service.
// All implementations must embed UnimplementedTruckServiceServer
// for forward compatibility
type TruckServiceServer interface {
	Get(context.Context, *TruckRequest) (*FindTruckResult, error)
	Find(context.Context, *FindTruckRequest) (*FindTruckResponse, error)
	Create(context.Context, *Truck) (*TruckResponse, error)
	Update(context.Context, *Truck) (*TruckResponse, error)
	Delete(context.Context, *TruckRequest) (*TruckResponse, error)
	mustEmbedUnimplementedTruckServiceServer()
}

// UnimplementedTruckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTruckServiceServer struct {
}

func (UnimplementedTruckServiceServer) Get(context.Context, *TruckRequest) (*FindTruckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTruckServiceServer) Find(context.Context, *FindTruckRequest) (*FindTruckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedTruckServiceServer) Create(context.Context, *Truck) (*TruckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTruckServiceServer) Update(context.Context, *Truck) (*TruckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTruckServiceServer) Delete(context.Context, *TruckRequest) (*TruckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTruckServiceServer) mustEmbedUnimplementedTruckServiceServer() {}

// UnsafeTruckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TruckServiceServer will
// result in compilation errors.
type UnsafeTruckServiceServer interface {
	mustEmbedUnimplementedTruckServiceServer()
}

func RegisterTruckServiceServer(s grpc.ServiceRegistrar, srv TruckServiceServer) {
	s.RegisterService(&TruckService_ServiceDesc, srv)
}

func _TruckService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TruckServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TruckService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TruckServiceServer).Get(ctx, req.(*TruckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TruckService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTruckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TruckServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TruckService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TruckServiceServer).Find(ctx, req.(*FindTruckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TruckService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Truck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TruckServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TruckService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TruckServiceServer).Create(ctx, req.(*Truck))
	}
	return interceptor(ctx, in, info, handler)
}

func _TruckService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Truck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TruckServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TruckService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TruckServiceServer).Update(ctx, req.(*Truck))
	}
	return interceptor(ctx, in, info, handler)
}

func _TruckService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TruckServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TruckService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TruckServiceServer).Delete(ctx, req.(*TruckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TruckService_ServiceDesc is the grpc.ServiceDesc for TruckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TruckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TruckService",
	HandlerType: (*TruckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TruckService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _TruckService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TruckService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TruckService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TruckService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "truck.proto",
}