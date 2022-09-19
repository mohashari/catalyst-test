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

func Test_brand_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx   context.Context
		model model.Brand
	}

	m := model.Brand{
		Name: "cata",
	}
	db, dbMock, _ := sqlmock.New()
	query := `INSERT INTO brand (name) VALUES ($1) returning id`
	dbMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(m.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	dbMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(m.Name).
		WillReturnError(sql.ErrConnDone)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "Test_brand_Insert positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				model: model.Brand{
					Name: "cata",
				},
			},
			wantId:  1,
			wantErr: false,
		},

		{
			name: "Test_brand_Insert error connection done",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				model: model.Brand{
					Name: "cata",
				},
			},
			wantId:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &brand{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotId, err := b.Insert(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("brand.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("brand.Insert() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_brand_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
		mu sync.RWMutex
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	db, dbmock, _ := sqlmock.New()
	query := ` select id,name from brand where id = $1`

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(int64(1), "cata"))

	dbmock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnError(sql.ErrConnDone)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantModel model.Brand
		wantErr   bool
	}{
		{
			name: "Test_brand_GetByID positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantModel: model.Brand{
				ID:   1,
				Name: "cata",
			},
			wantErr: false,
		},

		{
			name: "Test_brand_GetByID positif case",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantModel: model.Brand{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &brand{
				db: tt.fields.db,
				mu: tt.fields.mu,
			}
			gotModel, err := b.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("brand.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("brand.GetByID() = %v, want %v", gotModel, tt.wantModel)
			}
		})
	}
}
