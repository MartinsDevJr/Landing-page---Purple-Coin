package main

import (
	"log"
	"net/http"

	"github.com/MartinsDevJr/purple_coin/routes"
	"github.com/gorilla/mux"

	"github.com/MartinsDevJr/purple_coin/db"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)

	log.Println("API rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}