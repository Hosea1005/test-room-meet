package response

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type RegisterResponse struct {
	Status Status `json:"status"`
}

