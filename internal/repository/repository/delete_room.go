package repository

import (
	"context"
	"room-meet/domain/entity"
	"room-meet/internal/model"
	"time"
)

func (s *RoomRepository) DeleteRoom(ctx context.Context, id string) (*entity.Room, error) {
	// Temukan ruangan berdasarkan ID
	var room model.Room
	result := s.super.PostgresSql.First(&room, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Perbarui kolom deleted_at dengan waktu sekarang
	now := time.Now()
	room.DeletedAt = &now // Ubah properti secara langsung

	// Simpan perubahan ke database
	result = s.super.PostgresSql.Save(&room)
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}
