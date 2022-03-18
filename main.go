package main

import (
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

var books []Book

func main() {
	// inti router
	r := mux.NewRouter()

	//books
	books = append(books, Book{ID: "1", Isbn: "123456", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}} )
	books = append(books, Book{ID: "2", Isbn: "678912", Title: "Book Two", Author: &Author{FirstName: "Jane", LastName: "Doe"}} )

	//route handlers(endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}