package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/handler"
	"time"
)

func InitRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.GET("/signup", func(c *gin.Context) {
			c.File("../client/signup.html")
		})
		auth.GET("/login", func(c *gin.Context) {
			c.File("../client/login.html")
		})
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		api.GET("/t", func(c *gin.Context) {
			c.File("../client/index.html")
		})

	}
	router.GET("/mainpage", func(c *gin.Context) {
		c.File("../client/index.html")
	})
	router.Static("/js", "../client/js")

	return router
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(r *gin.Engine, port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
