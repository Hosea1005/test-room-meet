package domain

import (
	"context"
	"github.com/google/uuid"
	"room-meet/domain/entity"
	"room-meet/internal/http/request"
	"room-meet/internal/model"
)

//go:generate mockery --name ProductRepository --filename product_mock.go --inpackage --structname ProductMock --output=./
type RoomRepository interface {
	//User
	InsertUser(ctx context.Context, data *entity.User) (*model.Users, error)
	CheckDataUser(ctx context.Context, email string) (*model.Users, error)

	GetListRoom(ctx context.Context, req request.ListRoomRequest) (*[]model.Room, error)
	CheckIdProduct(ctx context.Context, id uuid.UUID) error
	InsertRoom(ctx context.Context, data *entity.Room) (*entity.Room, error)
	UpdateRoom(ctx context.Context, data *entity.Room) (*entity.Room, error)
	UpdateStatus(ctx context.Context, data *entity.Room) (*entity.Room, error)
	DeleteRoom(ctx context.Context, id string) (*entity.Room, error)
}
