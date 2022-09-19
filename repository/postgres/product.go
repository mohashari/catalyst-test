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
		model.Quality,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
