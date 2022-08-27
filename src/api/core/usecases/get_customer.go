package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type GetCustomer interface {
	Execute(id int) (entities.Customer, error)
}

type GetCustomerImp struct{}

func (uc *GetCustomerImp) Execute(id int) (entities.Customer, error) {
	customer, err := getCustomer(id)

	return customer, err
}

func getCustomer(id int) (entities.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT * FROM Customers WHERE Id = ?`, id)

	var customer entities.Customer
	err := row.Scan(&customer.Id,
		&customer.Name,
		&customer.LastName,
		&customer.CellNumber,
		&customer.Observation,
	)
	if err != nil {
		log.Println(err)
	}

	return customer, err
}
