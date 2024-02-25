package repository

import (
	"github.com/jmoiron/sqlx"
	"server/entity"
)

type Repository struct {
	AuthRepositoryImpl
	UserRepositoryImpl
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepositoryImpl: NewAuthRepository(db),
		UserRepositoryImpl: NewUserRepository(db),
	}
}

type AuthRepositoryImpl interface {
	Login(user entity.User) (int, error)
	SignUp(user entity.User) (int, error)
}

type UserRepositoryImpl interface {
	GetUser(user *entity.User) (int, error)
}