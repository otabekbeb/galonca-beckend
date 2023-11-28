package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tsl_server/api"
	"tsl_server/model"
)

type GeoService struct {
	api.UnimplementedGeoServiceServer
	storage *model.GeoStorage
}

func NewGeoService() *GeoService {
	return &GeoService{
		storage: model.NewGeoStorage(),
	}
}

func (server *GeoService) ListCountry(_ context.Context, req *api.CountryListRequest) (*api.CountryListResponse, error) {
	countries := server.storage.ListCountry(req.GetName())
	var result []*api.Country
	for _, v := range countries {
		result = append(result, v.ToAPI())
	}
	return &api.CountryListResponse{Countries: result}, nil
}

func (server *GeoService) ListRegion(_ context.Context, req *api.RegionListRequest) (*api.RegionListResponse, error) {
	regions := server.storage.ListRegion(req.GetName())
	var result []*api.Region
	for _, v := range regions {
		result = append(result, v.ToAPI())
	}
	return &api.RegionListResponse{Regions: result}, nil
}

func (server *GeoService) ListCity(_ context.Context, req *api.CityListRequest) (*api.CityListResponse, error) {
	cities := server.storage.ListCity(req.GetName())
	var result []*api.City
	for _, v := range cities {
		result = append(result, v.ToAPI())
	}
	return &api.CityListResponse{Cities: result}, nil
}

func (server *GeoService) GetCity(_ context.Context, req *api.GetCityRequest) (*api.City, error) {
	city := server.storage.GetCity(uint(req.GetId()))
	if city == nil {
		return nil, status.Errorf(codes.NotFound, "City not found")
	}
	return city.ToAPI(), nil
}

func (server *GeoService) GetCountry(_ context.Context, req *api.GetCountryRequest) (*api.Country, error) {
	country := server.storage.GetCountry(uint(req.GetId()))
	if country == nil {
		return nil, status.Errorf(codes.NotFound, "Country not found")
	}
	return country.ToAPI(), nil
}

func (server *GeoService) GetRegion(_ context.Context, req *api.GetRegionRequest) (*api.Region, error) {
	region := server.storage.GetRegion(uint(req.GetId()))
	if region == nil {
		return nil, status.Errorf(codes.NotFound, "Region not found")
	}
	return region.ToAPI(), nil
}
