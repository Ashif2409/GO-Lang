package controllers

import (
	"encoding/json"
	"net/http"

	"example.com1/database"
	models "example.com1/model"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []models.Book
	if err := database.DB.Preload("Author").Find(&books).Error; err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var book models.Book
	if err := database.DB.Preload("Author").First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the author exists
	var author models.Author
	if err := database.DB.Where("first_name = ? AND last_name = ?", book.Author.FirstName, book.Author.LastName).First(&author).Error; err != nil {
		// If author not found, create a new one
		author = models.Author{
			FirstName: book.Author.FirstName,
			LastName:  book.Author.LastName,
		}
		if err := database.DB.Create(&author).Error; err != nil {
			http.Error(w, "Failed to create author", http.StatusInternalServerError)
			return
		}
	}

	// Assign the AuthorID to the Book and save it
	book.AuthorID = author.ID
	if err := database.DB.Create(&book).Error; err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Ensure new author exists before proceeding
	if newBook.AuthorID != 0 {
		var author models.Author
		if err := database.DB.First(&author, newBook.AuthorID).Error; err != nil {
			http.Error(w, "New Author not found", http.StatusBadRequest)
			return
		}
	}

	// Delete the existing book
	if err := database.DB.Delete(&book).Error; err != nil {
		http.Error(w, "Failed to delete old book", http.StatusInternalServerError)
		return
	}

	// Assign the same ID to the new book
	newBook.ID = book.ID

	// Recreate the book with the same ID
	if err := database.DB.Create(&newBook).Error; err != nil {
		http.Error(w, "Failed to recreate book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newBook)
}
