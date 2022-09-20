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
}

//OrderDetailRepo ...
type OrderDetailRepo interface {
}

//OrderRepo ...
type OrderRepo interface {
}

//ProductRepo ...
type ProductRepo interface {
	Insert(ctx context.Context, model model.Product) (id int64, err error)
	GetByID(ctx context.Context, id int64) (product model.Product, err error)
	GetByBrandID(ctx context.Context, id int64) (products []model.Product, err error)
}
