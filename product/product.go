package product

type IProduct interface {
	Create() IProduct
	GetList() []IProduct
	SearchById(id int) IProduct
}

