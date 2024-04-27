package usecase

import (
	"context"
	"room-meet/internal/http/request"
	"room-meet/internal/http/response"
)

func (a RoomUseCase) GetListRoom(ctx context.Context, request request.ListRoomRequest) *response.ListRoomsResponse {
	resProduct, err := a.RoomRepository.GetListRoom(ctx, request)
	if err != nil {
		return &response.ListRoomsResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	var roomData []response.Rooms
	for _, product := range *resProduct {
		roomData = append(roomData, response.Rooms{
			ID:        product.ID,
			Name:      product.Name,
			Location:  product.Location,
			Capacity:  product.Capacity,
			CreatedAt: product.CreatedAt,
		})
	}

	return &response.ListRoomsResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success",
		},

		Data: roomData,
	}
}
