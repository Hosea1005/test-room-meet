package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"room-meet/helper"
	"room-meet/internal/config"
	"room-meet/internal/http/request"
	"room-meet/internal/http/response"
)

func (a RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.DeleteRequest
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
	vars := mux.Vars(r)
	log.Println(vars["id"])
	req.Id = vars["id"]

	res := a.usecase.DeleteRoom(r.Context(), req.Id)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
