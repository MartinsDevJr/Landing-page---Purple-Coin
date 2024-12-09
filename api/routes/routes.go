package routes

import (
	"github.com/MartinsDevJr/purple_coin/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
}