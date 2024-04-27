package usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/config"
	"room-meet/internal/http/response"
	"time"
)

func (a RoomUseCase) Login(ctx context.Context, request *entity.User) *response.LoginResponse {
	resUser, err := a.RoomRepository.CheckDataUser(ctx, request.Email())
	if err != nil {
		return &response.LoginResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	errVal := helper.CheckValidatePassword(resUser.Password, request.Password())
	if errVal != nil {
		return &response.LoginResponse{
			Status: response.Status{
				Code:    400,
				Message: "Password Incorrect",
			},
		}
	}
	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 10)
	claims := &config.JWTClaim{
		UserId: resUser.ID,
		Email:  resUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "technical-test",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// medeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		return &response.LoginResponse{
			Status: response.Status{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	err = a.CacheService.SetToken(ctx, token, token)
	if err != nil {
		return &response.LoginResponse{
			Status: response.Status{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	return &response.LoginResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Login",
		},
		Data: response.Data{
			Token: token,
		},
	}
}
