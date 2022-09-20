package model

//Product ...
type Product struct {
	ID       int64   `json:"id"`
	Brand    Brand   `json:"brand"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
