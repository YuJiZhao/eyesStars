package routes

import (
	"eyesStars/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 18:03
 */

func SetStarGroupRoutes(router *gin.RouterGroup) {
	// 添加星星
	router.POST("/starAdd", controller.StarAdd)
	// 根据id查询星星
	router.GET("/starGetById/:id", controller.StarGetById)
	// 批量获取星星
	router.GET("/starsGet", controller.StarsGet)
}
