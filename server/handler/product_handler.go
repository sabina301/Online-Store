package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/response"
)

type ProductResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.serv.GetAllProducts()

	pr := make([]ProductResponse, len(products))
	for i, product := range products {
		pr[i].Id = product.Id
		pr[i].Name = product.Name
		pr[i].Price = product.Price
	}
	if err != nil {
		response.NewError(c, "invalid products", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, pr)
}

func (h *Handler) AddProductInCart(c *gin.Context) {
	jsonStr, err := c.GetRawData()
	var data map[string]int

	err = json.Unmarshal([]byte(jsonStr), &data)

	userId := data["userId"]
	if err != nil {
		response.NewError(c, "user is not founded", http.StatusNotFound)
		return
	}
	productId := data["productId"]
	if err != nil {
		response.NewError(c, "product is not founded", http.StatusNotFound)
		return
	}
	err = h.serv.AddProductInCart(userId, productId)
	if err != nil {
		response.NewError(c, "invalid products", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
