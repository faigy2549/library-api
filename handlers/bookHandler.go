package handlers

import (
	"encoding/json"
	"library-api/models"
	"library-api/repositories"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newBook = repositories.AddBook(newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	if author != "" {
		books := repositories.SearchBooksByAuthor(author)
		if len(books) == 0 {
			http.Error(w, "No books found for author: "+author, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(books)
		return
	}

	books := repositories.GetAllBooks()
	if len(books) == 0 {
		http.Error(w, "No books found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book := repositories.GetBookByID(id)
	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	json.NewEncoder(w).Encode(book)
}


func DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if !repositories.DeleteBookByID(id) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
