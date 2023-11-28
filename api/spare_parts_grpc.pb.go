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

// SparePartServiceClient is the client API for SparePartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SparePartServiceClient interface {
	Get(ctx context.Context, in *SparePartRequest, opts ...grpc.CallOption) (*FindSparePartResult, error)
	Find(ctx context.Context, in *FindSparePartRequest, opts ...grpc.CallOption) (*FindSparePartResponse, error)
	Create(ctx context.Context, in *SparePart, opts ...grpc.CallOption) (*SparePartResponse, error)
	Update(ctx context.Context, in *SparePart, opts ...grpc.CallOption) (*SparePartResponse, error)
	Delete(ctx context.Context, in *SparePartRequest, opts ...grpc.CallOption) (*SparePartResponse, error)
}

type sparePartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSparePartServiceClient(cc grpc.ClientConnInterface) SparePartServiceClient {
	return &sparePartServiceClient{cc}
}

func (c *sparePartServiceClient) Get(ctx context.Context, in *SparePartRequest, opts ...grpc.CallOption) (*FindSparePartResult, error) {
	out := new(FindSparePartResult)
	err := c.cc.Invoke(ctx, "/api.SparePartService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparePartServiceClient) Find(ctx context.Context, in *FindSparePartRequest, opts ...grpc.CallOption) (*FindSparePartResponse, error) {
	out := new(FindSparePartResponse)
	err := c.cc.Invoke(ctx, "/api.SparePartService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparePartServiceClient) Create(ctx context.Context, in *SparePart, opts ...grpc.CallOption) (*SparePartResponse, error) {
	out := new(SparePartResponse)
	err := c.cc.Invoke(ctx, "/api.SparePartService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparePartServiceClient) Update(ctx context.Context, in *SparePart, opts ...grpc.CallOption) (*SparePartResponse, error) {
	out := new(SparePartResponse)
	err := c.cc.Invoke(ctx, "/api.SparePartService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparePartServiceClient) Delete(ctx context.Context, in *SparePartRequest, opts ...grpc.CallOption) (*SparePartResponse, error) {
	out := new(SparePartResponse)
	err := c.cc.Invoke(ctx, "/api.SparePartService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SparePartServiceServer is the server API for SparePartService service.
// All implementations must embed UnimplementedSparePartServiceServer
// for forward compatibility
type SparePartServiceServer interface {
	Get(context.Context, *SparePartRequest) (*FindSparePartResult, error)
	Find(context.Context, *FindSparePartRequest) (*FindSparePartResponse, error)
	Create(context.Context, *SparePart) (*SparePartResponse, error)
	Update(context.Context, *SparePart) (*SparePartResponse, error)
	Delete(context.Context, *SparePartRequest) (*SparePartResponse, error)
	mustEmbedUnimplementedSparePartServiceServer()
}

// UnimplementedSparePartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSparePartServiceServer struct {
}

func (UnimplementedSparePartServiceServer) Get(context.Context, *SparePartRequest) (*FindSparePartResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSparePartServiceServer) Find(context.Context, *FindSparePartRequest) (*FindSparePartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedSparePartServiceServer) Create(context.Context, *SparePart) (*SparePartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSparePartServiceServer) Update(context.Context, *SparePart) (*SparePartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSparePartServiceServer) Delete(context.Context, *SparePartRequest) (*SparePartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSparePartServiceServer) mustEmbedUnimplementedSparePartServiceServer() {}

// UnsafeSparePartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SparePartServiceServer will
// result in compilation errors.
type UnsafeSparePartServiceServer interface {
	mustEmbedUnimplementedSparePartServiceServer()
}

func RegisterSparePartServiceServer(s grpc.ServiceRegistrar, srv SparePartServiceServer) {
	s.RegisterService(&SparePartService_ServiceDesc, srv)
}

func _SparePartService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparePartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparePartServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SparePartService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparePartServiceServer).Get(ctx, req.(*SparePartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparePartService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSparePartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparePartServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SparePartService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparePartServiceServer).Find(ctx, req.(*FindSparePartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparePartService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparePart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparePartServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SparePartService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparePartServiceServer).Create(ctx, req.(*SparePart))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparePartService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparePart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparePartServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SparePartService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparePartServiceServer).Update(ctx, req.(*SparePart))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparePartService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparePartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparePartServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SparePartService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparePartServiceServer).Delete(ctx, req.(*SparePartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SparePartService_ServiceDesc is the grpc.ServiceDesc for SparePartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SparePartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SparePartService",
	HandlerType: (*SparePartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SparePartService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _SparePartService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SparePartService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SparePartService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SparePartService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spare_parts.proto",
}