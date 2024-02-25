package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/response"
	"server/service"
)

type AuthHandler struct {
	serv service.AuthServiceImpl
}

func NewAuthHandler(serv service.AuthServiceImpl) *AuthHandler {
	return &AuthHandler{serv}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var input *entity.User
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
	}
	_, err = ah.serv.Login(input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
	}
}

func (ah *AuthHandler) SignUp(c *gin.Context) {

}
