package http

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/http/request"
)

func (a RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.UpdateRoomRequest
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, []error{errors.New("unauthorized")})
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	log.Println(vars["id"])
	req.Id, _ = helper.ConvertStringToUUID(vars["id"])
	payload, _ := entity.NewRoom(&entity.RoomDTO{
		Id:       req.Id,
		Name:     req.Name,
		Location: req.Location,
		Status:   req.Status,
		Capacity: req.Capacity,
	})
	res := a.usecase.UpdateRoom(r.Context(), payload)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
