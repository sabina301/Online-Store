package repository

import "github.com/jmoiron/sqlx"

type ProductRepository struct {
}

func NewProductRepository(database *sqlx.DB) *ProductRepository {
	return &ProductRepository{}
}
