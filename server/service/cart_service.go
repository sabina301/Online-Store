package service

import (
	"server/entity"
	"server/repository"
)

type CartService struct {
	rep repository.CartRepositoryImpl
}

func NewCartService(rep repository.CartRepositoryImpl) *CartService {
	return &CartService{rep: rep}
}

func (c *CartService) GetProductFromCart(userId int) ([]entity.Product, error) {
	products, err := c.rep.GetProductFromCart(userId)
	if err != nil {
		return nil, err
	}
	return products, nil
}
