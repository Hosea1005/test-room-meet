package usecase

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/config"
	"room-meet/internal/http/response"
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

	token, _ := helper.GenerateJWT(config.JWT_KEY, resUser.ID.String(), resUser.Email, resUser.Role)
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
