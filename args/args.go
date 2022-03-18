package args

import (
	"flag"

	"github.com/hkaya15/PicusSecurity/Week_3_Homework/book"
)

type Runner interface {
	Init([]string) error
	GetList(obj interface{}) (interface{}, error)
	SearchByName(obj interface{}, name string) (interface{}, error)
	SearchById() (interface{}, error)
	DeleteById(obj interface{}, id int) error
	BuyById(obj interface{}, id int, count int)
}

func NewFuncCommand(arg *flag.FlagSet) *NewCommand {
	command := &NewCommand{
		arg: arg,
	}
	return command
}

type NewCommand struct {
	arg *flag.FlagSet
}

func (l *NewCommand) GetList(arg book.Book) ([]book.Book, error) {
	res := arg.GetList()
	return res, nil
}

func (l *NewCommand) SearchByName(arg book.Book, name string) ([]book.Book, error) {
	res := arg.SearchByName(name)
	return res, nil
}
func (l *NewCommand) SearchById(arg book.Book, id int) (book.Book, error) {
	res := arg.SearchById(id)
	return res, nil
}
func (l *NewCommand) DeleteById(arg book.Book, id int) error {
	arg.DeleteById(id)
	return nil
}

func (l *NewCommand) BuyById(arg book.Book, id int, count int) error {
	arg.BuyById(id, count)
	return nil
}
