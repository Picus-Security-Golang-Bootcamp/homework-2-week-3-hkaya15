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
	getList    = "2"
)

func GetUI() {
	newBook := book.Book{}
	for {
		fmt.Println("\n*** Book Systems ***")
		fmt.Print("\nPress '1' to create a new Book\nPress '2' to see all list\nPress 'q' to exit\n")
		scanner, err := getInput()
		checkError(err)
		result := scanner.Text()
		if result == createBook {

			fmt.Print("Kitap Adı Giriniz: ")
			scanner, _ := getInput()
			newBook.Name = scanner.Text()

			fmt.Print("Sayfa sayısı Giriniz: ")
			scanner, _ = getInput()
			value := scanner.Text()
			newBook.Page, _ = strconv.Atoi(value)

			fmt.Print("Stok sayısı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.StockCount, _ = strconv.Atoi(value)

			fmt.Print("Fiyatı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Price, _ = strconv.ParseFloat(value, 64)

			fmt.Print("ISBN Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.ISBN, _ = strconv.Atoi(value)

			fmt.Print("Yazar Adı Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Author.Name = value

			fmt.Print("Yazar ID'si Giriniz: ")
			scanner, _ = getInput()
			value = scanner.Text()
			newBook.Author.ID, _ = strconv.Atoi(value)

			newBook.BookID = rand.Int()
			newBook.StockID = rand.Int()

			newBook.AddList()
			fmt.Println("\nKitabınız Başarıyla Eklendi:\n", newBook)

		} else if result == getList {
			list := newBook.GetList()
			readList(list)
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

func readList(booklist []book.Book) {
	turkishLira := string(rune(8378))
	for _, v := range booklist {
		book := fmt.Sprintf("\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", v.BookID, v.Name, v.Page, v.StockCount, v.Price, turkishLira, v.StockID, v.ISBN, v.Author.Name, v.Author.ID)
		fmt.Println(book)
	}
}
