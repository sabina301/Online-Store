package service

import (
	"server/entity"
	"server/repository"
)

type AuthService struct {
	rep repository.AuthRepositoryImpl
}

func NewAuthService(rep repository.AuthRepositoryImpl) *AuthService {
	return &AuthService{rep: rep}
}

func (as *AuthService) SignUp(user *entity.User) (int, error) {
	return as.rep.SignUp(*user)
}

func (as *AuthService) Login(user *entity.User) (int, error) {
	return as.rep.Login(*user)
}
