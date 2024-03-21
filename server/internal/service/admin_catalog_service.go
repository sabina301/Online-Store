package service

import (
	"server/internal/entity"
	"server/internal/repository"
	"strconv"
)

type AdminCatalogService struct {
	rep repository.ProductRepositoryImpl
}

func NewAdminCatalogService(rep repository.ProductRepositoryImpl) *AdminCatalogService {
	return &AdminCatalogService{
		rep: rep,
	}
}

func (acs *AdminCatalogService) AddProduct(product entity.Product) (int, error) {
	price, err := strconv.ParseUint(product.Price, 10, 32)
	if err != nil {
		return -1, err
	}
	id, err := acs.rep.AddProduct(product.Category, product.Name, product.Color, product.Description, uint32(price))
	if err != nil {
		return -1, err
	}
	return id, nil
}
