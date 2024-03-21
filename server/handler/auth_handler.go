package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/response"
)

type inputUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input inputUser
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, "invalid input body", http.StatusBadRequest)
		return
	}
	token, _, err := h.serv.GenerateToken(input.Username, input.Password)
	if err != nil {
		response.NewError(c, "invalid username or password", http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.User
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, "invalid input body", http.StatusBadRequest)
		return
	}
	id, err := h.serv.SignUp(input)
	if err != nil {
		response.NewError(c, "duplicate username value", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetRole(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		response.NewError(c, err.Error(), http.StatusUnauthorized)
	}
	id, role, err := h.serv.ParseToken(token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"role": role,
	})
}
