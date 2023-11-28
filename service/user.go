package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"net/smtp"
	"time"
	"tsl_server/api"
	"tsl_server/model"
)

const secretCaptcha = "6LflwQMfAAAAAGdBOsOB64L21nwBzmUWAZMx0yg0"

type UserService struct {
	api.UnimplementedUserServiceServer
	storage *model.UserStorage
}

func NewUserService() *UserService {
	return &UserService{
		storage: model.NewUserStorage(),
	}
}

func (server *UserService) Create(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &api.CreateUserResponse{
			Id: 0,
		}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["captchatoken"]
	if len(values) == 0 {
		return &api.CreateUserResponse{
			Id: 0,
		}, status.Errorf(codes.Unauthenticated, "token is not provided")
	}
	captcha, err := recaptcha.NewReCAPTCHA(secretCaptcha, recaptcha.V3, 10*time.Second)
	if err != nil {
		return &api.CreateUserResponse{
			Id: 0,
		}, status.Errorf(codes.Unauthenticated, err.Error())
	}
	err = captcha.VerifyWithOptions(values[0], recaptcha.VerifyOption{Action: "submit", Threshold: 0.8})
	if err != nil {
		return &api.CreateUserResponse{
			Id: 0,
		}, status.Errorf(codes.Unauthenticated, err.Error())
	}
	user, err := server.storage.Create(uint64(req.GetId()), req.GetName(), req.GetPassword(), req.GetEmail(), req.GetPhone(), "admin")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user %v", err)
	}
	res := &api.CreateUserResponse{Id: uint64(user.ID)}
	return res, nil
}

func (server *UserService) Get(_ context.Context, req *api.GetUserRequest) (*api.TslUser, error) {
	user := server.storage.GetUser(req.GetId())
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found %v", req.GetId())
	}
	res := user.ToAPI()
	return res, nil
}

func (server *UserService) Delete(_ context.Context, req *api.DeleteUserRequest) (*api.CommonUserResponse, error) {
	result := server.storage.Delete(req.GetId())
	return result, nil
}

func (server *UserService) GetByEmail(_ context.Context, req *api.GetUserByEmailRequest) (*api.TslUser, error) {
	user := server.storage.GetUserByEmail(req.GetEmail())
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found %v", req.GetEmail())
	}
	res := user.ToAPI()
	return res, nil
}

func (server *UserService) ChangePassword(_ context.Context, req *api.ChangePasswordRequest) (*api.CommonUserResponse, error) {
	result := server.storage.ChangePassword(req)
	return result, nil
}

func (server *UserService) WhoIs(ctx context.Context, _ *api.WhoIsRequest) (*api.TslUser, error) {
	userId := ctx.Value("userId").(uint64)
	user := server.storage.GetUser(userId)
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found %v", userId)
	}
	res := user.ToAPI()
	return res, nil
}

func (server *UserService) CheckResetPasswordToken(ctx context.Context, _ *api.WhoIsRequest) (*api.CommonUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "metadata is not provided",
		}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["resetpasswordtoken"]
	if len(values) == 0 {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "token is not provided",
		}, status.Errorf(codes.Unauthenticated, "token is not provided")
	}

	accessToken := values[0]
	result, id, _ := VerifyResetPasswordToken(accessToken)
	return &api.CommonUserResponse{
		Success: result,
		Error:   id,
	}, nil
}

func (server *UserService) ResetPasswordEmail(ctx context.Context, req *api.GetUserByEmailRequest) (*api.CommonUserResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "token is not provided",
		}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["captchatoken"]
	if len(values) == 0 {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "token is not provided",
		}, status.Errorf(codes.Unauthenticated, "token is not provided")
	}
	captcha, err := recaptcha.NewReCAPTCHA(secretCaptcha, recaptcha.V3, 10*time.Second)
	if err != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   err.Error(),
		}, status.Errorf(codes.Unauthenticated, err.Error())
	}
	err = captcha.VerifyWithOptions(values[0], recaptcha.VerifyOption{Action: "submit", Threshold: 0.8})
	if err != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   err.Error(),
		}, status.Errorf(codes.Unauthenticated, err.Error())
	}
	user := server.storage.GetUserByEmail(req.GetEmail())
	if user == nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "user not found",
		}, status.Errorf(codes.NotFound, "user not found")
	}
	token, err := GenerateResetPasswordToken(user)
	if err != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "can not generate token",
		}, err
	}
	to := []string{
		user.Email,
	}
	auth := smtp.PlainAuth("", "admin@tsl.kz", "@tsl_kz!", "smtp.mail.ru")
	subject := "Subject: TSL.KZ reset password\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><p>Для смены пароля пройдите по этой</p><a href=\"https://tsl.kz/rp.html?t=" + token + "\">ссылке</a></body></html>"
	msg := []byte(subject + mime + body)
	err = smtp.SendMail("smtp.mail.ru:25", auth, "admin@tsl.kz", to, msg)
	if err != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}
	return &api.CommonUserResponse{
		Success: true,
		Error:   "",
	}, nil
}
