package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	RoomDTO struct {
		Id        uuid.UUID
		Name      string
		Location  string
		Status    string
		Capacity  int64
		CreatedAt *time.Time
		DeletedAt *time.Time
	}
	Room struct {
		id        uuid.UUID
		name      string
		location  string
		status    string
		capacity  int64
		createdAt *time.Time
		deletedAt *time.Time
	}
)

func NewRoom(payload *RoomDTO) (*Room, error) {
	return RebuildRoom(payload), nil
}
func RebuildRoom(payload *RoomDTO) *Room {
	return &Room{
		id:        payload.Id,
		name:      payload.Name,
		location:  payload.Location,
		status:    payload.Status,
		capacity:  payload.Capacity,
		createdAt: payload.CreatedAt,
		deletedAt: payload.DeletedAt,
	}
}

func (p *Room) Id() uuid.UUID {
	return p.id
}
func (p *Room) Name() string {
	return p.name
}
func (p *Room) Location() string {
	return p.location
}
func (p *Room) Status() string {
	return p.status
}
func (p *Room) SetStatus(req string) {
	p.status = req
}
func (p *Room) Capacity() int64 {
	return p.capacity
}
func (p *Room) CreatedAt() *time.Time {
	return p.createdAt
}

//	func (p *Room) UpdatedAt() *time.Time {
//		return p.updatedAt
//	}
func (p *Room) DeletedAt() *time.Time {
	return p.deletedAt
}
func (p *Room) SetDeletedAt(req *time.Time) {
	p.deletedAt = req
}
