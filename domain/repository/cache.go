package domain

import (
	"context"
	"room-meet/domain/entity"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value *entity.Room) error
	SetToken(ctx context.Context, key string, token string) error
	GetToken(ctx context.Context, key string) (string, error)
	DeleteToken(ctx context.Context, key string) error
}
