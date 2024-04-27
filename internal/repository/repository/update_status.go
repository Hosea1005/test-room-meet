package repository

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/model"
	"time"
)

func (s *RoomRepository) UpdateStatus(ctx context.Context, data *entity.Room) (*entity.Room, error) {
	// Temukan ruangan berdasarkan ID dari data
	var room model.Room
	result := s.super.PostgresSql.First(&room, "id = ?", data.Id())
	if result.Error != nil {
		return nil, result.Error
	}

	// Perbarui atribut ruangan dengan data baru
	now := time.Now()
	room.Status = data.Status()
	room.UpdatedAt = &now

	// Simpan perubahan ke database
	result = s.super.PostgresSql.Save(&room)
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}
