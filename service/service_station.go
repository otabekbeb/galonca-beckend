package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type StationService struct {
	api.UnimplementedServiceStationServiceServer
	storage *model.ServiceStationStorage
}

func NewServiceStationService() *StationService {
	return &StationService{
		storage: model.NewServiceStationStorage(),
	}
}

func (server *StationService) Create(ctx context.Context, req *api.ServiceStation) (*api.ServiceStationResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.ServiceStation: %v", req)
	serviceStation, err := server.storage.Create(req, userId)
	if err != nil || serviceStation == nil {
		res := &api.ServiceStationResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create serviceStation")
	}
	res := &api.ServiceStationResponse{
		Id:      uint64(serviceStation.ID),
		Success: true,
	}
	return res, nil
}

func (server *StationService) Update(ctx context.Context, req *api.ServiceStation) (*api.ServiceStationResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Station: %v", req)
	station, err := server.storage.Update(req, userId)
	if err != nil || station == nil {
		res := &api.ServiceStationResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create station")
	}
	res := &api.ServiceStationResponse{
		Id:      uint64(station.ID),
		Success: true,
	}
	return res, nil
}

func (server *StationService) Delete(_ context.Context, req *api.ServiceStationRequest) (*api.ServiceStationResponse, error) {
	serviceStationId := server.storage.Delete(req.GetId())
	if serviceStationId == 0 {
		return &api.ServiceStationResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete serviceStation")
	}
	return &api.ServiceStationResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *StationService) Get(_ context.Context, req *api.ServiceStationRequest) (*api.FindServiceStationResult, error) {
	serviceStation := server.storage.Get(req.GetId())
	if serviceStation == nil {
		return nil, status.Errorf(codes.NotFound, "ServiceStation not found")
	}

	return serviceStation.ToAPI(), nil
}

func (server *StationService) Find(_ context.Context, req *api.FindServiceStationRequest) (*api.FindServiceStationResponse, error) {
	serviceStations, count := server.storage.Find(
		req.GetCategories(),
		req.GetLocation(),
		req.GetLimit(),
		req.GetOffset(),
	)
	var result api.FindServiceStationResponse
	result.ServiceStations = make([]*api.FindServiceStationResult, len(serviceStations))
	for i, c := range serviceStations {
		ca := c.ToAPI()
		log.Printf("ServiceStationsResult: %v", ca)
		result.ServiceStations[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
