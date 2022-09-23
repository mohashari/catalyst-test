package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/mohashari/catalyst-test/model"
)

type order struct {
	db *sql.DB
	mu sync.RWMutex
}

func (o *order) Insert(ctx context.Context, order model.Order) (id int64, err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	query := `insert into order (customer_id,order_date,created_at,amount) values($1,$2,$3,$4) returning id`
	queryOrderDetail := `insert into order_detail (order_id,product_id,amount,quantity) values ($1,$2,$3,$4)`

	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return id, err
	}

	for _, orderDetail := range order.OrderDetails {
		_, err = tx.ExecContext(
			ctx,
			queryOrderDetail,
			orderDetail.OrderID,
			orderDetail.Product.ID,
			orderDetail.Amount,
			orderDetail.Quantity,
		)
		if err != nil {
			tx.Rollback()
			return id, err
		}
	}

	if err := tx.QueryRowContext(
		ctx,
		query,
		order.Customer.ID,
		order.OrderDate,
		order.CreatedAt,
		order.Amount).Scan(&id); err != nil {
		tx.Rollback()
		return id, err
	}
	tx.Commit()
	return id, nil
}

func (o *order) GetByID(ctx context.Context, id int64) (order model.Order, err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	query := `select id,customer_id,order_date,created_at,amount from order where id = $1`
	if err := o.db.QueryRowContext(ctx, query, id).Scan(
		&order.ID,
		&order.Customer.ID,
		&order.OrderDate,
		&order.CreatedAt,
		&order.Amount,
	); err != nil {
		return order, err
	}
	return order, nil
}
