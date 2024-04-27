package http

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"room-meet/domain/entity"
	"room-meet/helper"
	"room-meet/internal/config"
	"room-meet/internal/http/request"
	"room-meet/internal/http/response"
)

func (a RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.UpdateRoomRequest
	)
	token := r.Header.Get("Authorization")
	claim, err := helper.CheckJWT(token, config.JWT_KEY)
	log.Println(err, claim)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusForbidden, response.Status{
			Code:    403,
			Message: "Invalid token",
		})
		return
	}
	mapClaims := *claim
	role, _ := mapClaims["role"].(string)
	if role != "admin" {
		helper.RespondWithJSON(w, http.StatusForbidden, response.Status{
			Code:    403,
			Message: "only admin can update status",
		})
		return
	}
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
