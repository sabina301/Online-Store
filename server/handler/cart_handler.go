package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/response"
)

func (h *Handler) GetProductFromCart(c *gin.Context) {
	userId, _ := h.GetUserId(c)
	products, err := h.serv.GetProductFromCart(userId)
	if err != nil {
		response.NewError(c, "invalid products", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, products)
}
