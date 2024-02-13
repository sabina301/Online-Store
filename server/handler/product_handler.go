package handler

import "server/service"

type ProductHandler struct {
	services *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{services: service}
}
