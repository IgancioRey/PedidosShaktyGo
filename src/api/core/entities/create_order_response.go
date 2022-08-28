package entities

type CreateOrderResponse struct {
	Id       int      `json:"id"`
	Customer Customer `json:"customer"`
	Date     string   `json:"date"`
}

func NewOrderResult(order Order, customer Customer) CreateOrderResponse {
	return CreateOrderResponse{
		Id:       order.Id,
		Customer: customer,
		Date:     order.Date,
	}
}
