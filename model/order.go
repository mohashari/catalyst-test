package model

import "time"

//Order ...
type Order struct {
	ID        int64
	Customer  Customer
	OrderDate time.Time
	Amount    float64
}
