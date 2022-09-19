package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/mohashari/catalyst-test/model"
)

type brand struct {
	db *sql.DB
	mu sync.RWMutex
}

func (b *brand) Insert(ctx context.Context, model model.Brand) (id int64, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	query := `INSERT INTO brand (name) VALUES ($1) returning id`
	if err := b.db.QueryRowContext(ctx, query, model.Name).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (b *brand) GetByID(ctx context.Context, id int64) (model model.Brand, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	query := ` select id,name from brand where id = $1`

	if err := b.db.QueryRowContext(ctx, query, id).Scan(
		&model.ID,
		&model.Name,
	); err != nil {
		return model, err
	}
	return model, nil
}
