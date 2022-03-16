package product

type IProduct interface {
	Create() IProduct
	GetList() []IProduct
}

