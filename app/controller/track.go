package controller

import (
	"eyesStars/app/common"
	"eyesStars/app/common/result"
	"eyesStars/app/constant"
	"eyesStars/app/service"
	"eyesStars/app/utils"
	"eyesStars/global"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 14:55
 */

// TrackVisit
// @Summary			进站埋点
// @Description		权限：Visitor，无响应内容
// @Tags			track
// @Accept			json
// @Security		sToken,lToken
// @Param			sToken		header		string				false		"短token"
// @Param			lToken		header		string				false		"长token"
// @Param			package		query		string				true		"埋点数据加密包"
// @Router			/track/visit [get]
func TrackVisit(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Visitor) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 获取用户信息
	ip := utils.GetIP(c)
	uid, _ := c.Get(constant.Auth.CUid)

	// 参数校验
	encryptPkg, _ := url.QueryUnescape(c.Query("package"))
	err, decryptPkg := utils.AesDecrypt(encryptPkg)
	if err != nil {
		var uidStr string
		if uid != nil {
			uidStr = "nil"
		} else {
			uidStr = strconv.Itoa(uid.(int))
		}
		global.Log.Warn("非法埋点请求！uid:" + uidStr + ";" + "IP:" + ip + ";package:" + encryptPkg)
		return
	}

	// 执行业务
	service.TrackService.TrackVisit(uid, ip, decryptPkg)
}

// TrackLogin
// @Summary			登录埋点
// @Description		权限：Visitor，无响应内容
// @Tags			track
// @Accept			json
// @Security		sToken,lToken
// @Param			sToken		header		string				false		"短token"
// @Param			lToken		header		string				false		"长token"
// @Param			package		query		string				true		"埋点数据加密包"
// @Router			/track/login [get]
func TrackLogin(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Visitor) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 获取用户信息
	ip := utils.GetIP(c)
	uid, _ := c.Get(constant.Auth.CUid)

	// 参数校验
	encryptPkg, _ := url.QueryUnescape(c.Query("package"))
	err, decryptPkg := utils.AesDecrypt(encryptPkg)
	if err != nil {
		var uidStr string
		if uid != nil {
			uidStr = "nil"
		} else {
			uidStr = strconv.Itoa(uid.(int))
		}
		global.Log.Warn("非法埋点请求！uid:" + uidStr + ";" + "IP:" + ip + ";package:" + encryptPkg)
		return
	}

	// 执行业务
	service.TrackService.TrackLogin(uid, ip, decryptPkg)
}
