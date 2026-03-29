package router

import (
	"fmt"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/middleware"

	"github.com/gin-gonic/gin"
)

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

func NewRouter(cfg *config.Config, productHandler productHandler, userHandler userHandler, cartHandler cartHandler) (*gin.Engine, error) {
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

	r.GET("/api/products", productHandler.Products)

	auth := r.Group("/api/auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	admin := api.Group("/admin")
	admin.Use(middleware.AdminMiddleware())

	product := api.Group("/products")
	product.GET("/:id", productHandler.ProductById)
	product.GET("/title/:title", productHandler.ProductByTitle)

	cart := api.Group("/cart")
	cart.GET("/", cartHandler.Cart)
	cart.POST("/:id", cartHandler.AddToCart)
	cart.DELETE("/", cartHandler.ClearCart)

	adminUsers := admin.Group("/users")
	adminUsers.POST("/:id/promote", userHandler.PromoteToAdmin)
	adminUsers.POST("/:id/downgrade", userHandler.DowngradeToCustomer)

	adminProduct := admin.Group("/products")
	adminProduct.POST("/", productHandler.PostProduct)
	adminProduct.PUT("/:id", productHandler.PutProduct)
	adminProduct.DELETE("/:id", productHandler.DeleteProduct)

	return r, nil
}
