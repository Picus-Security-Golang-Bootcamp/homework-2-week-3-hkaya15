package book

import (
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

func (b Book) SearchById(id int) Book{
	return b
}