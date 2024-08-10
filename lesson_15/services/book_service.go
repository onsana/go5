package services

import (
	"errors"
	"sync"

	"library-server/models"
)

var mu sync.Mutex

func GetAllBooks() []models.Book {
	return models.Books
}

func GetAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range models.Books {
		if !book.IsBorrowed {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func BorrowBook(id string) (*models.Book, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, book := range models.Books {
		if book.ID == id {
			if book.IsBorrowed {
				return nil, errors.New("book is already borrowed")
			}
			models.Books[i].IsBorrowed = true
			return &models.Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}
