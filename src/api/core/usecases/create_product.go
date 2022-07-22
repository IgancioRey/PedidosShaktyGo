package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type CreateProduct interface {
	Execute(prod entities.CreateProductRequest) error
}

type CreateProductImp struct{}

func (uc *CreateProductImp) Execute(prod entities.CreateProductRequest) error {
	err := createProduct(prod)
	if err != nil {
		log.Println(err.Error())
	}
	return err

}

func createProduct(prod entities.CreateProductRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `INSERT INTO Products (ProductName, Stock, Price, SellPrice, Date)
											VALUES (?,?,?,?,?)`,
		prod.ProductName,
		prod.Stock,
		prod.Price,
		prod.SellPrice,
		prod.Date)
	return err
}
