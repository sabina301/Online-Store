package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	connStr   = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	userTable = "users"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Database struct {
	db *sqlx.DB
}

func NewDatabase(conf DatabaseConfig) (*Database, error) {
	db, err := sqlx.Open("postgres", "host=ps-psql port=5432 user=postgres password=Ertyu55555 dbname=online-store-db sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}

func (d *Database) CloseDB() {
	err := d.db.Close()
	if err != nil {
		return
	}
}
