package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

func InitRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
	}
	router.GET("/", h.Test)

	return router
}

func Start(router *gin.Engine, addr string) error {
	return router.Run(addr)
}
