package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type TruckService struct {
	api.UnimplementedTruckServiceServer
	storage *model.TruckStorage
}

func NewTruckService() *TruckService {
	return &TruckService{
		storage: model.NewTruckStorage(),
	}
}

func (server *TruckService) Create(ctx context.Context, req *api.Truck) (*api.TruckResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Truck: %v", req)
	truck, err := server.storage.Create(req, userId)
	if err != nil || truck == nil {
		res := &api.TruckResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create truck")
	}
	res := &api.TruckResponse{
		Id:      uint64(truck.ID),
		Success: true,
	}
	return res, nil
}

func (server *TruckService) Update(ctx context.Context, req *api.Truck) (*api.TruckResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Truck: %v", req)
	truck, err := server.storage.Update(req, userId)
	if err != nil || truck == nil {
		res := &api.TruckResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create truck")
	}
	res := &api.TruckResponse{
		Id:      uint64(truck.ID),
		Success: true,
	}
	return res, nil
}

func (server *TruckService) Delete(_ context.Context, req *api.TruckRequest) (*api.TruckResponse, error) {
	truckId := server.storage.Delete(req.GetId())
	if truckId == 0 {
		return &api.TruckResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete truck")
	}
	return &api.TruckResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *TruckService) Get(_ context.Context, req *api.TruckRequest) (*api.FindTruckResult, error) {
	truck := server.storage.Get(req.GetId())
	if truck == nil {
		return nil, status.Errorf(codes.NotFound, "Truck not found")
	}

	return truck.ToAPI(), nil
}

func (server *TruckService) Find(_ context.Context, req *api.FindTruckRequest) (*api.FindTruckResponse, error) {
	trucks, count := server.storage.Find(
		req.GetGeoFrom(),
		req.GetGeoTo(),
		req.GetWeightFrom(),
		req.GetWeightTo(),
		req.GetVolumeFrom(),
		req.GetVolumeTo(),
		req.GetFrom().AsTime(),
		req.GetTill().AsTime(),
		req.GetLimit(),
		req.GetOffset(),
		req.GetTypeLoading(),
		req.GetTimeFilter().AsTime(),
		req.GetCargoType(),
	)
	var result api.FindTruckResponse
	result.Trucks = make([]*api.FindTruckResult, len(trucks))
	for i, c := range trucks {
		ca := c.ToAPI()
		log.Printf("Truck: %v,\n TrucksResult: %v", c, ca)
		result.Trucks[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
