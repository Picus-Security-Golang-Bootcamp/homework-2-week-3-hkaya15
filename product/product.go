package product

type IProduct interface {
	Create()
}

type ProductManager struct {
	Pr IProduct
}

// func (pr ProductManager) GetList() *[]IProduct{
// 	return &ProductList;
// }

func (pr ProductManager) Create() {
	pr.Pr.Create()
}

// var ProductList = []IProduct{book.Book{}}
