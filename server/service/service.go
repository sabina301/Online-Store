package service

import (
	"server/entity"
	"server/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service struct {
	AuthServiceImpl
	UserServiceImpl
	AdminCatalogServiceImpl
	ProductServiceImpl
	CartServiceImpl
	OrderServiceImpl
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		AuthServiceImpl:         NewAuthService(rep.AuthRepositoryImpl),
		UserServiceImpl:         NewUserService(rep.UserRepositoryImpl),
		AdminCatalogServiceImpl: NewAdminCatalogService(rep.ProductRepositoryImpl),
		ProductServiceImpl:      NewProductService(rep.ProductRepositoryImpl),
		CartServiceImpl:         NewCartService(rep.CartRepositoryImpl),
		OrderServiceImpl:        NewOrderService(rep.OrderRepositoryImpl),
	}
}

type AuthServiceImpl interface {
	SignUp(user entity.User) (int, error)
	GenerateToken(username string, password string) (string, int, error)
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
	AddProductInCart(userId int, productId int) error
}

type CartServiceImpl interface {
	GetProductFromCart(userId int) ([]entity.Product, error)
}

type OrderServiceImpl interface {
	MakeOrder(userId int) (int, error)
}
