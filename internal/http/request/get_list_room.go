package request

type ListRoomRequest struct {
	SearchValue string `json:"search_value"`
	SearchBy    string `json:"search_by"`
	SortBy      string `json:"sort_by"`
	SortDir     string `json:"sort_dir"`
}
