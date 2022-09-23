package repository

import (
	"context"
	"io"

	"github.com/mohashari/catalyst-test/model"
)

//Repository ...
type Repository struct {
	CustomerRepo    CustomerRepo
	OrderRepo       OrderRepo
	BrandRepo       BrandRepo
	ProductRepo     ProductRepo
	OrderDetailRepo OrderDetailRepo
	io.Closer
}

//BrandRepo ...
type BrandRepo interface {
	Insert(ctx context.Context, model model.Brand) (id int64, err error)
	GetByID(ctx context.Context, id int64) (model model.Brand, err error)
}

//CustomerRepo ...
type CustomerRepo interface {
	GetByID(ctx context.Context, id int64) (customer model.Customer, err error)
}

//OrderDetailRepo ...
type OrderDetailRepo interface {
	Insert(ctx context.Context, orderDetail model.OrderDetail) (err error)
	GetDetailByOrderID(ctx context.Context, id int64) (orderDetail model.OrderDetail, err error)
}

//OrderRepo ...
type OrderRepo interface {
	Insert(ctx context.Context, order model.Order) (id int64, err error)
	GetByID(ctx context.Context, id int64) (order model.Order, err error)
}

//ProductRepo ...
type ProductRepo interface {
	Insert(ctx context.Context, model model.Product) (id int64, err error)
	GetByID(ctx context.Context, id int64) (product model.Product, err error)
	GetByBrandID(ctx context.Context, id int64) (products []model.Product, err error)
}
