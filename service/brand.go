package service

import (
	"context"
	"fmt"
	"log"

	"github.com/mohashari/catalyst-test/model"
)

//BrandRequest ...
type BrandRequest struct {
	Name string `json:"name" validate:"required"`
}

//Valid ...
func (b *BrandRequest) Valid() error {
	if b.Name == "" {
		return fmt.Errorf("name required")
	}
	return nil
}

func (s *service) CreateBrand(ctx context.Context, req BrandRequest) (resp DefaultResponse, err error) {

	if err := req.Valid(); err != nil {
		log.Println("level ", "err ", "method ", "validate ", "message ", err.Error())
		return DefaultResponse{}, err
	}

	id, err := s.repo.BrandRepo.Insert(ctx, model.Brand{
		Name: req.Name,
	})

	if err != nil {
		return resp, err
	}

	return DefaultResponse{
		Message: "success",
		Data:    id,
	}, nil
}
