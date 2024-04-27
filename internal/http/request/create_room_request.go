package request

type CreateRoomRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int64  `json:"capacity"`
}
