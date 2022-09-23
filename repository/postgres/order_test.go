package postgres

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mohashari/catalyst-test/model"
)

func Test_order_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx   context.Context
		order model.Order
	}

	now := time.Now()

	orderDetailModel := model.OrderDetail{

		OrderID: 1,
		Product: model.Product{
			ID: 1,
		},
		Amount:   100000,
		Quantity: 10,
	}
	orderModel := model.Order{
		ID:           1,
		Customer:     model.Customer{ID: 1},
		OrderDate:    now,
		CreatedAt:    now,
		OrderDetails: []model.OrderDetail{orderDetailModel},
		Amount:       30000,
	}

	db, dbmock, _ := sqlmock.New()
	queryOrderDetail := `insert into order_detail (order_id,product_id,amount,quantity) values ($1,$2,$3,$4)`
	query := `insert into order (customer_id,order_date,created_at,amount) values($1,$2,$3,$4) returning id`

	dbmock.ExpectBegin()
	dbmock.ExpectExec(regexp.QuoteMeta(queryOrderDetail)).
		WithArgs(orderDetailModel.OrderID,
			orderDetailModel.Product.ID,
			orderDetailModel.Amount,
			orderDetailModel.Quantity).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			orderModel.Customer.ID,
			orderModel.OrderDate,
			orderModel.CreatedAt,
			orderModel.Amount).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(int64(1)))
	dbmock.ExpectCommit()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "Test_order_Insert positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:   context.Background(),
				order: orderModel,
			},
			wantId:  1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &order{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotId, err := o.Insert(tt.args.ctx, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("order.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("order.Insert() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_order_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	query := `select id,customer_id,order_date,created_at,amount from order where id = $1`

	orderModel := model.Order{}
	db, dbmock, _ := sqlmock.New()
	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "order_date", "created_at", "amount"}).
			AddRow(
				orderModel.ID,
				orderModel.Customer.ID,
				orderModel.OrderDate,
				orderModel.CreatedAt,
				orderModel.Amount,
			))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).WillReturnError(sqlmock.ErrCancelled)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOrder model.Order
		wantErr   bool
	}{
		{
			name: "Test_order_GetByID positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantOrder: orderModel,
			wantErr:   false,
		},

		{
			name: "Test_order_GetByID err sql",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantOrder: model.Order{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &order{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotOrder, err := o.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("order.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("order.GetByID() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
