package handler

import (
	"errors"
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

func (h *Handler) GetUserId(c *gin.Context) (int, error) {
	token, err := c.Cookie("token")

	if err != nil {
		return 0, errors.New("user id not found")
	}
	userId, _, err := h.serv.ParseToken(token)
	if err != nil {
		return 0, errors.New("user id not found")
	}
	if err != nil {
		return 0, errors.New("user id is of invalid type")
	}
	return userId, nil
}
