package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/mohashari/catalyst-test/repository"
)

//ConnParam ...
type ConnParam struct {
	Host        string
	Port        string
	DBName      string
	User        string
	Pass        string
	Options     string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime time.Duration
}

type postgres struct {
	db *sql.DB
}

//NewPostgres ...
func NewPostgres(p *ConnParam) (*repository.Repository, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s", p.User, p.Pass, p.Host, p.Port, p.DBName, p.Options)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("db open: %v", err)
	}
	db.SetMaxOpenConns(p.MaxOpenConn)
	db.SetMaxIdleConns(p.MaxIdleConn)
	db.SetConnMaxLifetime(p.MaxLifetime)
	return &repository.Repository{
		Closer: &postgres{
			db: db,
		},
		BrandRepo: &brand{
			db: db,
		},
		ProductRepo: &product{
			db: db,
		},
		CustomerRepo: &customer{
			db: db,
		},
		OrderRepo: &order{
			db: db,
		},
		OrderDetailRepo: &orderDetail{
			db: db,
		},
	}, nil
}

// Close ...
func (p *postgres) Close() error {
	if p.db != nil {
		if err := p.db.Close(); err != nil {
			return err
		}
		p.db = nil
	}
	return nil
}
