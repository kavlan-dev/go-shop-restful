package router

import (
	"fmt"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerInterface interface {
	productHandler
	userHandler
	cartHandler
}

type productHandler interface {
	Products(c *gin.Context)
	ProductById(c *gin.Context)
	ProductByTitle(c *gin.Context)
	PostProduct(c *gin.Context)
	PutProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type userHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	PromoteToAdmin(c *gin.Context)
	DowngradeToCustomer(c *gin.Context)
}

type cartHandler interface {
	Cart(c *gin.Context)
	AddToCart(c *gin.Context)
	ClearCart(c *gin.Context)
}

func Router(cfg *config.Config, handler handlerInterface) (*http.Server, error) {
	var r *gin.Engine
	switch cfg.Environment {
	case "dev":
		r = gin.Default()
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger(), gin.Recovery())
	default:
		return nil, fmt.Errorf("Не известное окружение %s", cfg.Environment)
	}
	r.Use(middleware.CORSMiddleware(cfg.CORS))

	auth := r.Group("/api/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	admin := api.Group("/admin")
	admin.Use(middleware.AdminMiddleware())

	product := api.Group("/products")
	product.GET("/", handler.Products)
	product.GET("/:id", handler.ProductById)
	product.GET("/title/:title", handler.ProductByTitle)

	cart := api.Group("/cart")
	cart.GET("/", handler.Cart)
	cart.POST("/:id", handler.AddToCart)
	cart.DELETE("/", handler.ClearCart)

	adminUsers := admin.Group("/users")
	adminUsers.POST("/:id/promote", handler.PromoteToAdmin)
	adminUsers.POST("/:id/downgrade", handler.DowngradeToCustomer)

	adminProduct := admin.Group("/products")
	adminProduct.POST("/", handler.PostProduct)
	adminProduct.PUT("/:id", handler.PutProduct)
	adminProduct.DELETE("/:id", handler.DeleteProduct)

	server := &http.Server{
		Addr:    cfg.ServerAddress(),
		Handler: r,
	}

	return server, nil
}
