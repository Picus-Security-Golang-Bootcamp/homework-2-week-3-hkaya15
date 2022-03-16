package main

import (
	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
	"github.com/hkaya15/PicusSecurity/Week_3_Homework/product"
)

func main() {
	book := book.Book{}
	product := product.ProductManager{Pr: book}
	product.Create()

}
