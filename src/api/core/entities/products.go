package entities

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
	SellPrice   float32 `json:"sell_price"`
	Date        string  `json:"date"`
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
