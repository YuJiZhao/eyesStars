package middleware

import (
	"eyesStars/app/common/result"
	"eyesStars/global"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 11:08
 */

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.Config.Log.RootDir + "/" + global.Config.Log.LogFolderServer + "/" + time.Now().Format("2006-01-02") + ".log",
			MaxSize:    global.Config.Log.MaxSize,
			MaxBackups: global.Config.Log.MaxBackups,
			MaxAge:     global.Config.Log.MaxAge,
			Compress:   global.Config.Log.Compress,
		},
		serverError)
}

func serverError(c *gin.Context, err interface{}) {
	msg := "server error"
	// 非生产环境显示具体错误信息
	if global.Config.App.Env != "prod" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, result.Response{
		Code: http.StatusInternalServerError,
		Msg:  msg,
	})
	c.Abort()
}
