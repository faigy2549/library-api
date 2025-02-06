package main

import (
	"library-api/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET") 
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.DeleteBookByID).Methods("DELETE")
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
