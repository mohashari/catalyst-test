package repository

import (
	"context"

	"github.com/mohashari/catalyst-test/model"
)

//BrandRepo ...
type BrandRepo interface {
	Insert(ctx context.Context, model model.Brand) (id int64, err error)
	GetByID(ctx context.Context, id int64) (model model.Brand, err error)
}
