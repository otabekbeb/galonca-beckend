package service

import (
	"context"
	"github.com/unrolled/render"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strconv"
	"tsl_server/api"
	"tsl_server/model"
)

type AuthService struct {
	api.UnimplementedAuthServiceServer
	storage *model.UserStorage
	manager *JWTManager
}

var as *AuthService

func NewAuthService() *AuthService {
	if as == nil {
		as = &AuthService{
			storage: model.NewUserStorage(),
			manager: NewManager(),
		}
	}
	return as
}

func (server *AuthService) Login(_ context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	user := server.storage.GetUserByEmail(req.GetUsername())
	if user == nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", req.GetUsername())
	}

	if !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.manager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &api.LoginResponse{AccessToken: token}
	return res, nil
}

type AuthInterceptor struct {
	jwtManager      *JWTManager
	storage         *model.UserStorage
	accessibleRoles map[string][]string
}

var ui *AuthInterceptor

func NewAuthInterceptor() *AuthInterceptor {
	const userServicePath = "/api.UserService/"
	const cargoServicePath = "/api.CargoService/"
	const truckServicePath = "/api.TruckService/"
	const serviceStationPath = "/api.ServiceStationService/"
	const sparePartsPath = "/api.SparePartService/"
	const roadsideServicePath = "/api.RoadsideServiceService/"
	const transportServicePath = "/api.TransportService/"
	const companyServicePath = "/api.CompanyService/"
	const profileServicePath = "/api.Profile/"
	if ui == nil {
		ui = &AuthInterceptor{
			jwtManager: NewManager(),
			storage:    model.NewUserStorage(),
			accessibleRoles: map[string][]string{
				userServicePath + "GetByEmail": {"admin", "user"},
				userServicePath + "WhoIs":      {"admin", "user"},
				cargoServicePath + "Create":    {"admin", "user"},
				cargoServicePath + "Update":    {"admin", "user"},
				cargoServicePath + "Delete":    {"admin", "user"},
				//cargoServicePath + "Find":                    {"admin", "user"},
				cargoServicePath + "Get":    {"admin", "user"},
				truckServicePath + "Create": {"admin", "user"},
				truckServicePath + "Update": {"admin", "user"},
				truckServicePath + "Delete": {"admin", "user"},
				//truckServicePath + "Find":                    {"admin", "user"},
				truckServicePath + "Get":      {"admin", "user"},
				serviceStationPath + "Create": {"admin", "user"},
				serviceStationPath + "Update": {"admin", "user"},
				serviceStationPath + "Delete": {"admin", "user"},
				//serviceStationPath + "Find":                  {"admin", "user"},
				serviceStationPath + "Get": {"admin", "user"},
				sparePartsPath + "Create":  {"admin", "user"},
				sparePartsPath + "Update":  {"admin", "user"},
				sparePartsPath + "Delete":  {"admin", "user"},
				//sparePartsPath + "Find":                      {"admin", "user"},
				sparePartsPath + "Get":         {"admin", "user"},
				roadsideServicePath + "Create": {"admin", "user"},
				roadsideServicePath + "Update": {"admin", "user"},
				roadsideServicePath + "Delete": {"admin", "user"},
				//roadsideServicePath + "Find":    {"admin", "user"},
				roadsideServicePath + "Get":     {"admin", "user"},
				transportServicePath + "Create": {"admin", "user"},
				transportServicePath + "Update": {"admin", "user"},
				transportServicePath + "Delete": {"admin", "user"},
				//transportServicePath + "Find":                {"admin", "user"},
				transportServicePath + "Get":        {"admin", "user"},
				companyServicePath + "Create":       {"admin", "user"},
				companyServicePath + "CreateReview": {"admin", "user"},
				companyServicePath + "Delete":       {"admin", "user"},
				companyServicePath + "ListReview":   {"admin", "user"},
				//companyServicePath + "Get":                   {"admin", "user"},
				//companyServicePath + "GetByUser":             {"admin", "user"},
				companyServicePath + "AddEmployee":           {"admin", "user"},
				companyServicePath + "ListEmployees":         {"admin", "user"},
				companyServicePath + "RemoveEmployee":        {"admin", "user"},
				profileServicePath + "ListMyCargo":           {"admin", "user"},
				profileServicePath + "ListMyTransport":       {"admin", "user"},
				profileServicePath + "ListMyRoadsideService": {"admin", "user"},
				profileServicePath + "ListMyServiceStation":  {"admin", "user"},
				profileServicePath + "ListMySpareParts":      {"admin", "user"},
				profileServicePath + "ListMyTruck":           {"admin", "user"},
				profileServicePath + "ListFavorites":         {"admin", "user"},
				profileServicePath + "AddFavorite":           {"admin", "user"},
				profileServicePath + "DeleteFavorite":        {"admin", "user"},
			},
		}
	}
	return ui
}

func NewWebAuthMiddleware() *AuthInterceptor {
	if ui == nil {
		ui = &AuthInterceptor{
			jwtManager: NewManager(),
			storage:    model.NewUserStorage(),
			accessibleRoles: map[string][]string{
				"imageManager/refresh":     {"admin", "user"},
				"/api.UserService/WhoIs":   {"admin", "user"},
				"/api.CargoService/Create": {"admin", "user"},
				"/api.CargoService/Update": {"admin", "user"},
				"/api.CargoService/Get":    {"admin", "user"},
				"/api.CargoService/Delete": {"admin", "user"},
				//"/api.CargoService/Find":              {"admin", "user"},
				"/api.TruckService/Create": {"admin", "user"},
				"/api.TruckService/Update": {"admin", "user"},
				"/api.TruckService/Get":    {"admin", "user"},
				"/api.TruckService/Delete": {"admin", "user"},
				//"/api.TruckService/Find":              {"admin", "user"},
				"/api.SparePartService/Create": {"admin", "user"},
				"/api.SparePartService/Update": {"admin", "user"},
				"/api.SparePartService/Get":    {"admin", "user"},
				"/api.SparePartService/Delete": {"admin", "user"},
				//"/api.SparePartService/Find":          {"admin", "user"},
				"/api.ServiceStationService/Create": {"admin", "user"},
				"/api.ServiceStationService/Update": {"admin", "user"},
				"/api.ServiceStationService/Get":    {"admin", "user"},
				"/api.ServiceStationService/Delete": {"admin", "user"},
				//"/api.ServiceStationService/Find":     {"admin", "user"},
				"/api.RoadsideServiceService/Create": {"admin", "user"},
				"/api.RoadsideServiceService/Update": {"admin", "user"},
				"/api.RoadsideServiceService/Get":    {"admin", "user"},
				"/api.RoadsideServiceService/Delete": {"admin", "user"},
				//"/api.RoadsideServiceService/Find":    {"admin", "user"},
				"/api.TransportService/Create":        {"admin", "user"},
				"/api.TransportServiceService/Update": {"admin", "user"},
				"/api.TransportServiceService/Get":    {"admin", "user"},
				"/api.TransportServiceService/Delete": {"admin", "user"},
				//"/api.TransportServiceService/Find":   {"admin", "user"},
				"/api.CompanyService/Create":       {"admin", "user"},
				"/api.CompanyService/CreateReview": {"admin", "user"},
				"/api.CompanyService/Delete":       {"admin", "user"},
				"/api.CompanyService/ListReview":   {"admin", "user"},
				//"/api.CompanyService/Get":             {"admin", "user"},
				//"/api.CompanyService/GetByUser":       {"admin", "user"},
				"/api.CompanyService/AddEmployee":    {"admin", "user"},
				"/api.CompanyService/ListEmployees":  {"admin", "user"},
				"/api.CompanyService/RemoveEmployee": {"admin", "user"},
				"/api.ProfileService/ListFavorites":  {"admin", "user"},
				"/api.ProfileService/AddFavorite":    {"admin", "user"},
				"/api.ProfileService/DeleteFavorite": {"admin", "user"},
			},
		}
	}
	return ui
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (error, uint64) {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	log.Printf("mthod: %s, ok: %v", method, ok)
	if !ok {
		log.Println("free method")
		return nil, 0
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided"), 0
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided"), 0
	}

	accessToken := values[0]
	log.Println("accessToken:", accessToken)
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err), 0
	}
	userId, err := strconv.ParseInt(claims.Id, 10, 64)
	user := interceptor.storage.GetUser(uint64(userId))
	if user == nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", userId), 0
	}
	for _, role := range accessibleRoles {
		if role == user.Role {
			return nil, uint64(userId)
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC"), uint64(userId)
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		err, userId := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(context.WithValue(ctx, "userId", userId), req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		err, _ := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) WebAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.URL.Path
		accessibleRoles, ok := interceptor.accessibleRoles[method]
		log.Printf("method: %s , ok: %v", method, ok)
		if !ok {
			// everyone can access
			next.ServeHTTP(w, r)
			return
		}
		format := render.New()
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			_ = format.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "access token not provided"})
			return
		}

		claims, err := interceptor.jwtManager.Verify(accessToken)
		if err != nil {
			_ = format.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "access token is invalid"})
			return
		}
		userId, err := strconv.ParseInt(claims.Id, 10, 64)
		user := interceptor.storage.GetUser(uint64(userId))
		if user == nil {
			_ = format.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "access token is invalid"})
			return
		}

		for _, role := range accessibleRoles {
			if role == user.Role {
				r.Header.Set("user-id", strconv.Itoa(int(user.ID)))
				next.ServeHTTP(w, r)
				return
			}
		}
		_ = format.JSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "no permission to access this RPC"})
	})
}
