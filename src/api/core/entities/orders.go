package entities

type Order struct {
	Id         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	Date       string `json:"date"`
}
