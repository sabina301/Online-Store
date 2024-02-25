package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	message string `json:"message"`
}

func NewError(c *gin.Context, message string, status int) {
	logrus.Error(message)
	c.AbortWithStatusJSON(status, error{message: message})
}
