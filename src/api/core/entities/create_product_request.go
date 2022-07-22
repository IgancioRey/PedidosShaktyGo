package entities

import "time"

type CreateProductRequest struct {
	ProductName string
	Stock       int
	Price       float32
	SellPrice   float32
	Date        string
}

func (cpr *CreateProductRequest) IsValid() bool {
	date := validDate(cpr.Date)
	price := validPrices(cpr.Price, cpr.SellPrice)
	return date && price
}

func validDate(date string) bool {
	_, err := time.Parse("20060102", date)

	return err == nil
}

func validPrices(price float32, sellPrice float32) bool {
	return (0 < price) && (price < sellPrice)
}
