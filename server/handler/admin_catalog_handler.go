package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/entity"
	"server/response"
)

func (h *Handler) AddProduct(c *gin.Context) {
	var input entity.Product
	err := c.BindJSON(&input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.serv.AddProduct(input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
