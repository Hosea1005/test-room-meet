package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/config"
	"room-meet/internal/http/request"
	"room-meet/internal/http/response"
)

func (a RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.CreateRoomRequest
	)
	token := r.Header.Get("Authorization")
	err := helper.CheckJWT(token, config.JWT_KEY)
	log.Println(err)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusForbidden, response.Status{
			Code:    403,
			Message: "Invalid token",
		})
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	payload, _ := entity.NewRoom(&entity.RoomDTO{
		Id:       helper.GeneratedUUID(),
		Name:     req.Name,
		Location: req.Location,
		Capacity: req.Capacity,
	})
	res := a.usecase.CreateRoom(r.Context(), payload)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
