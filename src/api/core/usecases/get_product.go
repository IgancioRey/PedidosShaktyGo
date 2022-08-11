package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type GetProduct interface {
	Execute(id int) (entities.Product, error)
}

type GetProductImp struct{}

func (uc *GetProductImp) Execute(id int) (entities.Product, error) {

	product, err := getProduct(id)

	return product, err
}

func getProduct(id int) (entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT	Id, ProductName, Stock, Price, SellPrice, DATE_FORMAT(Date, '%d-%m-%Y') as Date
		FROM Products
		WHERE Id = ?`, id)

	var product entities.Product
	err := row.Scan(&product.Id,
		&product.ProductName,
		&product.Stock,
		&product.Price,
		&product.SellPrice,
		&product.Date,
	)
	if err != nil {
		log.Println(err)
	}

	return product, err
}
