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

func Test_customer_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	customerModel := model.Customer{}
	db, dbmock, _ := sqlmock.New()

	query := `select id,name,email from customer where id = $1`
	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.
			NewRows([]string{
				"id", "name", "email",
			}).AddRow(
			customerModel.ID,
			customerModel.Name,
			customerModel.Email,
		))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(int64(1)).
		WillReturnError(sql.ErrNoRows)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCustomer model.Customer
		wantErr      bool
	}{
		{
			name: "Test_customer_GetByID",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantCustomer: customerModel,
			wantErr:      false,
		},
		{
			name: "Test_customer_GetByID err no row",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantCustomer: model.Customer{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotCustomer, err := c.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("customer.GetByID() = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}
