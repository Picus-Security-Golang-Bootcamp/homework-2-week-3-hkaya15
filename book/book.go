package book

import (
	"fmt"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/author"
)

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

func (b Book) Create(){
	fmt.Println("Book Created")
}
