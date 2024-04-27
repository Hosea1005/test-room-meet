package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"room-meet/helper"
	"room-meet/internal/http/request"
)

func (a RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.DeleteRequest
	)
	vars := mux.Vars(r)
	log.Println(vars["id"])
	req.Id = vars["id"]

	res := a.usecase.DeleteRoom(r.Context(), req.Id)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
