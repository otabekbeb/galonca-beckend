package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"tsl_server/model"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

const secretKey = "!tslWebSite!"

func NewManager() *JWTManager {
	tokenDuration, _ := strconv.ParseInt(os.Getenv("token_duration_min"), 10, 32)
	return &JWTManager{secretKey, time.Duration(tokenDuration) * time.Minute}
}

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func GenerateResetPasswordToken(user *model.TslUser) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
			Id:        strconv.Itoa(int(user.ID)),
			Subject:   "смена пароля",
		},
		Email: user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyResetPasswordToken(accessToken string) (bool, string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return false, "", fmt.Errorf("invalid token: %w", err)
	}

	user, ok := token.Claims.(*UserClaims)
	if !ok {
		return false, "", fmt.Errorf("invalid token claims")
	}

	return true, user.Id, nil
}

func (manager *JWTManager) Generate(user *model.TslUser) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration * time.Minute).Unix(),
			Id:        strconv.Itoa(int(user.ID)),
			Subject:   "для доступа к сайту tsl",
		},
		Email: user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
