package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 10:16
 */

// Response 响应结构体
type Response struct {
	Code int         `json:"code"` // 自定义错误码
	Msg  string      `json:"msg"`  // 信息
	Data interface{} `json:"data"` // 数据
}

// Success 成功响应
func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// SuccessByCustom 成功响应(使用customResult信息)
func SuccessByCustom(c *gin.Context, result customResult, data interface{}) {
	Success(c, result.code, result.msg, data)
}

// SuccessDefault 成功响应(默认模式)
func SuccessDefault(c *gin.Context, data interface{}) {
	SuccessByCustom(c, Results.DefaultSuccess, data)
}

// SuccessAttachedCode 成功响应(默认模式，自选状态码)
func SuccessAttachedCode(c *gin.Context, data interface{}, code int) {
	Success(c, code, Results.DefaultSuccess.msg, data)
}

// SuccessAttachedMsg 成功响应(默认模式，自选信息)
func SuccessAttachedMsg(c *gin.Context, data interface{}, msg string) {
	Success(c, Results.DefaultSuccess.code, msg, data)
}

// Fail 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		nil,
	})
}

// FailByCustom 失败响应(使用customResult信息)
func FailByCustom(c *gin.Context, result customResult) {
	Fail(c, result.code, result.msg)
}

// FailDefault 失败响应(默认模式)
func FailDefault(c *gin.Context) {
	FailByCustom(c, Results.DefaultFail)
}

// FailAttachedCode 失败响应(默认模式，自选状态码)
func FailAttachedCode(c *gin.Context, code int) {
	Fail(c, code, Results.DefaultFail.msg)
}

// FailAttachedMsg 失败响应(默认模式，自选信息)
func FailAttachedMsg(c *gin.Context, msg string) {
	Fail(c, Results.DefaultFail.code, msg)
}
