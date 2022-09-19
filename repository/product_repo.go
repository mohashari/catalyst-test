package repository

import (
	"context"

	"github.com/mohashari/catalyst-test/model"
)

//ProductRepo ...
type ProductRepo interface {
	Insert(ctx context.Context, model model.Product) (id int64, err error)
}
