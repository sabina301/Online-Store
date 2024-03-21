package service

import (
	"server/repository"
)

type OrderService struct {
	rep repository.OrderRepositoryImpl
}

func NewOrderService(rep repository.OrderRepositoryImpl) *OrderService {
	return &OrderService{rep: rep}
}

func (s *OrderService) MakeOrder(userId int) (int, error) {
	products, err := s.rep.GetProductFromCart(userId)
	if err != nil {
		return -1, err
	}
	orderId, err := s.rep.MakeOrder(userId, products)
	if err != nil {
		return -1, err
	}
	return orderId, nil
}
