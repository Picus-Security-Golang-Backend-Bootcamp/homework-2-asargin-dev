package library_models

// Created Struct for Book Informations

type Book struct {
	Id            int
	BookName      string
	StockCode     int
	IsbnNo        int
	NumberOfPages int
	Price         float32
	StockQuantity int
	AuthorInfo    Author
	IsDeleted     bool
}
