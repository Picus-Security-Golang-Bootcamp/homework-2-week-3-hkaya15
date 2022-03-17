package product

type IProduct interface {
	Create() IProduct
	GetList() []IProduct
	AddList()
	SearchById(id int) IProduct
	SearchByName(name string) []IProduct
	DeleteById(id int)
	BuyById(id int, count int) 
}
