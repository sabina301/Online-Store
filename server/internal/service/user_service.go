package service

import (
	"errors"
	"server/internal/entity"
	"server/internal/repository"
)

type UserService struct {
	rep repository.UserRepositoryImpl
}

func NewUserService(rep repository.UserRepositoryImpl) *UserService {
	return &UserService{rep}
}

func (us *UserService) GetUser(user *entity.User) (int, error) {
	us.rep.GetUser(user)
	return 1, errors.New("lol")
}

func (us *UserService) Test() string {
	return "kek"
}
