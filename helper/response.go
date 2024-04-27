package helper

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}
