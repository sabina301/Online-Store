package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"server/entity"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (ar *AuthRepository) SignUp(user entity.User) (int, error) {
	tx, err := ar.db.Begin()

	if err != nil {
		tx.Rollback()
		return -1, err
	}
	var id int
	query := "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id"
	row := tx.QueryRow(query, user.Username, user.Password, "user")
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	var cartId int
	query = "INSERT INTO cart (user_id) VALUES ($1) RETURNING id"
	row = tx.QueryRow(query, id)
	err = row.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	log.Println("IDDD = ", cartId, id)
	query = "UPDATE users SET cart_id = $1 WHERE id = $2"
	_, err = tx.Exec(query, cartId, id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return id, nil
}

func (ar *AuthRepository) GetUser(username string, passwordHash string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password_hash AS password, role  FROM users WHERE username=$1 AND password_hash=$2"
	err := ar.db.Get(&user, query, username, passwordHash)
	return user, err
}

func (ar *AuthRepository) CreateAdmin(username string, passwordHash string) (int, error) {
	var id int
	query := "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id"
	row := ar.db.QueryRow(query, username, passwordHash, "admin")
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
