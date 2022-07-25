package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type GetCustomers interface {
	Execute() ([]entities.Customer, error)
}

type GetCustomersImp struct{}

func (uc *GetCustomersImp) Execute() ([]entities.Customer, error) {
	customersList := getCustomersList()

	return customersList, nil
}

func getCustomersList() []entities.Customer {
	customersList := make([]entities.Customer, 0)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := database.DbConn.Query(`SELECT * FROM Customers`)

	if err != nil {
		log.Println(err.Error())
		return customersList
	}

	defer results.Close()

	var customer entities.Customer
	for results.Next() {
		results.Scan(
			&customer.Id,
			&customer.Name,
			&customer.LastName,
			&customer.CellNumber,
			&customer.Observation,
		)
		customersList = append(customersList, customer)
	}
	return customersList
}
