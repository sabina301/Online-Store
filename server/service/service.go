package service

import (
	"server/entity"
	"server/repository"
)

type Service struct {
	AuthServiceImpl
	UserServiceImpl
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthServiceImpl: NewAuthService(repo.AuthRepositoryImpl),
		UserServiceImpl: NewUserService(repo.UserRepositoryImpl),
	}
}

type AuthServiceImpl interface {
	Login(user entity.User) (int, error)
	SignUp(user entity.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserServiceImpl interface {
	GetUser(user *entity.User) (int, error)
	Test() string
}
