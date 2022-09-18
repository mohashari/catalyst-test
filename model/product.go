package model

//Product ...
type Product struct {
	ID      int64
	Brand   Brand
	Name    string
	Price   float64
	Quality int
}
