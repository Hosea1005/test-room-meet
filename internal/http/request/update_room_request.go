package request

import "github.com/google/uuid"

type UpdateRoomRequest struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Location string    `json:"location"`
	Capacity int64     `json:"capacity"`
}
