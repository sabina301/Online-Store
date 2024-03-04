package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/response"
	"strings"
)

// parsing token
func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		response.NewError(c, "Empty auth header", http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewError(c, "Invalid auth header", http.StatusUnauthorized)
		return
	}

	userId, err := h.serv.ParseToken(headerParts[1])
	if err != nil {
		response.NewError(c, "Invalid token", http.StatusUnauthorized)
		return
	}

	c.Set("userId", userId)
}
