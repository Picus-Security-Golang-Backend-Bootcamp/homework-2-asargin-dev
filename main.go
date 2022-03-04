package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/utils"
)

func init() {
}

func main() {

	args := os.Args

	books := []string{
		"A Tale of Two Cities",
		"The Hobbit",
		"The Little Prince",
		"The Alchemist",
		"Harry Potter and the Chamber of Secrets",
	}

	//Checking args and routing associated functions
	if len(args) > 1 {

		switch args[1] {

		case "buy":
			fmt.Println("Uygulama Yapım Aşamasındadır")

		case "delete":
			fmt.Println("Uygulama Yapım Aşamasındadır")

		case "list":

			utils.List_book(books)

		case "search":
			// search argümanı sonrası girilececek değerleri Join ederek string bir arama querysi oluşturarak
			// kitaplarla beraber search_book fonksiyonumuza parametre olarak gönderiyoruz.
			utils.Search_book(strings.Join(args[2:], " "), books)

		}

	} else {
		fmt.Println(

			`Hoşgeldiniz, *Kitapp* uygulamasında kullanabileceğiniz komutlar :

			Kitap satın almak için => buy 

			Kitap aramak için => search  

			Kitapları listelemek için => list 

			Kitaplıktan kitap silmek için => delete`)
	}

}
