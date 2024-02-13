package service

import "server/repository"

type ProductService struct {
}

func NewProductService(productRep *repository.ProductRepository) *ProductService {
	return &ProductService{}
}
