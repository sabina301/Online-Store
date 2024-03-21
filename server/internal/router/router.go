package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/handler"

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
		catalog := admin.Group("/catalog")
		{
			catalog.GET("/edit", func(c *gin.Context) {
				c.File("../client/catalog_edit.html")
			})
			catalog.GET("/add", func(c *gin.Context) {
				c.File("../client/catalog_add.html")
			})
			catalog.POST("/edit/add", h.AddProduct)
		}
		order := admin.Group("/order")
		{
			order.GET("/all", h.GetAllOrdersAdmin)
			order.GET("", func(c *gin.Context) {
				c.File("../client/admin_order.html")
			})
		}
	}

	user := router.Group("/user")
	{
		cart := user.Group("/cart")
		{
			cart.GET("/get", h.GetProductFromCart)
			cart.GET("", func(c *gin.Context) {
				c.File("../client/cart.html")
			})
		}
		product := user.Group("/product")
		{
			product.POST("/add", h.AddProductInCart)
		}
		order := user.Group("/order")
		{
			order.POST("/make", h.MakeOrder)
			order.GET("/all", h.GetAllOrdersUser)
			order.GET("", func(c *gin.Context) {
				c.File("../client/user_order.html")
			})
		}
	}

	catalog := router.Group("/catalog")
	{
		catalog.GET("/product/get/all", h.GetAllProducts)
		catalog.GET("", func(c *gin.Context) {
			c.File("../client/catalog.html")
		})
		catalog.GET("/getrole", h.GetRole)
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

func (s *Server) Finish(c context.Context) error {
	return s.httpServer.Shutdown(c)
}
