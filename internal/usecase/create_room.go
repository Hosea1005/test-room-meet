package usecase

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/http/response"
)

func (a RoomUseCase) CreateRoom(ctx context.Context, request *entity.Room) *response.CreateRoomResponse {
	request.SetStatus("available")
	_, err := a.RoomRepository.InsertRoom(ctx, request)
	if err != nil {
		return &response.CreateRoomResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &response.CreateRoomResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Create Room",
		},
	}
}
