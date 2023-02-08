package common

import (
	"eyesStars/app/constant"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 17:44
 */

func CheckAuth(c *gin.Context, role string) bool {
	// 没有token的情况
	if flag, exists := c.Get(constant.Auth.AuthFlag); !(flag.(bool) && exists) {
		return role == constant.Roles.Visitor
	}

	// 有token的情况
	tRole, _ := c.Get(constant.Auth.CRole)
	if role == constant.Roles.Admin && tRole != constant.Roles.Admin {
		return false
	}

	return true
}
