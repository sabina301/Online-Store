package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/entity"
	"server/internal/response"
)

func (h *Handler) GetUser(c *gin.Context) {
	var input *entity.User
	err := c.BindJSON(input)
	if err != nil {
		response.NewError(c, err.Error(), http.StatusBadRequest)
	}
	h.serv.GetUser(input)
}

func (h *Handler) Test(c *gin.Context) {
	id, _ := c.Get("userId")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
