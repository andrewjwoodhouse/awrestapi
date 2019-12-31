package main

import "github.com/gorilla/mux"

func main() {

	// init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

}
