package repository

import "github.com/jmoiron/sqlx"

type UserRepository struct {
}

func NewUserRepository(database *sqlx.DB) *UserRepository {
	return &UserRepository{}
}
