package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

var r *gin.Engine

func InitRouter(userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", userHandler.SignUp)
		auth.POST("/sign-in", userHandler.SignIn)
	}

	return router
}
