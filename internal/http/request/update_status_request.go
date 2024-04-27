package request

import "github.com/google/uuid"

type UpdateStatusRequest struct {
	Id     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}
