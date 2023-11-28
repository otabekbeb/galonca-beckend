package service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"html"
	"log"
	"net/http"
	"tsl_server/api"
	"tsl_server/model"
)

type ReferenceService struct {
	api.UnimplementedReferenceServiceServer
	storage *model.ReferenceStorage
}

func NewServiceReferenceService() *ReferenceService {
	return &ReferenceService{
		storage: model.NewReferenceStorage(),
	}
}

func (server *ReferenceService) Get(_ context.Context, req *api.ReferenceRequest) (*api.Reference, error) {
	switch req.GetType() {
	case api.RType_cargoType:
		result := server.storage.GetCargoType(req.GetName())
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_transportType:
		result := server.storage.GetTransportType(req.GetName(), 0)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_busTransportType:
		result := server.storage.GetTransportType(req.GetName(), 1)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_specialTransportType:
		result := server.storage.GetTransportType(req.GetName(), 2)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_autoBrand:
		result := server.storage.GetBrand(req.GetName(), 0)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_busBrand:
		result := server.storage.GetBrand(req.GetName(), 1)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_specialBrand:
		result := server.storage.GetTransportType(req.GetName(), 3)
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_loadingType:
		result := server.storage.GetLoadingType(req.GetName())
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_addition:
		result := server.storage.GetAddition(req.GetName())
		if result == nil {
			break
		}
		return result.ToAPI(), nil
	case api.RType_serviceCategory:
		result := server.storage.GetServiceCategoryByName(req.GetName())
		if result == nil {
			break
		}
		return result.ToAPI(), nil

	}
	return nil, status.Errorf(codes.NotFound, "Reference not found")
}

func (server *ReferenceService) Find(_ context.Context, req *api.ReferenceListRequest) (*api.ReferenceListResponse, error) {
	var result api.ReferenceListResponse
	switch req.GetType() {
	case api.RType_cargoType:
		references := server.storage.GetCargoTypeList(api.SortType_name[int32(req.GetSort())])
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_transportType:
		references := server.storage.GetTransportTypeList(api.SortType_name[int32(req.GetSort())], 0)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_busTransportType:
		references := server.storage.GetTransportTypeList(api.SortType_name[int32(req.GetSort())], 1)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_specialTransportType:
		references := server.storage.GetTransportTypeList(api.SortType_name[int32(req.GetSort())], 2)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_autoBrand:
		references := server.storage.GetBrandList(api.SortType_name[int32(req.GetSort())], 0)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_busBrand:
		references := server.storage.GetBrandList(api.SortType_name[int32(req.GetSort())], 1)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_specialBrand:
		references := server.storage.GetBrandList(api.SortType_name[int32(req.GetSort())], 2)
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_loadingType:
		references := server.storage.GetLoadingTypeList(api.SortType_name[int32(req.GetSort())])
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_addition:
		references := server.storage.GetAdditionList(api.SortType_name[int32(req.GetSort())])
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	case api.RType_serviceCategory:
		references := server.storage.GetServiceCategoryList(api.SortType_name[int32(req.GetSort())])
		for _, c := range references {
			result.Reference = append(result.Reference, c.ToAPI())
		}
		result.Found = uint32(len(result.Reference))
		return &result, nil
	}
	return nil, nil
}

func (server *ReferenceService) AddLike(ctx context.Context, req *api.Like) (*api.LikeResponse, error) {
	p, _ := peer.FromContext(ctx)
	server.storage.AddLike(req, p.Addr.String())
	lr := api.LikeRequest{
		EntityId:   req.EntityId,
		EntityType: html.EscapeString(req.EntityType),
	}
	return server.storage.GetLikes(&lr), nil
}

func (server *ReferenceService) GetLikes(_ context.Context, req *api.LikeRequest) (*api.LikeResponse, error) {
	return server.storage.GetLikes(req), nil
}

func (server *ReferenceService) Stat(_ context.Context, req *api.StatRequest) (*api.StatResponse, error) {
	url := "https://stat.gov.kz/api/juridical/counter/api/?lang=ru&bin=" + req.GetBin()
	resp, err := http.Get(url)
	log.Printf("result: %d %s", resp.StatusCode, resp.Status)
	if resp.StatusCode != http.StatusOK {
		return &api.StatResponse{Success: false, Description: resp.Status}, nil
	}
	if err != nil {
		return &api.StatResponse{Success: false, Description: err.Error()}, err
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if result["success"].(bool) == false {
		return &api.StatResponse{Success: false, Description: "BIN not found"}, nil
	}
	obj := result["obj"].(map[string]interface{})
	if err != nil {
		return &api.StatResponse{Success: false, Description: err.Error()}, err
	}
	response := &api.StatResponse{
		Success: result["success"].(bool),
		Obj: &api.O{
			Id:          int64(obj["id"].(float64)),
			Bin:         html.EscapeString(obj["bin"].(string)),
			Name:        html.EscapeString(obj["name"].(string)),
			OkedCode:    html.EscapeString(obj["okedCode"].(string)),
			OkedName:    html.EscapeString(obj["okedName"].(string)),
			KatoCode:    html.EscapeString(obj["katoCode"].(string)),
			KatoId:      int64(obj["katoId"].(float64)),
			KatoAddress: html.EscapeString(obj["katoAddress"].(string)),
			Fio:         html.EscapeString(obj["fio"].(string)),
		},
	}
	if obj["registerDate"] == nil {
		response.Obj.RegisterDate = "не указана"
	} else {
		response.Obj.RegisterDate = html.EscapeString(obj["registerDate"].(string))
	}

	if obj["secondOkeds"] == nil {
		response.Obj.SecondOkeds = "не указан"
	} else {
		response.Obj.SecondOkeds = html.EscapeString(obj["secondOkeds"].(string))
	}

	if obj["krpCode"] == nil {
		response.Obj.KrpCode = "не указан"
	} else {
		response.Obj.KrpCode = html.EscapeString(obj["krpCode"].(string))
	}
	if obj["krpName"] == nil {
		response.Obj.KrpName = "не указан"
	} else {
		response.Obj.KrpName = html.EscapeString(obj["krpName"].(string))
	}
	if obj["krpBfCode"] == nil {
		response.Obj.KrpBfCode = "не указан"
	} else {
		response.Obj.KrpBfCode = html.EscapeString(obj["krpCode"].(string))
	}
	if obj["krpBfName"] == nil {
		response.Obj.KrpBfName = "не указан"
	} else {
		response.Obj.KrpBfName = html.EscapeString(obj["krpBfName"].(string))
	}

	return response, nil
}
