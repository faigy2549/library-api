 
package handlers

import (
	"bytes"
	"encoding/json"
	"library-api/handlers"
	"library-api/models"
	"library-api/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

func TestCreateBook(t *testing.T) {
	repositories.ResetBooks()

	newBook := models.Book{
		Title:         "Programming",
		Author:        "Faigy",
		PublishedYear: 2025,
	}

	bookJSON, err := json.Marshal(newBook)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateBook)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	var createdBook models.Book
	err = json.NewDecoder(rr.Body).Decode(&createdBook)
	if err != nil {
		t.Fatal(err)
	}

	if createdBook.Title != newBook.Title {
		t.Errorf("Expected title %s, got %s", newBook.Title, createdBook.Title)
	}
}

func TestGetAllBooks(t *testing.T) {
	repositories.ResetBooks()

	book1 := models.Book{
		Title:         "Go Programming",
		Author:        "Faigy",
		PublishedYear: 2024,
	}
	book2 := models.Book{
		Title:         "Python Programming",
		Author:        "Faigy",
		PublishedYear: 2022,
	}

	repositories.AddBook(book1)
	repositories.AddBook(book2)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var returnedBooks []models.Book
	err = json.NewDecoder(rr.Body).Decode(&returnedBooks)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(returnedBooks) != 2 {
		t.Errorf("Expected 2 books, but got %d", len(returnedBooks))
	}

	if returnedBooks[0].Title != book1.Title || returnedBooks[1].Title != book2.Title {
		t.Errorf("Expected books %+v, but got %+v", []models.Book{book1, book2}, returnedBooks)
	}
}

func TestGetBookByID(t *testing.T) {
	repositories.ResetBooks()

	book := models.Book{
		Title:        "Programming",
		Author:       "Faigy",
		PublishedYear: 2023,
	}
	repositories.AddBook(book)


	req, err := http.NewRequest("GET", "/books/1", nil) 
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}


	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", handlers.GetBookByID).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var returnedBook models.Book
	err = json.NewDecoder(rr.Body).Decode(&returnedBook)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if returnedBook.ID != 1 || returnedBook.Title != book.Title || returnedBook.Author != book.Author {
		t.Errorf("Expected book %+v, got %+v", book, returnedBook)
	}
}
func TestDeleteBookByID(t *testing.T) {

	repositories.ResetBooks()

	book := models.Book{
		Title:         "Programming",
		Author:        "Faigy",
		PublishedYear: 2021,
	}

	repositories.AddBook(book)

	req, err := http.NewRequest("DELETE", "/books/1", nil) 
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBookByID).Methods("DELETE")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rr.Code)
	}

	if len(repositories.GetAllBooks()) != 0 {
		t.Errorf("Expected no books, but got %d", len(repositories.GetAllBooks()))
	}
}
