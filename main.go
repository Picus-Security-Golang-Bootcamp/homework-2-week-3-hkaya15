package main

import (
	"fmt"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
)

func main() {
	bookCreation := book.Book{Name: "HK"}
	x := bookCreation.GetList()
	fmt.Println(x)

}
