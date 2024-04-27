package repository

import (
	"go.uber.org/zap"
	"room-meet/internal/config"
)

type RoomRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}

func NewRoomRepository(super config.BaseRepository,
	logger *zap.Logger) *RoomRepository {
	return &RoomRepository{
		super:  super,
		logger: logger,
	}
}
func NewCacheRepository(super config.BaseRepository,
	logger *zap.Logger) *CacheRepository {
	return &CacheRepository{
		super:  super,
		logger: logger,
	}
}

type CacheRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}
