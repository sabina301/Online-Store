package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUser(user *entity.User) (int, error) {
	return 1, errors.New("lol")
}
