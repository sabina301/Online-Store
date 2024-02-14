package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

func InitRouter(userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", userHandler.SignUp)
		auth.POST("/login", userHandler.Login)
	}

	router.GET("/", userHandler.Test)

	return router
}

func Start(router *gin.Engine, addr string) error {
	return router.Run(addr)
}
