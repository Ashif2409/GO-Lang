package routes

import (
	"example.com1/controllers"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes all book-related routes
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PATCH")
}
