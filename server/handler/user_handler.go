package handler

import (
	"github.com/gin-gonic/gin"
	"server/service"
)

type UserHandler struct {
	services *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{services: service}
}

func (uh *UserHandler) SignIn(c *gin.Context) {

}

func (uh *UserHandler) SignUp(c *gin.Context) {

}
