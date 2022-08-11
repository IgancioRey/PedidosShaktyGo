package entities

type Product struct {
	Id          int
	ProductName string
	Stock       int
	Price       float32
	SellPrice   float32
	Date        string
}

/*
func NewProduct(id int, name string, stock int, price float32, sellPrice float32, date time.Time) Product {
	return Product{
		Id:          id,
		ProductName: name,
		Stock:       stock,
		Price:       price,
		SellPrice:   sellPrice,
		Date:        date,
	}
}
*/
