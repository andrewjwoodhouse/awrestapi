package main

// implement a RESTful API in Golang

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (model)

func main() {

	// init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
