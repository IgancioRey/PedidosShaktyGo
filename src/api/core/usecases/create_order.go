package usecases

import (
	"context"
	"log"
	"time"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/repository/database"
)

type CreateOrder interface {
	Execute(order entities.CreateOrderRequest) (entities.CreateOrderResponse, error)
}

type CreateOrderImp struct{}

func (uc *CreateOrderImp) Execute(order entities.CreateOrderRequest) (entities.CreateOrderResponse, error) {
	newOrder, err := createOrder(order)
	return newOrder, err
}

func createOrder(order entities.CreateOrderRequest) (entities.CreateOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `INSERT INTO Orders (CustomerID, Date)
											VALUES (?,?)`,
		order.CustomerID,
		order.Date)
	if err != nil {
		log.Println(err)
		return entities.CreateOrderResponse{}, err
	}
	newOrder, err := getLastOrder()
	if err != nil {
		log.Println(err)
		return entities.CreateOrderResponse{}, err
	}
	customer, err := getCustomerFromOrder(newOrder.CustomerID)
	if err != nil {
		log.Println(err)
		return entities.CreateOrderResponse{}, err
	}
	newOrderResponse := entities.NewOrderResult(newOrder, customer)
	return newOrderResponse, err
}

func getLastOrder() (entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT	Id, CustomerID, DATE_FORMAT(Date, '%d-%m-%Y') as Date
		FROM Orders
		ORDER BY Id DESC`)

	var order entities.Order
	err := row.Scan(&order.Id,
		&order.CustomerID,
		&order.Date)
	if err != nil {
		log.Println(err)
	}

	return order, err
}

func getCustomerFromOrder(id int) (entities.Customer, error) {
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
