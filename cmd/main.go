package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	dUsecase "room-meet/domain/usecase"
	"room-meet/internal/config"
	httpHandler "room-meet/internal/http"
	"room-meet/internal/repository/repository"
	"room-meet/internal/usecase"
)

func main() {
	logger := config.Logger()
	config.Environment()
	logger.Info("initializing service product ")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	repo := config.BaseRepository{
		PostgresSql: config.PostgresDB(logger),
		Redis:       config.Redis(ctx, logger),
	}
	// Initialize repository instances
	userRepo := repository.NewRoomRepository(repo, logger)
	cacheRepo := repository.NewCacheRepository(repo, logger)
	// Initialize use case instances
	userService := usecase.NewRoomUseCase(logger, userRepo, cacheRepo, repo)

	r := mux.NewRouter()
	initHandler(r, userService)
	http.Handle("/", r)

}
func initHandler(r *mux.Router, usecase dUsecase.RoomUseCase) {
	httpHandler.NewDeliveryHttpArea(r, usecase)
}
