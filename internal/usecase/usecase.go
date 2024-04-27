package usecase

import (
	"go.uber.org/zap"
	domain "room-meet/domain/repository"
	"room-meet/domain/usecase"
	"room-meet/internal/config"
)

type RoomUseCase struct {
	Logger         *zap.Logger
	RoomRepository domain.RoomRepository
	CacheService   domain.CacheRepository
	Super          config.BaseRepository
}

func NewRoomUseCase(
	logger *zap.Logger,
	userRepository domain.RoomRepository,
	cacheService domain.CacheRepository,
	super config.BaseRepository,
) usecase.RoomUseCase {
	return &RoomUseCase{Logger: logger, RoomRepository: userRepository, CacheService: cacheService, Super: super}
}
