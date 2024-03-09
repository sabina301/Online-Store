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
	cookie, _ := c.Cookie("token")
	if header == "" && cookie == "" {
		response.NewError(c, "empty auth header name or httpOnlyCookie", http.StatusUnauthorized)
		return
	}
	var token string
	if header != "" {
		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[1] == " " || headerParts[1] == "" {
			response.NewError(c, "token is empty", http.StatusUnauthorized)
			return
		}
		if headerParts[0] != "Bearer" {
			response.NewError(c, "invalid header value", http.StatusUnauthorized)
			return
		}
		token = headerParts[1]
	} else {
		token = cookie
	}

	userId, userRole, err := h.serv.ParseToken(token)
	if err != nil {
		response.NewError(c, "invalid token", http.StatusUnauthorized)
		return
	}
	c.Set("userId", userId)
	c.Set("role", userRole)
}
