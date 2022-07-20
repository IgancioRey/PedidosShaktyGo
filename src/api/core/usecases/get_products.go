package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type GetProducts interface {
	Execute() ([]entities.Product, error)
}

type GetProductsImp struct{}

func (uc *GetProductsImp) Execute() ([]entities.Product, error) {
	productsList := getProductsList()

	return productsList, nil
}

func getProductsList() []entities.Product {
	productsList := make([]entities.Product, 0)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := database.DbConn.Query(`SELECT *
	FROM Products`)

	if err != nil {
		log.Println(err.Error())
		return productsList
	}

	defer results.Close()

	var product entities.Product
	for results.Next() {
		results.Scan(
			&product.Id,
			&product.ProductName,
			&product.Stock,
			&product.Price,
			&product.SellPrice,
			&product.Date,
		)
		productsList = append(productsList, product)
	}
	return productsList
}
