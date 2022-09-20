package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/mohashari/catalyst-test/model"
)

type product struct {
	db *sql.DB
	mu sync.RWMutex
}

func (p *product) Insert(ctx context.Context, model model.Product) (id int64, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := `insert into product(brand_id,name,price,quantity) values($1,$2,$3,$4) returning id`
	if err := p.db.QueryRowContext(
		ctx,
		query,
		model.Brand.ID,
		model.Name,
		model.Price,
		model.Quantity,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (p *product) GetByID(ctx context.Context, id int64) (product model.Product, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := `SELECT id, brand_id,name,price,quantity FROM product where id = $1`

	if err := p.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Brand.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
	); err != nil {
		return product, err
	}
	return product, nil
}
