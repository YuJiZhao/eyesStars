package middleware

import (
	"eyesStars/app/common/result"
	"eyesStars/app/constant"
	"eyesStars/app/rpc/rpc"
	"eyesStars/global"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 17:05
 */

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sToken := c.Request.Header.Get(constant.Auth.SToken)
		lToken := c.Request.Header.Get(constant.Auth.LToken)

		// 如果缺失token，则终止鉴权，由controller决定接口是否允许无权限访问
		if sToken == "" || lToken == "" {
			c.Set(constant.Auth.AuthFlag, false)
			return
		}
		c.Set(constant.Auth.AuthFlag, true)

		// 连接耶瞳用户中心
		err, client, transport := rpc.Auth()
		if err != nil {
			global.Log.Error("thrift权限校验调用错误！" + err.Error())
			result.FailByCustom(c, result.Results.ServerError)
			c.Abort()
			return
		}
		defer func(transport thrift.TTransport) {
			closeErr := transport.Close()
			if closeErr != nil {
				global.Log.Error("AuthService:thrift连接关闭异常！" + closeErr.Error())
			}
		}(transport)

		// 权限校验，通过则获取身份信息
		info, err := client.CheckAuthByDouble(c, global.Config.App.Name, sToken, lToken)
		if err != nil {
			global.Log.Info("thrift鉴权错误！" + err.Error())
			result.FailAttachedMsg(c, err.Error())
			c.Abort()
			return
		}

		// 设置sToken、lToken响应头
		c.Header(constant.Auth.SToken, info.SToken)
		c.Header(constant.Auth.LToken, info.LToken)
		// 设置用户信息
		c.Set(constant.Auth.CUid, uint32(info.UID))
		c.Set(constant.Auth.CRole, info.Role)
	}
}
