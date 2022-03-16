package main

import (
	"fmt"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
)

func main() {
	//ui.GetUI()
	book1 := book.Book{}
	fmt.Println(book1.SearchByName("a"))
}
