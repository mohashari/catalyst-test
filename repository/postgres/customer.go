package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/mohashari/catalyst-test/model"
)

type customer struct {
	db *sql.DB
	mu sync.RWMutex
}

func (c *customer) GetByID(ctx context.Context, id int64) (customer model.Customer, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	query := `select id,name,email from customer where id = $1`

	if err := c.db.QueryRowContext(ctx, query, id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
	); err != nil {
		return customer, err
	}
	return customer, nil

}
