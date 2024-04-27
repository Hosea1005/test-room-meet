package usecase

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/http/response"
)

func (a RoomUseCase) UpdateRoom(ctx context.Context, request *entity.Room) *response.UpdateRoomResponse {
	_, err := a.RoomRepository.UpdateRoom(ctx, request)
	if err != nil {
		return &response.UpdateRoomResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &response.UpdateRoomResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Update Status Room",
		},
	}
}
