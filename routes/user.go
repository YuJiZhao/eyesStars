package routes

import (
	"eyesStars/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 18:02
 */

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 获取用户基本信息
	router.GET("/userInfoGet", controller.UserInfoGet)
	// 修改用户基本信息
	router.PUT("/userInfoUpdate", controller.UserInfoUpdate)
	// 修改用户头像
	router.POST("/userAvatarUpdate", controller.UserAvatarUpdate)
}
