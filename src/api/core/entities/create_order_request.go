package entities

import "time"

type CreateOrderRequest struct {
	CustomerID int32  `json:"customer_id"`
	Date       string `json:"date"`
}

func (ccr *CreateOrderRequest) IsValid() bool {
	date := validDateRequest(ccr.Date)
	return date
}

func validDateRequest(date string) bool {
	_, err := time.Parse("20060102", date)
	return err == nil
}
