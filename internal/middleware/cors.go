package middleware

import (
	"golang-shop-restful/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(cfg config.Config) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = cfg.AllowOrigins
	config.AddAllowHeaders("Authorization")
	return cors.New(config)
}
