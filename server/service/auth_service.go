package service

import (
	"crypto/sha256"
	"fmt"
	"server/entity"
	"server/repository"
)

const (
	salt = "4444lolkekcheburek4444"
)

type AuthService struct {
	rep repository.AuthRepositoryImpl
}

func NewAuthService(rep repository.AuthRepositoryImpl) *AuthService {
	return &AuthService{rep: rep}
}

func (as *AuthService) SignUp(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.rep.SignUp(user)
}

func (as *AuthService) Login(user entity.User) (int, error) {
	return as.rep.Login(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
