package postgres

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mohashari/catalyst-test/model"
)

func Test_orderDetail_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx         context.Context
		orderDetail model.OrderDetail
	}

	orderDetailModel := model.OrderDetail{
		OrderID: 1,
		Product: model.Product{
			ID: 1,
		},
		Amount:   100000,
		Quantity: 10,
	}

	db, dbmock, _ := sqlmock.New()
	query := `insert into order_detail (order_id,product_id,amount,quantity) values ($1,$2,$3,$4)`

	dbmock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(
			orderDetailModel.OrderID,
			orderDetailModel.Product.ID,
			orderDetailModel.Amount,
			orderDetailModel.Quantity,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	dbmock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(
			orderDetailModel.OrderID,
			orderDetailModel.Product.ID,
			orderDetailModel.Amount,
			orderDetailModel.Quantity,
		).WillReturnError(sqlmock.ErrCancelled)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test_orderDetail_GetDetailByID",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:         context.Background(),
				orderDetail: orderDetailModel,
			},
			wantErr: false,
		},

		{
			name: "Test_orderDetail_GetDetailByID sql error",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:         context.Background(),
				orderDetail: orderDetailModel,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDetail{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			if err := o.Insert(tt.args.ctx, tt.args.orderDetail); (err != nil) != tt.wantErr {
				t.Errorf("orderDetail.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_orderDetail_GetDetailByOrderID(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	var orderDetailModel model.OrderDetail
	db, dbmock, _ := sqlmock.New()
	query := `select id,order_id,product_id,amount,quantity from order_detail where order_id = $1`

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "product_id", "amount", "quantity"}).
			AddRow(
				orderDetailModel.ID,
				orderDetailModel.OrderID,
				orderDetailModel.Product.ID,
				orderDetailModel.Amount,
				orderDetailModel.Quantity))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).
		WillReturnError(sqlmock.ErrCancelled)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOrder model.OrderDetail
		wantErr   bool
	}{
		{
			name: "Test_orderDetail_GetDetailByOrderID",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantOrder: orderDetailModel,
			wantErr:   false,
		},

		{
			name: "Test_orderDetail_GetDetailByOrderID sql err",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantOrder: model.OrderDetail{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDetail{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotOrder, err := o.GetDetailByOrderID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderDetail.GetDetailByOrderID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("orderDetail.GetDetailByOrderID() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
