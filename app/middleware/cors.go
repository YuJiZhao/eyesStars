package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 11:08
 */

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "sToken", "lToken"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"sToken", "lToken"}

	return cors.New(config)
}
