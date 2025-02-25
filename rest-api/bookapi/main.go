package main

import (
	"log"
	"net/http"

	"example.com1/database"
	"example.com1/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
