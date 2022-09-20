package service

import (
	"context"

	"github.com/mohashari/catalyst-test/repository"
)

const (
	success = "success"
)

type service struct {
	repo *repository.Repository
}

//Service ...
type Service interface {
	CreateBrand(ctx context.Context, req BrandRequest) (resp DefaultResponse, err error)
	CreateProduct(ctx context.Context, req ProductCreateReq) (resp DefaultResponse, err error)
	GetProductByID(ctx context.Context, id int64) (resp DefaultResponse, err error)
	GetProductByBrandID(ctx context.Context, id int64) (resp DefaultResponse, err error)
}

//NewService ...
func NewService(repo *repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

//DefaultResponse ...
type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
