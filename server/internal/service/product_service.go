package service

import (
	"server/internal/entity"
	"server/internal/repository"
)

type ProductService struct {
	rep repository.ProductRepositoryImpl
}

func NewProductService(rep repository.ProductRepositoryImpl) *ProductService {
	return &ProductService{rep: rep}
}

func (ps *ProductService) GetAllProducts() ([]entity.Product, error) {
	products, err := ps.rep.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) AddProductInCart(userId int, productId int) error {
	err := ps.rep.AddProductInCart(userId, productId)
	if err != nil {
		return err
	}
	return nil
}
