package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/response"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.serv.GetAllProducts()
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, products)
}
