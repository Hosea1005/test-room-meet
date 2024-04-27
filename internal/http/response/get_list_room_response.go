package response

import (
	"github.com/google/uuid"
	"time"
)

type ListRoomsResponse struct {
	Status Status  `json:"status"`
	Data   []Rooms `json:"data"`
}

type Rooms struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Capacity  int64      `json:"capacity"`
	CreatedAt *time.Time `json:"created_at"`
}
