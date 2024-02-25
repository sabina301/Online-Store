package handler

import (
	"github.com/gin-gonic/gin"
	"server/service"
)

type Handler struct {
	AuthHandlerImpl
	UserHandlerImpl
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		AuthHandlerImpl: NewAuthHandler(services.AuthServiceImpl),
		UserHandlerImpl: NewUserHandler(services.UserServiceImpl),
	}
}

type AuthHandlerImpl interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type UserHandlerImpl interface {
	GetUser(c *gin.Context)
	Test(c *gin.Context)
}
