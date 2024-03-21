package repository

import (
	"github.com/jmoiron/sqlx"
	"server/internal/entity"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db}
}

func (r *CartRepository) GetProductFromCart(userId int) ([]entity.Product, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var cartId int
	query := "SELECT cart_id FROM users WHERE id = $1"
	row := tx.QueryRow(query, userId)
	err = row.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	query = "SELECT p.id, p.category, p.name, p.color, p.description, p.price FROM products p JOIN cart_products cp ON p.id = cp.product_id WHERE cp.cart_id = $1"
	rows, err := tx.Query(query, cartId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Category, &product.Name, &product.Color, &product.Description, &product.Price)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		products = append(products, product)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return products, nil
}
