package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"room-meet/domain/usecase"
)

type RoomHandler struct {
	usecase usecase.RoomUseCase
}

func NewDeliveryHttpArea(r *mux.Router, usecase usecase.RoomUseCase) RoomHandler {
	handler := RoomHandler{
		usecase: usecase,
	}
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	//r.HandleFunc("/logout", handler.Logout).Methods("GET")
	api := r.PathPrefix("/room").Subrouter()
	api.HandleFunc("", handler.GetListRoom).Methods("GET")
	api.HandleFunc("/add-room", handler.CreateRoom).Methods("POST")
	api.HandleFunc("/delete-room/{id}", handler.DeleteRoom).Methods("DELETE")
	api.HandleFunc("/update-room/{id}", handler.UpdateRoom).Methods("PUT")
	api.HandleFunc("/update-status/{id}", handler.UpdateStatus).Methods("PUT")
	api1 := r.PathPrefix("/book").Subrouter()
	api1.HandleFunc("/request-room/{id}", handler.RequestRoom).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
	return handler
}
