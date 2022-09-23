package model

import "time"

//Order ...
type Order struct {
	ID           int64         `json:"id,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	OrderDate    time.Time     `json:"order_date,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	Amount       float64       `json:"amount,omitempty"`
	OrderDetails []OrderDetail `json:"order_details,omitempty"`
}
