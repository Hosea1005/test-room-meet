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

func (a RoomHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var (
		req request.UpdateStatusRequest
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
	vars := mux.Vars(r)
	log.Println(vars["id"])
	req.Id, _ = helper.ConvertStringToUUID(vars["id"])
	payload, _ := entity.NewRoom(&entity.RoomDTO{
		Id:     req.Id,
		Status: req.Status,
	})
	res := a.usecase.UpdateStatus(r.Context(), payload)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
