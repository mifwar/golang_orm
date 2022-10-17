package routes

import (
	"gamestore/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterGameStore(r *mux.Router) {
	r.HandleFunc("/game/", controllers.CreateGame).Methods("POST")
	r.HandleFunc("/game/", controllers.GetGame).Methods("GET")
	r.HandleFunc("/game/{id}", controllers.GetGameByID).Methods("GET")
	r.HandleFunc("/game/{id}", controllers.DeleteGame).Methods("DELETE")
	r.HandleFunc("/game/{id}", controllers.UpdateGame).Methods("PUT")
}
