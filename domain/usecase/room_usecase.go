package usecase

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/http/request"
	"room-meet/internal/http/response"
)

type RoomUseCase interface {
	UpdateRoom(ctx context.Context, request *entity.Room) *response.UpdateRoomResponse
	UpdateStatus(ctx context.Context, request *entity.Room) *response.UpdateRoomResponse
	DeleteRoom(ctx context.Context, request string) *response.DeleteResponse
	Register(ctx context.Context, request *entity.User) *response.RegisterResponse
	Login(ctx context.Context, request *entity.User) *response.LoginResponse
	CreateRoom(ctx context.Context, request *entity.Room) *response.CreateRoomResponse
	GetListRoom(ctx context.Context, request request.ListRoomRequest) *response.ListRoomsResponse
}
