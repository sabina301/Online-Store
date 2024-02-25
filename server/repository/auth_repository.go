package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"server/entity"
)

type AuthRepository struct {
	db *sqlx.DB
}

func (ar *AuthRepository) Login(user entity.User) (int, error) {
	return 1, gin.Error{}
}

func (ar *AuthRepository) SignUp(user entity.User) (int, error) {
	return 1, gin.Error{}
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}
