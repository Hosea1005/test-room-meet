package http

import (
	"net/http"
	"room-meet/helper"
	"room-meet/internal/http/request"
)

func (a RoomHandler) GetListRoom(w http.ResponseWriter, r *http.Request) {
	var (
		req request.ListRoomRequest
	)
	req.SearchValue = r.FormValue("search_value")
	req.SearchBy = r.FormValue("search_by")
	req.SortBy = r.FormValue("sort_by")
	req.SortDir = r.FormValue("sort_dir")
	res := a.usecase.GetListRoom(r.Context(), req)
	helper.RespondWithJSON(w, http.StatusOK, res)
}
