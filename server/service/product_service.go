package service

import (
	"server/entity"
	"server/repository"
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
