package product

import "github.com/hkaya15/PicusSecurity/Week_3_Homework/book"

type IProduct interface {
	Create() book.Book
	GetList() []book.Book
	AddList()
	SearchById(id int) book.Book
	SearchByName(name string) []book.Book
	DeleteById(id int)
	BuyById(id int, count int) 
}
