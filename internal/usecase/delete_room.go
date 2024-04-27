package usecase

import (
	"context"
	"room-meet/internal/http/response"
)

func (a RoomUseCase) DeleteRoom(ctx context.Context, request string) *response.DeleteResponse {
	_, err := a.RoomRepository.DeleteRoom(ctx, request)
	if err != nil {
		return &response.DeleteResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &response.DeleteResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Delete Room",
		},
	}
}
