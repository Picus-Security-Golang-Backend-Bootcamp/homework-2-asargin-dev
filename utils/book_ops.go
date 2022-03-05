package utils

import (
	"fmt"
	"strings"

	library "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/models/library_models"
)

func Buy_book() {

}

func Delete_book() {

}

// The function that gets all books in the library.
func List_book(books []library.Book) {

	if len(books) == 0 {

		fmt.Println("Kitaplıkta kitap bulunmamaktadır.")

	} else {

		for _, book := range books {

			fmt.Printf("%d - %s\n", book.Id, book.BookName)

		}

	}

}

// The function that searches books within written arg inputs
func Search_book(search_params string, books []library.Book) {

	if len(search_params) == 0 {

		fmt.Println("Lütfen aramak istediğiniz kitabın adını yazınız.")

	} else {
		// Sonuçlar liste şeklinde istenildiğinden, İşlemler sonrası eşleşen sonuçları 'result' slice'ının içine append edeceğiz.
		result := []string{}

		for _, book := range books {
			//Büyük-küçük harf duyarlılığını kaldırmak için hem kitapları hem de arama parametrelerini lowercase'e çeviriyoruz.
			if strings.Contains(strings.ToLower(book.BookName), strings.ToLower(search_params)) {
				result = append(result, book.BookName)
			}
		}

		if len(result) == 0 {
			fmt.Println("Aradığınız kriterde kitap bulunamamıştır.")
		} else {
			fmt.Println(result)
		}

	}
}
