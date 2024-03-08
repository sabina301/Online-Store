package repository

import "github.com/jmoiron/sqlx"

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) AddProduct(category string, name string, color string, description string, price uint32) (int, error) {
	var id int
	query := "INSERT INTO products(category, name, price, description, color) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	row := pr.db.QueryRow(query, category, name, price, description, color)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
