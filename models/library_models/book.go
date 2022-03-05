package library_models

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/utils"
)

// Created Struct for Book Informations

type Book struct {
	Id, StockCode, IsbnNo, NumberOfPages, StockQuantity int
	BookName                                            string
	Price                                               float32
	AuthorInfo                                          Author
	IsDeleted                                           bool
}

/* Constructor */
func (b Book) CreateLibrary(books_slice []string) []Book {

	var books []Book

	for index, value := range books_slice {
		var book Book
		book.Id = index + 1
		book.BookName = value
		book.StockCode = utils.CreateRandomNumberGenerator(16)
		book.IsbnNo = 1
		book.StockQuantity = 1
		book.Price = 1
		book.IsDeleted = false
		book.AuthorInfo = Author{FirstName: "", LastName: ""}
		books = append(books, book)
	}

	return books

}
