package repository

import (
	"github.com/jmoiron/sqlx"
	"server/entity"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) MakeOrder(userId int, products []entity.Product) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	var cartId int
	query := "SELECT cart_id FROM users WHERE id = $1"
	row := tx.QueryRow(query, userId)
	err = row.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	var orderId int
	query = "INSERT INTO order_from_cart (cart_id) VALUES ($1) RETURNING id"
	row = tx.QueryRow(query, cartId)
	err = row.Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	query = "INSERT INTO order_product (order_id, product_id) VALUES ($1, $2)"
	for i := 0; i < len(products); i++ {
		_, err = tx.Exec(query, orderId, products[i].Id)
		if err != nil {
			tx.Rollback()
			return -1, nil
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return orderId, nil
}

func (r *OrderRepository) GetProductFromCart(userId int) ([]entity.Product, error) {
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
