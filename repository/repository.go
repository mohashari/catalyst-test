package repository

import "io"

//Repository ...
type Repository struct {
	CustomerRepo    CustomerRepo
	OrderRepo       OrderRepo
	BrandRepo       BrandRepo
	ProductRepo     ProductRepo
	OrderDetailRepo OrderDetailRepo
	io.Closer
}
