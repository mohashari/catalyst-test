package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mohashari/catalyst-test/model"
)

func Test_product_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx   context.Context
		model model.Product
	}
	prd := model.Product{}
	db, dbmock, _ := sqlmock.New()
	query := `insert into product(brand_id,name,price,quantity) values($1,$2,$3,$4) returning id`
	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			prd.Brand.ID,
			prd.Name,
			prd.Price,
			prd.Quantity,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			prd.Brand.ID,
			prd.Name,
			prd.Price,
			prd.Quantity,
		).
		WillReturnError(sql.ErrConnDone)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "Test_product_Insert positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:   context.Background(),
				model: model.Product{},
			},
			wantId:  1,
			wantErr: false,
		},
		{
			name: "Test_product_Insert connection done",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:   context.Background(),
				model: model.Product{},
			},
			wantId:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &product{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotId, err := p.Insert(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("product.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("product.Insert() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
