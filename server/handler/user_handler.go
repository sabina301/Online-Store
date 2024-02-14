package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/service"
)

type UserHandler struct {
	services *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{services: service}
}

func (uh *UserHandler) Login(c *gin.Context) {

}

func (uh *UserHandler) SignUp(c *gin.Context) {

}

func (uh *UserHandler) Test(c *gin.Context) {
	fmt.Println("!!!!! TEST")
}
