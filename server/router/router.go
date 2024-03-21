package router

import (
	"context"
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
	admin := router.Group("/admin", h.UserIdentity)
	{
		admin.GET("/catalog/edit", func(c *gin.Context) {
			c.File("../client/catalog_edit.html")
		})
		admin.GET("/catalog/add", func(c *gin.Context) {
			c.File("../client/catalog_add.html")
		})
		admin.POST("/catalog/edit/add", h.AddProduct)
	}
	router.GET("/user/cart", func(c *gin.Context) {
		c.File("../client/cart.html")
	})
	router.POST("/user/product/add", h.AddProductInCart)
	router.GET("/user/cart/get", h.GetProductFromCart)
	router.POST("/user/order/make", h.MakeOrder)
	router.GET("/catalog/get/products/all", h.GetAllProducts)
	router.GET("/catalog", func(c *gin.Context) {
		c.File("../client/catalog.html")
	})
	router.GET("/catalog/getrole", h.GetRole)
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

func (s *Server) Finish(c context.Context) error {
	return s.httpServer.Shutdown(c)
}
