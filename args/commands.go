package args

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
)

var (
	list      = flag.NewFlagSet("list", flag.ExitOnError)
	search    = flag.NewFlagSet("search", flag.ExitOnError)
	get       = flag.NewFlagSet("get", flag.ExitOnError)
	id        = get.Int("id", 0, "Book Id")
	delete    = flag.NewFlagSet("delete", flag.ExitOnError)
	deletedId = delete.Int("id", 0, "Delete Id")
	buy       = flag.NewFlagSet("buy", flag.ExitOnError)
	bookId    = buy.Int("bookID", 0, "Book Id")
	quantity  = buy.Int("quantity", 0, "quantity")
)

var usage = `Usage: ./main.go <arg> <arg>
Args:
	-list The list of the all book.
	-search  Search by book name. It works as that whether contain as a word.
	-get	Get by book Id.
	-delete Delete by book Id.
	-buy Buy by book Id.
`

func ManageCommands() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", usage)

	}

	if len(os.Args) < 2 {
		sendError()
	}
	switch os.Args[1] {
	case "list":
		flag.Parse()
		if len(os.Args) == 2 {
			res, err := NewFuncCommand(list).GetList(book.Book{})
			if err != nil {
				fmt.Println(err)
			}
			turkishLira := string(rune(8378))
			for _, v := range res {
				book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", v.BookID, v.Name, v.Page, v.StockCount, v.Price, turkishLira, v.StockID, v.ISBN, v.Author.Name, v.Author.ID)
				fmt.Println(book)
			}

		} else {
			sendError()
		}

	case "search":
		search.Parse(os.Args[2:])
		if len(os.Args) > 2 {
			val := strings.Join(search.Args(), " ")
			res, err := NewFuncCommand(search).SearchByName(book.Book{}, val)
			if err != nil {
				fmt.Println(err)
			}
			if len(res) > 0 {
				turkishLira := string(rune(8378))
				for _, v := range res {
					book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", v.BookID, v.Name, v.Page, v.StockCount, v.Price, turkishLira, v.StockID, v.ISBN, v.Author.Name, v.Author.ID)
					fmt.Println(book)
				}
			} else {
				fmt.Println("İlgili kitap bulunamadı.")
			}

		} else {
			sendError()
		}

	case "get":
		get.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Please enter a valid Id, examples: -id")
		} else {
			v, err := NewFuncCommand(get).SearchById(book.Book{}, *id)
			if err != nil {
				fmt.Println(err)
			}
			if v.BookID == 0 {
				fmt.Println("Bu Id'de bir kitap bulunmamaktadır")
			} else {
				turkishLira := string(rune(8378))
				book := fmt.Sprintf("\n\tKitap ID'si: %v\n \tKitap Adı: %s\n \tSayfa Sayısı: %v\n \tStock Sayısı %v\n \tFiyat: %.2f %v\n \tStok Kodu: %v\n \tISBN : %v\n \tYazar Adı : %s\n \tYazar ID'si : %v\n", v.BookID, v.Name, v.Page, v.StockCount, v.Price, turkishLira, v.StockID, v.ISBN, v.Author.Name, v.Author.ID)
				fmt.Println(book)
			}

		}

	case "delete":
		delete.Parse(os.Args[2:])
		if *deletedId == 0 {
			fmt.Println("Please enter a valid Id: examples: -id")
		} else {
			err := NewFuncCommand(delete).DeleteById(book.Book{}, *deletedId)
			if err != nil {
				fmt.Println(err)
			}
		}

	case "buy":
		buy.Parse(os.Args[2:])
		if *bookId == 0 {
			fmt.Println("Please enter a valid args: examples: -bookID -quantity")
		} else {
			if *quantity <= 0 {
				fmt.Println("Lütfen Geçerli bir adet değeri giriniz")
			} else {
				err := NewFuncCommand(buy).BuyById(book.Book{}, *bookId, *quantity)
				if err != nil {
					fmt.Println(err)
				}
			}

		}

	default:
		sendError()

	}
}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, "%s", msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func sendError() {
	err := fmt.Errorf("\nUsage Fault!!!")
	usageAndExit(err.Error())
}
