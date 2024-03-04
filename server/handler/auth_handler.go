package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/response"
	"server/service"
)

type inputUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input inputUser
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.serv.GenerateToken(input.Username, input.Password)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.User
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.serv.SignUp(input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetService() service.AuthServiceImpl {
	return h.serv
}
