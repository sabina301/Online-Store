package repository

import (
	"github.com/jmoiron/sqlx"
	"server/entity"
)

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

func (pr *ProductRepository) GetAllProducts() ([]entity.Product, error) {
	query := "SELECT * FROM products"
	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Category, &product.Name, &product.Color, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepository) AddProductInCart(userId int, productId int) error {
	tx, err := pr.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	var cartId int
	query1 := "SELECT cart_id FROM users WHERE id = $1"
	row1 := tx.QueryRow(query1, userId)
	err = row1.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query2 := "INSERT INTO cart_products (cart_id, product_id) VALUES ($1, $2)"
	_, err = tx.Exec(query2, cartId, productId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
