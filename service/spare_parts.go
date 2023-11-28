package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type SparePartService struct {
	api.UnimplementedSparePartServiceServer
	storage *model.SparePartStorage
}

func NewSparePartService() *SparePartService {
	return &SparePartService{
		storage: model.NewSparePartStorage(),
	}
}

func (server *SparePartService) Create(ctx context.Context, req *api.SparePart) (*api.SparePartResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.SparePart: %v", req)
	SparePart, err := server.storage.Create(req, userId)
	if err != nil || SparePart == nil {
		res := &api.SparePartResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create SparePart")
	}
	res := &api.SparePartResponse{
		Id:      uint64(SparePart.ID),
		Success: true,
	}
	return res, nil
}

func (server *SparePartService) Update(ctx context.Context, req *api.SparePart) (*api.SparePartResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.SparePart: %v", req)
	sparePart, err := server.storage.Update(req, userId)
	if err != nil || sparePart == nil {
		res := &api.SparePartResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create sparePart")
	}
	res := &api.SparePartResponse{
		Id:      uint64(sparePart.ID),
		Success: true,
	}
	return res, nil
}

func (server *SparePartService) Delete(_ context.Context, req *api.SparePartRequest) (*api.SparePartResponse, error) {
	SparePartId := server.storage.Delete(req.GetId())
	if SparePartId == 0 {
		return &api.SparePartResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete SparePart")
	}
	return &api.SparePartResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *SparePartService) Get(_ context.Context, req *api.SparePartRequest) (*api.FindSparePartResult, error) {
	SparePart := server.storage.Get(req.GetId())
	if SparePart == nil {
		return nil, status.Errorf(codes.NotFound, "SparePart not found")
	}

	return SparePart.ToAPI(), nil
}

func (server *SparePartService) Find(_ context.Context, req *api.FindSparePartRequest) (*api.FindSparePartResponse, error) {
	SpareParts, count := server.storage.Find(
		req.GetCategories(),
		req.GetTransportTypes(),
		req.GetLocation(),
		req.GetLimit(),
		req.GetOffset(),
	)
	var result api.FindSparePartResponse
	result.SpareParts = make([]*api.FindSparePartResult, len(SpareParts))
	for i, c := range SpareParts {
		ca := c.ToAPI()
		log.Printf("SparePartsResult: %v", ca)
		result.SpareParts[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
