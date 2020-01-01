package main

// implement a RESTful API in Golang

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// init books var as slice Book struct
var books []Book

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	// loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //MOCK ID - NOT SAFE!
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	// init router
	r := mux.NewRouter()

	// mock data - TODO: implement a database backend
	books = append(books, Book{ID: "1", Isbn: "11111111", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "22222222", Title: "Book Two", Author: &Author{Firstname: "Peter", Lastname: "Smith"}})

	// Route Handlers / Endpoints

	r.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/books", createBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
