package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"tdd/config"
	"time"
)

type MySQL struct {
	DB *sqlx.DB
}

func NewMySQL(conf config.MySQLTemplate) (*MySQL, error) {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true",
		conf.Username, conf.Password, conf.Address, conf.DBName,
	)

	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(time.Minute * 10)

	return &MySQL{DB: db}, nil
}

func (p *MySQL) Close() (err error) {
	err = p.DB.Close()
	if err != nil {
		return
	}

	return
}
