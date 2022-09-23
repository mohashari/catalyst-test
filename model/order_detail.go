package model

//OrderDetail ...
type OrderDetail struct {
	ID       int64
	OrderID  int64
	Product  Product
	Amount   float64
	Quantity int
}
