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

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	Get(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*FindCompanyResponse, error)
	GetByUser(ctx context.Context, in *GetCompanyByUserRequest, opts ...grpc.CallOption) (*FindCompanyResponse, error)
	Create(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	Delete(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*CommonCompanyResponse, error)
	CreateReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*CommonCompanyResponse, error)
	ListReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error)
	AddEmployee(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ListEmployees(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*EmployeesResponse, error)
	RemoveEmployee(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonUserResponse, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) Get(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*FindCompanyResponse, error) {
	out := new(FindCompanyResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetByUser(ctx context.Context, in *GetCompanyByUserRequest, opts ...grpc.CallOption) (*FindCompanyResponse, error) {
	out := new(FindCompanyResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/GetByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Create(ctx context.Context, in *Company, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Delete(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*CommonCompanyResponse, error) {
	out := new(CommonCompanyResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CreateReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*CommonCompanyResponse, error) {
	out := new(CommonCompanyResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error) {
	out := new(ReviewResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/ListReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) AddEmployee(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListEmployees(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*EmployeesResponse, error) {
	out := new(EmployeesResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/ListEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) RemoveEmployee(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonUserResponse, error) {
	out := new(CommonUserResponse)
	err := c.cc.Invoke(ctx, "/api.CompanyService/RemoveEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility
type CompanyServiceServer interface {
	Get(context.Context, *GetCompanyRequest) (*FindCompanyResponse, error)
	GetByUser(context.Context, *GetCompanyByUserRequest) (*FindCompanyResponse, error)
	Create(context.Context, *Company) (*CreateCompanyResponse, error)
	Delete(context.Context, *DeleteCompanyRequest) (*CommonCompanyResponse, error)
	CreateReview(context.Context, *Review) (*CommonCompanyResponse, error)
	ListReview(context.Context, *ReviewRequest) (*ReviewResponse, error)
	AddEmployee(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ListEmployees(context.Context, *GetCompanyRequest) (*EmployeesResponse, error)
	RemoveEmployee(context.Context, *DeleteUserRequest) (*CommonUserResponse, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (UnimplementedCompanyServiceServer) Get(context.Context, *GetCompanyRequest) (*FindCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCompanyServiceServer) GetByUser(context.Context, *GetCompanyByUserRequest) (*FindCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUser not implemented")
}
func (UnimplementedCompanyServiceServer) Create(context.Context, *Company) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCompanyServiceServer) Delete(context.Context, *DeleteCompanyRequest) (*CommonCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCompanyServiceServer) CreateReview(context.Context, *Review) (*CommonCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedCompanyServiceServer) ListReview(context.Context, *ReviewRequest) (*ReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReview not implemented")
}
func (UnimplementedCompanyServiceServer) AddEmployee(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedCompanyServiceServer) ListEmployees(context.Context, *GetCompanyRequest) (*EmployeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEmployees not implemented")
}
func (UnimplementedCompanyServiceServer) RemoveEmployee(context.Context, *DeleteUserRequest) (*CommonUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEmployee not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Get(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/GetByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetByUser(ctx, req.(*GetCompanyByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Create(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Delete(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateReview(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/ListReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListReview(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddEmployee(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/ListEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListEmployees(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_RemoveEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).RemoveEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CompanyService/RemoveEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).RemoveEmployee(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CompanyService_Get_Handler,
		},
		{
			MethodName: "GetByUser",
			Handler:    _CompanyService_GetByUser_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CompanyService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CompanyService_Delete_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _CompanyService_CreateReview_Handler,
		},
		{
			MethodName: "ListReview",
			Handler:    _CompanyService_ListReview_Handler,
		},
		{
			MethodName: "AddEmployee",
			Handler:    _CompanyService_AddEmployee_Handler,
		},
		{
			MethodName: "ListEmployees",
			Handler:    _CompanyService_ListEmployees_Handler,
		},
		{
			MethodName: "RemoveEmployee",
			Handler:    _CompanyService_RemoveEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile.proto",
}

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	ListFavorites(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*FavoritesResponse, error)
	AddFavorite(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*CommonUserResponse, error)
	DeleteFavorite(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*CommonUserResponse, error)
	ListMyCargo(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	ListMyTransport(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	ListMyRoadsideService(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	ListMyServiceStation(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	ListMySpareParts(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	ListMyTruck(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) ListFavorites(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*FavoritesResponse, error) {
	out := new(FavoritesResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddFavorite(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*CommonUserResponse, error) {
	out := new(CommonUserResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/AddFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) DeleteFavorite(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*CommonUserResponse, error) {
	out := new(CommonUserResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/DeleteFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMyCargo(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMyCargo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMyTransport(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMyTransport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMyRoadsideService(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMyRoadsideService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMyServiceStation(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMyServiceStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMySpareParts(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMySpareParts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListMyTruck(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/api.Profile/ListMyTruck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	ListFavorites(context.Context, *FavoritesRequest) (*FavoritesResponse, error)
	AddFavorite(context.Context, *Favorite) (*CommonUserResponse, error)
	DeleteFavorite(context.Context, *Favorite) (*CommonUserResponse, error)
	ListMyCargo(context.Context, *MyRequest) (*MyResponse, error)
	ListMyTransport(context.Context, *MyRequest) (*MyResponse, error)
	ListMyRoadsideService(context.Context, *MyRequest) (*MyResponse, error)
	ListMyServiceStation(context.Context, *MyRequest) (*MyResponse, error)
	ListMySpareParts(context.Context, *MyRequest) (*MyResponse, error)
	ListMyTruck(context.Context, *MyRequest) (*MyResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) ListFavorites(context.Context, *FavoritesRequest) (*FavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFavorites not implemented")
}
func (UnimplementedProfileServer) AddFavorite(context.Context, *Favorite) (*CommonUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorite not implemented")
}
func (UnimplementedProfileServer) DeleteFavorite(context.Context, *Favorite) (*CommonUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorite not implemented")
}
func (UnimplementedProfileServer) ListMyCargo(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyCargo not implemented")
}
func (UnimplementedProfileServer) ListMyTransport(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyTransport not implemented")
}
func (UnimplementedProfileServer) ListMyRoadsideService(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyRoadsideService not implemented")
}
func (UnimplementedProfileServer) ListMyServiceStation(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyServiceStation not implemented")
}
func (UnimplementedProfileServer) ListMySpareParts(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMySpareParts not implemented")
}
func (UnimplementedProfileServer) ListMyTruck(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyTruck not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_ListFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListFavorites(ctx, req.(*FavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/AddFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddFavorite(ctx, req.(*Favorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_DeleteFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).DeleteFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/DeleteFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).DeleteFavorite(ctx, req.(*Favorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMyCargo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMyCargo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMyCargo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMyCargo(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMyTransport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMyTransport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMyTransport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMyTransport(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMyRoadsideService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMyRoadsideService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMyRoadsideService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMyRoadsideService(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMyServiceStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMyServiceStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMyServiceStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMyServiceStation(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMySpareParts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMySpareParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMySpareParts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMySpareParts(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListMyTruck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListMyTruck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Profile/ListMyTruck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListMyTruck(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFavorites",
			Handler:    _Profile_ListFavorites_Handler,
		},
		{
			MethodName: "AddFavorite",
			Handler:    _Profile_AddFavorite_Handler,
		},
		{
			MethodName: "DeleteFavorite",
			Handler:    _Profile_DeleteFavorite_Handler,
		},
		{
			MethodName: "ListMyCargo",
			Handler:    _Profile_ListMyCargo_Handler,
		},
		{
			MethodName: "ListMyTransport",
			Handler:    _Profile_ListMyTransport_Handler,
		},
		{
			MethodName: "ListMyRoadsideService",
			Handler:    _Profile_ListMyRoadsideService_Handler,
		},
		{
			MethodName: "ListMyServiceStation",
			Handler:    _Profile_ListMyServiceStation_Handler,
		},
		{
			MethodName: "ListMySpareParts",
			Handler:    _Profile_ListMySpareParts_Handler,
		},
		{
			MethodName: "ListMyTruck",
			Handler:    _Profile_ListMyTruck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile.proto",
}
