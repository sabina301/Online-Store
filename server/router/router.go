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
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
	}
	api := router.Group("/test", h.UserIdentity)
	{
		api.GET("/t", h.Test)
	}

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
