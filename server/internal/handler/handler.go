package handler

import (
	"server/internal/service"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		serv: services,
	}
}
