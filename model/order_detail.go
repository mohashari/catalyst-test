package model

//OrderDetail ...
type OrderDetail struct {
	ID       int64
	Order    Order
	Product  Product
	Amount   float64
	Quantity int
}
