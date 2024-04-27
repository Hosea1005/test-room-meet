package response

type Data struct {
	Token string `json:"token"`
}
type LoginResponse struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}
