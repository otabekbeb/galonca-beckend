package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"tsl_server/api"
	"tsl_server/model"
)

type ProfileService struct {
	api.UnimplementedProfileServer
	storage *model.FavoriteStorage
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		storage: model.NewFavoriteStorage(),
	}
}

func (server *ProfileService) AddFavorite(ctx context.Context, req *api.Favorite) (*api.CommonUserResponse, error) {
	userId := ctx.Value("userId").(uint64)
	favorite := server.storage.AddFavorite(req, userId)
	var result api.CommonUserResponse
	if favorite == nil {
		result.Success = false
	} else {
		result.Success = true
	}
	return &result, nil
}

func (server *ProfileService) DeleteFavorite(ctx context.Context, req *api.Favorite) (*api.CommonUserResponse, error) {
	userId := ctx.Value("userId").(uint64)
	favorite := server.storage.DeleteFavorite(req, userId)
	var result api.CommonUserResponse
	if favorite == nil {
		result.Success = false
	} else {
		result.Success = true
	}
	return &result, nil
}

func (server *ProfileService) ListFavorites(ctx context.Context, req *api.FavoritesRequest) (*api.FavoritesResponse, error) {
	userId := ctx.Value("userId").(uint64)
	return server.storage.ListFavorites(req, userId), nil
}

func (server *ProfileService) ListMyCargo(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.Cargo{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_Cargo
	result.Cargo = new(api.FindCargoResponse)
	for _, en := range ent.([]model.Cargo) {
		result.Cargo.Cargos = append(result.Cargo.Cargos, en.ToAPI())
	}
	result.Cargo.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

func (server *ProfileService) ListMyTransport(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.Transport{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_Transport
	result.Transport = new(api.FindTransportResponse)
	for _, en := range ent.([]model.Transport) {
		result.Transport.Transports = append(result.Transport.Transports, en.ToAPI())
	}
	result.Transport.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

func (server *ProfileService) ListMyRoadsideService(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.RoadsideService{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_RoadsideService
	result.RoadsideService = new(api.FindRoadsideServiceResponse)
	for _, en := range ent.([]model.RoadsideService) {
		result.RoadsideService.RoadsideServices = append(result.RoadsideService.RoadsideServices, en.ToAPI())
	}
	result.RoadsideService.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

func (server *ProfileService) ListMyServiceStation(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.ServiceStation{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_ServiceStation
	result.ServiceStation = new(api.FindServiceStationResponse)
	for _, en := range ent.([]model.ServiceStation) {
		result.ServiceStation.ServiceStations = append(result.ServiceStation.ServiceStations, en.ToAPI())
	}
	result.ServiceStation.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

func (server *ProfileService) ListMySpareParts(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.SparePart{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_SpareParts
	result.SpareParts = new(api.FindSparePartResponse)
	for _, en := range ent.([]model.SparePart) {
		result.SpareParts.SpareParts = append(result.SpareParts.SpareParts, en.ToAPI())
	}
	result.SpareParts.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

func (server *ProfileService) ListMyTruck(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	ent, qty := server.storage.ListMy(model.Truck{}, req.GetLimit(), req.GetOffset(), uint(userId))
	var result api.MyResponse_Truck
	result.Truck = new(api.FindTruckResponse)
	for _, en := range ent.([]model.Truck) {
		result.Truck.Trucks = append(result.Truck.Trucks, en.ToAPI())
	}
	result.Truck.Found = uint64(qty)
	return &api.MyResponse{My: &result}, nil
}

type CompanyServiceService struct {
	api.UnimplementedCompanyServiceServer
	storage *model.CompanyStorage
}

func NewCompanyService() *CompanyServiceService {
	return &CompanyServiceService{
		storage: model.NewCompanyStorage(),
	}
}

func (server *CompanyServiceService) Create(ctx context.Context, req *api.Company) (*api.CreateCompanyResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.CompanyServiceService: %v", req)
	CompanyServiceService, err := server.storage.Create(req, userId)
	if err != nil || CompanyServiceService == nil {
		res := &api.CreateCompanyResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create Company")
	}
	res := &api.CreateCompanyResponse{
		Id:      uint64(CompanyServiceService.ID),
		Success: true,
	}
	return res, nil
}

func (server *CompanyServiceService) Delete(_ context.Context, req *api.DeleteCompanyRequest) (*api.CommonCompanyResponse, error) {
	CompanyServiceServiceId := server.storage.Delete(req.GetId())
	if CompanyServiceServiceId == 0 {
		return &api.CommonCompanyResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete CompanyServiceService")
	}
	return &api.CommonCompanyResponse{
		Success: true,
	}, nil
}

func (server *CompanyServiceService) Get(_ context.Context, req *api.GetCompanyRequest) (*api.FindCompanyResponse, error) {
	company := server.storage.GetCompany(req.GetId())
	if company == nil {
		return nil, status.Errorf(codes.NotFound, "CompanyServiceService not found")
	}
	result := api.FindCompanyResponse{
		Id:           int64(company.ID),
		Name:         company.Name,
		Bin:          company.BIN,
		Email:        company.Email,
		Phone:        company.Phone,
		ActivityType: company.ActivityType,
		CityId:       company.CityId,
		Address:      company.Address,
		CreatedAt:    company.CreatedAt.Format("02.01.2006"),
		Documents:    strings.Split(company.Documents, ";"),
		Images:       strings.Split(company.Images, ";"),
	}
	result.Owner.Id = uint64(company.Owner.ID)
	result.Owner.Name = company.Owner.Name
	return &result, nil
}

func (server *CompanyServiceService) GetByUser(_ context.Context, req *api.GetCompanyByUserRequest) (*api.FindCompanyResponse, error) {
	company := server.storage.GetCompanyByUser(req.GetEmail())
	if company == nil {
		return nil, status.Errorf(codes.NotFound, "CompanyServiceService not found")
	}
	result := api.FindCompanyResponse{
		Id:           int64(company.ID),
		Name:         company.Name,
		Bin:          company.BIN,
		Email:        company.Email,
		Phone:        company.Phone,
		ActivityType: company.ActivityType,
		CityId:       company.CityId,
		Address:      company.Address,
		CreatedAt:    company.CreatedAt.Format("02.01.2006"),
		Documents:    strings.Split(company.Documents, ";"),
		Images:       strings.Split(company.Images, ";"),
	}
	log.Printf("company %v", company)
	result.Owner = new(api.FindCompanyResponse_Owner)
	result.Owner.Id = company.OwnerID
	result.Owner.Name = company.Owner.Name
	return &result, nil
}

func (server *CompanyServiceService) AddEmployee(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	userId := ctx.Value("userId").(uint64)
	user := server.storage.AddEmployee(req, userId)
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "Cannot create employee")
	}
	result := api.CreateUserResponse{
		Id: uint64(user.ID),
	}
	return &result, nil
}

func (server *CompanyServiceService) ListEmployees(_ context.Context, req *api.GetCompanyRequest) (*api.EmployeesResponse, error) {
	return server.storage.ListEmployees(req), nil
}

func (server *CompanyServiceService) RemoveEmployee(ctx context.Context, req *api.DeleteUserRequest) (*api.CommonUserResponse, error) {
	return server.storage.RemoveEmployees(req), nil
}
