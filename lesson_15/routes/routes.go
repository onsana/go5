package routes

import (
	"library-server/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/available", handlers.GetAvailableBooks).Methods("GET")
	router.HandleFunc("/books/borrow/{id}", handlers.BorrowBook).Methods("POST")
}
