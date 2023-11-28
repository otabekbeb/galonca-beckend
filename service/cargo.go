package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type CargoService struct {
	api.UnimplementedCargoServiceServer
	storage *model.CargoStorage
}

func NewCargoService() *CargoService {
	return &CargoService{
		storage: model.NewCargoStorage(),
	}
}

func (server *CargoService) Create(ctx context.Context, req *api.Cargo) (*api.CargoResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Cargo: %v", req)
	cargo, err := server.storage.Create(req, userId)
	if err != nil || cargo == nil {
		res := &api.CargoResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create cargo")
	}
	res := &api.CargoResponse{
		Id:      uint64(cargo.ID),
		Success: true,
	}
	return res, nil
}

func (server *CargoService) Update(ctx context.Context, req *api.Cargo) (*api.CargoResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Cargo: %v", req)
	cargo, err := server.storage.Update(req, userId)
	if err != nil || cargo == nil {
		res := &api.CargoResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create cargo")
	}
	res := &api.CargoResponse{
		Id:      uint64(cargo.ID),
		Success: true,
	}
	return res, nil
}

func (server *CargoService) Delete(_ context.Context, req *api.CargoRequest) (*api.CargoResponse, error) {
	cargoId := server.storage.Delete(req.GetId())
	if cargoId == 0 {
		return &api.CargoResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete cargo")
	}
	return &api.CargoResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *CargoService) Get(_ context.Context, req *api.CargoRequest) (*api.FindCargoResult, error) {
	cargo := server.storage.Get(req.GetId())
	if cargo == nil {
		return nil, status.Errorf(codes.NotFound, "Cargo not found")
	}

	return cargo.ToAPI(), nil
}

func (server *CargoService) Find(_ context.Context, req *api.FindCargoRequest) (*api.FindCargoResponse, error) {
	cargos, count := server.storage.Find(
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
	var result api.FindCargoResponse
	result.Cargos = make([]*api.FindCargoResult, len(cargos))
	for i, c := range cargos {
		ca := c.ToAPI()
		log.Printf("CargosResult: %v", ca)
		result.Cargos[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
