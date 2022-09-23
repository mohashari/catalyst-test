package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/mohashari/catalyst-test/model"
)

type orderDetail struct {
	db *sql.DB
	mu sync.RWMutex
}

func (o *orderDetail) Insert(ctx context.Context, orderDetail model.OrderDetail) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	query := `insert into order_detail (order_id,product_id,amount,quantity) values ($1,$2,$3,$4)`

	_, err = o.db.ExecContext(ctx, query,
		orderDetail.OrderID,
		orderDetail.Product.ID,
		orderDetail.Amount,
		orderDetail.Quantity,
	)
	if err != nil {
		return err
	}
	return nil
}

func (o *orderDetail) GetDetailByOrderID(ctx context.Context, id int64) (orderDetail []model.OrderDetail, err error) {

	o.mu.Lock()
	defer o.mu.Unlock()

	query := `select id,order_id,product_id,amount,quantity from order_detail where order_id = $1`
	rows, err := o.db.QueryContext(ctx, query, id)
	if err != nil {
		return orderDetail, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.OrderDetail
		if err := rows.Scan(
			&order.ID,
			&order.OrderID,
			&order.Product.ID,
			&order.Amount,
			&order.Quantity,
		); err != nil {
			return orderDetail, err
		}
		orderDetail = append(orderDetail, order)
	}

	return orderDetail, nil
}
