package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"room-meet/domain/entity"
	"room-meet/internal/http/response"
)

func (a RoomUseCase) Register(ctx context.Context, request *entity.User) *response.RegisterResponse {
	// hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password()), bcrypt.DefaultCost)
	request.SetPassword(string(hashPassword))
	user, err := a.RoomRepository.InsertUser(ctx, request)
	//a.Logger.Info(user)
	log.Println(user)
	if err != nil {
		return &response.RegisterResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &response.RegisterResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Register",
		},
	}
}
