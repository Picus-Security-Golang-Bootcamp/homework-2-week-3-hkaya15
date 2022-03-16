package ui

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
)

const (
	createBook = "1"
	quit       = "q"
)

func GetUI() {

	for {
		fmt.Println("\n*** Book Systems ***")
		fmt.Print("\nPress '1' to create a new Book\nPress 'q' to exit\n")
		scanner, err := getInput()
		checkError(err)
		result := scanner.Text()
		if result == createBook {
			newBook := book.Book{}
			fmt.Print("Kitap Adı Giriniz: ")
			scanner, _ := getInput()
			newBook.Name = scanner.Text()

			fmt.Print("Sayfa sayısı Giriniz: ")
			scanner, _ = getInput()
			value := scanner.Text()
			newBook.Page,_ = strconv.Atoi(value)

			fmt.Print("Stok sayısı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.StockCount,_ = strconv.Atoi(value)

			fmt.Print("Fiyatı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Price,_ = strconv.ParseFloat(value,64)

			fmt.Print("ISBN Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.ISBN,_ = strconv.Atoi(value)

			fmt.Print("Yazar Adı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Author.Name = value

			fmt.Print("Yazar ID'si Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Author.ID,_ = strconv.Atoi(value)

			newBook.BookID=rand.Int()
			newBook.StockID=rand.Int()

			fmt.Println("\nKitabınız :",newBook)

		} else if strings.ToLower(result) == quit {
			break
		} else {
			fmt.Println("Lütfen girdiğiniz değeri kontrol edin!")
			continue
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getInput() (*bufio.Scanner, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	return scanner, err
}
