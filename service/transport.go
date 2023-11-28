package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"tsl_server/api"
	"tsl_server/model"
)

type TransportService struct {
	api.UnimplementedTransportServiceServer
	storage *model.TransportStorage
}

func NewTransportService() *TransportService {
	return &TransportService{
		storage: model.NewTransportStorage(),
	}
}

func (server *TransportService) Create(ctx context.Context, req *api.Transport) (*api.TransportResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Transport: %v", req)
	transport, err := server.storage.Create(req, userId)
	if err != nil || transport == nil {
		res := &api.TransportResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create transport")
	}
	res := &api.TransportResponse{
		Id:      uint64(transport.ID),
		Success: true,
	}
	return res, nil
}

func (server *TransportService) Update(ctx context.Context, req *api.Transport) (*api.TransportResponse, error) {
	userId := ctx.Value("userId").(uint64)
	log.Printf("api.Transport: %v", req)
	transport, err := server.storage.Update(req, userId)
	if err != nil || transport == nil {
		res := &api.TransportResponse{
			Success: false,
		}
		return res, status.Errorf(codes.Internal, "failed to create transport")
	}
	res := &api.TransportResponse{
		Id:      uint64(transport.ID),
		Success: true,
	}
	return res, nil
}

func (server *TransportService) Delete(_ context.Context, req *api.TransportRequest) (*api.TransportResponse, error) {
	transportId := server.storage.Delete(req.GetId())
	if transportId == 0 {
		return &api.TransportResponse{
			Success: false,
		}, status.Errorf(codes.NotFound, "Cannot delete transport")
	}
	return &api.TransportResponse{
		Id:      req.GetId(),
		Success: true,
	}, nil
}

func (server *TransportService) Get(_ context.Context, req *api.TransportRequest) (*api.FindTransportResult, error) {
	transport := server.storage.Get(req.GetId())
	if transport == nil {
		return nil, status.Errorf(codes.NotFound, "Transport not found")
	}

	return transport.ToAPI(), nil
}

func (server *TransportService) Find(_ context.Context, req *api.FindTransportRequest) (*api.FindTransportResponse, error) {
	transports, count := server.storage.Find(
		req.GetType(),
		req.GetPlacement(),
		req.GetCondition(),
		req.GetTransportType(),
		req.GetFuelType(),
		req.GetBrand(),
		req.GetCostFrom(),
		req.GetCostTill(),
		req.GetLimit(),
		req.GetOffset(),
	)
	var result api.FindTransportResponse
	result.Transports = make([]*api.FindTransportResult, len(transports))
	for i, c := range transports {
		ca := c.ToAPI()
		//log.Printf("TransportsResult: %v", ca)
		result.Transports[i] = ca
	}
	result.Found = uint64(count)
	return &result, nil
}
