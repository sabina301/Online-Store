package service

import "server/repository"

type UserService struct {
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{}
}
