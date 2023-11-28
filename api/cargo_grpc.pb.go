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

// CargoServiceClient is the client API for CargoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CargoServiceClient interface {
	Get(ctx context.Context, in *CargoRequest, opts ...grpc.CallOption) (*FindCargoResult, error)
	Find(ctx context.Context, in *FindCargoRequest, opts ...grpc.CallOption) (*FindCargoResponse, error)
	Create(ctx context.Context, in *Cargo, opts ...grpc.CallOption) (*CargoResponse, error)
	Update(ctx context.Context, in *Cargo, opts ...grpc.CallOption) (*CargoResponse, error)
	Delete(ctx context.Context, in *CargoRequest, opts ...grpc.CallOption) (*CargoResponse, error)
}

type cargoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCargoServiceClient(cc grpc.ClientConnInterface) CargoServiceClient {
	return &cargoServiceClient{cc}
}

func (c *cargoServiceClient) Get(ctx context.Context, in *CargoRequest, opts ...grpc.CallOption) (*FindCargoResult, error) {
	out := new(FindCargoResult)
	err := c.cc.Invoke(ctx, "/api.CargoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) Find(ctx context.Context, in *FindCargoRequest, opts ...grpc.CallOption) (*FindCargoResponse, error) {
	out := new(FindCargoResponse)
	err := c.cc.Invoke(ctx, "/api.CargoService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) Create(ctx context.Context, in *Cargo, opts ...grpc.CallOption) (*CargoResponse, error) {
	out := new(CargoResponse)
	err := c.cc.Invoke(ctx, "/api.CargoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) Update(ctx context.Context, in *Cargo, opts ...grpc.CallOption) (*CargoResponse, error) {
	out := new(CargoResponse)
	err := c.cc.Invoke(ctx, "/api.CargoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cargoServiceClient) Delete(ctx context.Context, in *CargoRequest, opts ...grpc.CallOption) (*CargoResponse, error) {
	out := new(CargoResponse)
	err := c.cc.Invoke(ctx, "/api.CargoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CargoServiceServer is the server API for CargoService service.
// All implementations must embed UnimplementedCargoServiceServer
// for forward compatibility
type CargoServiceServer interface {
	Get(context.Context, *CargoRequest) (*FindCargoResult, error)
	Find(context.Context, *FindCargoRequest) (*FindCargoResponse, error)
	Create(context.Context, *Cargo) (*CargoResponse, error)
	Update(context.Context, *Cargo) (*CargoResponse, error)
	Delete(context.Context, *CargoRequest) (*CargoResponse, error)
	mustEmbedUnimplementedCargoServiceServer()
}

// UnimplementedCargoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCargoServiceServer struct {
}

func (UnimplementedCargoServiceServer) Get(context.Context, *CargoRequest) (*FindCargoResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCargoServiceServer) Find(context.Context, *FindCargoRequest) (*FindCargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedCargoServiceServer) Create(context.Context, *Cargo) (*CargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCargoServiceServer) Update(context.Context, *Cargo) (*CargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCargoServiceServer) Delete(context.Context, *CargoRequest) (*CargoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCargoServiceServer) mustEmbedUnimplementedCargoServiceServer() {}

// UnsafeCargoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CargoServiceServer will
// result in compilation errors.
type UnsafeCargoServiceServer interface {
	mustEmbedUnimplementedCargoServiceServer()
}

func RegisterCargoServiceServer(s grpc.ServiceRegistrar, srv CargoServiceServer) {
	s.RegisterService(&CargoService_ServiceDesc, srv)
}

func _CargoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CargoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).Get(ctx, req.(*CargoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCargoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CargoService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).Find(ctx, req.(*FindCargoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cargo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CargoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).Create(ctx, req.(*Cargo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cargo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CargoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).Update(ctx, req.(*Cargo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CargoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CargoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoServiceServer).Delete(ctx, req.(*CargoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CargoService_ServiceDesc is the grpc.ServiceDesc for CargoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CargoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CargoService",
	HandlerType: (*CargoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CargoService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _CargoService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CargoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CargoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CargoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cargo.proto",
}