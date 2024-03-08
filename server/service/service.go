package service

import (
	"server/entity"
	"server/repository"
)

type Service struct {
	AuthServiceImpl
	UserServiceImpl
	AdminCatalogServiceImpl
	ProductServiceImpl
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		AuthServiceImpl:         NewAuthService(rep.AuthRepositoryImpl),
		UserServiceImpl:         NewUserService(rep.UserRepositoryImpl),
		AdminCatalogServiceImpl: NewAdminCatalogService(rep.ProductRepositoryImpl),
		ProductServiceImpl:      NewProductService(rep.ProductRepositoryImpl),
	}
}

type AuthServiceImpl interface {
	Login(user entity.User) (int, error)
	SignUp(user entity.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, string, error)
	CreateAdmin() (int, error)
}

type UserServiceImpl interface {
	GetUser(user *entity.User) (int, error)
	Test() string
}

type AdminCatalogServiceImpl interface {
	AddProduct(product entity.Product) (int, error)
}

type ProductServiceImpl interface {
	GetAllProducts() ([]entity.Product, error)
}
