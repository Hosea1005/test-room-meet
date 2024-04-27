package repository

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/model"
)

func (s *RoomRepository) InsertUser(ctx context.Context, data *entity.User) (*model.Users, error) {
	user := model.Users{
		ID:       helper.GeneratedUUID(),
		Name:     data.Name(),
		Password: data.Password(),
		Email:    data.Email(),
		Role:     "general",
	}
	result := s.super.PostgresSql.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
