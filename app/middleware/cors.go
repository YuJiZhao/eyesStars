package middleware

import (
	"eyesStars/app/constant"
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
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", constant.Auth.SToken, constant.Auth.LToken}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{constant.Auth.SToken, constant.Auth.LToken}

	return cors.New(config)
}
