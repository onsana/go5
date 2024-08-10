package models

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	IsBorrowed bool   `json:"is_borrowed"`
}

var Books = []Book{
	{ID: "1", Title: "Book One", Author: "Author One", IsBorrowed: false},
	{ID: "2", Title: "Book Two", Author: "Author Two", IsBorrowed: false},
	{ID: "3", Title: "Book Three", Author: "Author Three", IsBorrowed: false},
}
