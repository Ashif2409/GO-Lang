package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request Body ", http.StatusBadRequest)
		return
	}
	book.ID = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			var updatedBook Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil {
				http.Error(w, "Invalid request Body ", http.StatusBadRequest)
				return
			}
			updatedBook.ID = id
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {

	books = append(books, Book{
		ID:   "1",
		Isbn: "2345678",
		Name: "Book one",
		Author: &Author{
			FirstName: "John",
			LastName:  "Smith",
		},
	})
	books = append(books, Book{
		ID:   "2",
		Isbn: "0987654",
		Name: "Book two",
		Author: &Author{
			FirstName: "Steve",
			LastName:  "Smith",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", r))
}
