package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"room-meet/internal/model"
)

func (s *RoomRepository) CheckIdProduct(ctx context.Context, id uuid.UUID) error {
	var bank model.Room
	result := s.super.PostgresSql.Where("id = ?", id).First(&bank)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("Bank not found")
		}
		return result.Error
	}

	// Jika ID bank ditemukan di database, return nil (tanpa kesalahan).
	return nil
}
