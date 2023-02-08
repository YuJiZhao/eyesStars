package utils

import (
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 14:09
 */

func GetIP(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
