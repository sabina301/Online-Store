package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/response"
	"server/service"
)

type UserHandler struct {
	serv service.UserServiceImpl
}

func NewUserHandler(serv service.UserServiceImpl) *UserHandler {
	return &UserHandler{serv}
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	var input *entity.User
	err := c.BindJSON(input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
	}
	uh.serv.GetUser(input)
}

func (uh *UserHandler) Test(c *gin.Context) {
	str := uh.serv.Test()
	c.JSON(http.StatusOK, str)
}
