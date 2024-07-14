package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	IsBorrowed bool   `json:"is_borrowed"`
}

var books = []Book{
	{ID: "1", Title: "Book One", Author: "Author One", IsBorrowed: false},
	{ID: "2", Title: "Book Two", Author: "Author Two", IsBorrowed: false},
	{ID: "3", Title: "Book Three", Author: "Author Three", IsBorrowed: false},
}

var mu sync.Mutex

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getAvailableBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var availableBooks []Book
	for _, book := range books {
		if !book.IsBorrowed {
			availableBooks = append(availableBooks, book)
		}
	}
	json.NewEncoder(w).Encode(availableBooks)
}

func borrowBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mu.Lock()
	defer mu.Unlock()
	for i, book := range books {
		if book.ID == params["id"] {
			if book.IsBorrowed {
				http.Error(w, "Book is already borrowed", http.StatusConflict)
				return
			}
			books[i].IsBorrowed = true
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/available", getAvailableBooks).Methods("GET")
	router.HandleFunc("/books/borrow/{id}", borrowBook).Methods("POST")

	http.ListenAndServe(":8000", router)
}
