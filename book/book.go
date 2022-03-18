package book

import (
	"fmt"
	"strings"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/author"
)

//Create Slice that contains Book type
var bookList = []Book{
	{BookID: 1, Name: "Harry Potter ve Felsefe Taşı", Page: 258, StockCount: 5, Price: 32.00, StockID: 111, ISBN: 12345678910, Author: author.Author{ID: 745, Name: "J.K. Rowling"}},
	{BookID: 2, Name: "Yüzüklerin Efendisi Yüzük Kardeşliği", Page: 321, StockCount: 15, Price: 62.54, StockID: 112, ISBN: 2341567890, Author: author.Author{ID: 654, Name: "J.R.R. Tolkien"}},
	{BookID: 3, Name: "The Return of the King: The Lord of the Rings", Page: 178, StockCount: 1, Price: 139.00, StockID: 555, ISBN: 1236574897, Author: author.Author{ID: 654, Name: "J.R.R. Tolkien"}},
	{BookID: 4, Name: "Açlık", Page: 132, StockCount: 22, Price: 16.50, StockID: 432, ISBN: 456789654, Author: author.Author{ID: 222, Name: "Knut Hamsun"}},
	{BookID: 5, Name: "Satranç", Page: 111, StockCount: 12, Price: 6.30, StockID: 889, ISBN: 8796578904, Author: author.Author{ID: 155, Name: "Stefan Zweig"}},
}

// Create Book type
type Book struct {
	BookID     int
	Name       string
	Page       int
	StockCount int
	Price      float64
	StockID    int
	ISBN       int
	Author     author.Author
}

func (b Book) Create() Book {

	return b
}

func (b Book) GetList() []Book {
	return bookList
}

func (b Book) AddList() {
	bookList = append(bookList, b)
}

func (b Book) SearchById(id int) Book {
	var selectedBook Book
	for _, v := range bookList {
		if v.BookID == id {
			selectedBook = v
		}
	}
	return selectedBook
}

func (b Book) SearchByName(name string) []Book {
	newBookList := []Book{}
	for _, v := range bookList {
		lowerBookName := strings.ToLower(v.Name)
		lowerName := strings.ToLower(name)
		if strings.Contains(lowerBookName, lowerName) {
			newBookList = append(newBookList, v)
		}
	}
	return newBookList
}

func (b Book) DeleteById(id int) {
	var deletedBook Book
	for i, v := range bookList {
		if v.BookID == id {
			deletedBook = v
			bookList = append(bookList[:i], bookList[i+1:]...)
			break
		}
	}
	turkishLira := string(rune(8378))
	if deletedBook.BookID == 0 {
		fmt.Println("Bu Id'de bir kitap bulunmamaktadır")
	} else {
		book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", deletedBook.BookID, deletedBook.Name, deletedBook.Page, deletedBook.StockCount, deletedBook.Price, turkishLira, deletedBook.StockID, deletedBook.ISBN, deletedBook.Author.Name, deletedBook.Author.ID)

		fmt.Println("\nSilinen Kitap:\n", book)
		fmt.Println("Yeni Liste:")
		for _, v := range bookList {
			book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", v.BookID, v.Name, v.Page, v.StockCount, v.Price, turkishLira, v.StockID, v.ISBN, v.Author.Name, v.Author.ID)
			fmt.Println(book)
		}
	}

}

func (b Book) BuyById(id int, count int) {
	var updatedBook *Book
	for i, v := range bookList {
		if v.BookID == id {
			result := v.StockCount - count
			if result >= 0 {
				v.StockCount = result
				updatedBook = &v
				bookList[i] = *updatedBook
				break
			}
			fmt.Println("Yeterli sayıda kitap bulunmamaktadır")
		}
	}
	if updatedBook != nil {
		turkishLira := string(rune(8378))
		book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", updatedBook.BookID, updatedBook.Name, updatedBook.Page, updatedBook.StockCount, updatedBook.Price, turkishLira, updatedBook.StockID, updatedBook.ISBN, updatedBook.Author.Name, updatedBook.Author.ID)
		fmt.Println(book)
	}

}
