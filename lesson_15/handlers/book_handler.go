package handlers

import (
	"encoding/json"
	"net/http"

	"library-server/services"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := services.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	availableBooks := services.GetAvailableBooks()
	json.NewEncoder(w).Encode(availableBooks)
}

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	book, err := services.BorrowBook(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(book)
}
