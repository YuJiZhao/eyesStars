package routes

import (
	"eyesStars/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:52
 */

func SetContextGroupRoutes(router *gin.RouterGroup) {
	// 获取配置
	router.GET("/initSite", controller.ContextInitSite)
	// 添加配置
	router.POST("/contextAdd", controller.ContextAdd)
	// 修改配置
	router.PUT("/contextUpdate", controller.ContextUpdate)
}
