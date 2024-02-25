package service

import "server/repository"

type ProductService struct {
	rep *repository.ProductRepository
}

func NewProductService(productRep *repository.ProductRepository) *ProductService {
	return &ProductService{}
}
