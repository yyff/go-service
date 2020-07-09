package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yyff/go-service/conf"
)

type Dao struct {
	DB *sqlx.DB
	// rdb redis.Client
}

func New(config *conf.Config) (*Dao, error) {
	db, err := sqlx.Open("mysql", config.MySQL.DSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Dao{db}, nil
}
