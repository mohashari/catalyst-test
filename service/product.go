package service

import (
	"context"
	"fmt"
	"log"

	"github.com/mohashari/catalyst-test/model"
)

//ProductCreateReq ...
type ProductCreateReq struct {
	BrandID  int64   `json:"brand_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

//Valid ...
func (p *ProductCreateReq) Valid() error {
	if p.BrandID <= 0 {
		return fmt.Errorf("brand id required")
	}
	if p.Name == "" {
		return fmt.Errorf("name required")
	}
	if p.Price <= 0 {
		return fmt.Errorf("price required")
	}
	if p.Quantity <= 0 {
		return fmt.Errorf("quantity reuired")
	}
	return nil
}

func (s *service) CreateProduct(ctx context.Context, req ProductCreateReq) (resp DefaultResponse, err error) {
	if err := req.Valid(); err != nil {
		log.Println("level: ", "err ", "method: ", "Valid req product ", "message: ", err.Error())
		return resp, err
	}

	brand, err := s.repo.BrandRepo.GetByID(ctx, req.BrandID)
	if err != nil {
		log.Println("level: ", "err ", "method: ", "Get Brand ByID ", "message: ", err.Error())
		return resp, err
	}

	id, err := s.repo.ProductRepo.Insert(ctx, model.Product{
		Brand:    brand,
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	})

	if err != nil {
		log.Println("level: ", "err ", "method: ", "Insert Product ", "message: ", err.Error())
		return resp, err
	}

	return DefaultResponse{
		Message: success,
		Data:    id,
	}, nil
}
