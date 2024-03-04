package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"server/entity"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (ar *AuthRepository) Login(user entity.User) (int, error) {
	return 1, gin.Error{}
}

func (ar *AuthRepository) SignUp(user entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id"
	row := ar.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ar *AuthRepository) GetUser(username string, passwordHash string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password_hash AS password FROM users WHERE username=$1 AND password_hash=$2"
	err := ar.db.Get(&user, query, username, passwordHash)
	return user, err
}
