package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type RoadsideService struct {
	api.UnimplementedRoadsideServiceServiceServer
	storage *model.RoadsideServiceStorage
}

func NewRoadsideServiceService() *RoadsideService {
	return &RoadsideService{
		storage: model.NewRoadsideServiceStorage(),
	}
}

func (server *RoadsideService) Create(ctx context.Context, req *api.RoadsideService) (*api.RoadsideServiceResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.RoadsideService: %v", req)
	roadsideService, err := server.storage.Create(req, userId)
	if err != nil || roadsideService == nil {
		res := &api.RoadsideServiceResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create roadsideService")
	}
	res := &api.RoadsideServiceResponse{
		Id:      uint64(roadsideService.ID),
		Success: true,
	}
	return res, nil
}

func (server *RoadsideService) Update(ctx context.Context, req *api.RoadsideService) (*api.RoadsideServiceResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Roadside: %v", req)
	roadside, err := server.storage.Update(req, userId)
	if err != nil || roadside == nil {
		res := &api.RoadsideServiceResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create roadside")
	}
	res := &api.RoadsideServiceResponse{
		Id:      uint64(roadside.ID),
		Success: true,
	}
	return res, nil
}

func (server *RoadsideService) Delete(_ context.Context, req *api.RoadsideServiceRequest) (*api.RoadsideServiceResponse, error) {
	roadsideServiceId := server.storage.Delete(req.GetId())
	if roadsideServiceId == 0 {
		return &api.RoadsideServiceResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete roadsideService")
	}
	return &api.RoadsideServiceResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *RoadsideService) Get(_ context.Context, req *api.RoadsideServiceRequest) (*api.FindRoadsideServiceResult, error) {
	roadsideService := server.storage.Get(req.GetId())
	if roadsideService == nil {
		return nil, status.Errorf(codes.NotFound, "RoadsideService not found")
	}

	return roadsideService.ToAPI(), nil
}

func (server *RoadsideService) Find(_ context.Context, req *api.FindRoadsideServiceRequest) (*api.FindRoadsideServiceResponse, error) {
	roadsideServices, count := server.storage.Find(
		int32(req.GetType()),
		req.GetLocation(),
		req.GetLimit(),
		req.GetOffset(),
	)
	var result api.FindRoadsideServiceResponse
	result.RoadsideServices = make([]*api.FindRoadsideServiceResult, len(roadsideServices))
	for i, c := range roadsideServices {
		ca := c.ToAPI()
		log.Printf("RoadsideServicesResult: %v", ca)
		result.RoadsideServices[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
