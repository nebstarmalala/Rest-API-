package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
	
}

func updateBook(w http.ResponseWriter, r *http.Request){
	
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	
}