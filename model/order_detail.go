package model

//OrderDetail ...
type OrderDetail struct {
	ID       int64   `json:"id,omitempty"`
	OrderID  int64   `json:"order_id,omitempty"`
	Product  Product `json:"product,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
