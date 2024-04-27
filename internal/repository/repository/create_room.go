package repository

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/model"
)

func (s *RoomRepository) InsertRoom(ctx context.Context, data *entity.Room) (*entity.Room, error) {
	room := model.Room{
		ID:        data.Id(),
		Name:      data.Name(),
		Location:  data.Location(),
		Status:    data.Status(),
		Capacity:  data.Capacity(),
		CreatedAt: data.CreatedAt(),
	}
	result := s.super.PostgresSql.Create(&room)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
