package routes

import (
	"eyesStars/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 14:56
 */

func SetTrackGroupRoutes(router *gin.RouterGroup) {
	// 进站埋点
	router.GET("/visit", controller.TrackVisit)
	// 登录埋点
	router.GET("/login", controller.TrackLogin)
}
