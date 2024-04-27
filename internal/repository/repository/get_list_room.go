package repository

import (
	"context"
	"errors"
	"fmt"
	"room-meet/internal/http/request"
	"room-meet/internal/model"
)

func (s *RoomRepository) GetListRoom(ctx context.Context, req request.ListRoomRequest) (*[]model.Room, error) {
	var products []model.Room
	db := s.super.PostgresSql
	// Memulai query dengan tabel produk
	query := db.Model(&model.Room{})

	// Pencarian berdasarkan SearchValue
	if req.SearchValue != "" && req.SearchBy != "" {
		// Jika tidak valid, itu adalah pencarian berdasarkan nama
		query = query.Where(req.SearchBy+" LIKE ?", "%"+req.SearchValue+"%")
	}

	// Pengurutan berdasarkan SortBy dan SortDir
	if req.SortBy != "" {
		order := req.SortBy
		if req.SortDir != "" {
			if req.SortDir != "DESC" {
				order += " ASC"
			} else {
				order += " DESC"

			}

		}
		query = query.Order(order)
	}

	// Eksekusi query
	result := query.Find(&products)
	sqlString := query.Statement.SQL.String()
	fmt.Println("Query SQL yang dijalankan:", sqlString, result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(products) == 0 {
		return nil, errors.New("rooms not found")
	}

	return &products, nil
}
