package controller

import (
	"eyesStars/app/common"
	"eyesStars/app/common/result"
	"eyesStars/app/common/validate"
	"eyesStars/app/constant"
	"eyesStars/app/model/receiver"
	"eyesStars/app/service"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:52
 */

// ContextInitSite
// @Summary			网站信息初始化
// @Description		权限：Visitor；该接口处理方式要求前端必须是服务端渲染，因为接口会明文返回密钥等信息
// @Tags			context
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string		false		"短token"
// @Param			lToken		header		string 		false		"长token"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/context/initSite [get]
func ContextInitSite(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Visitor) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 执行业务
	if err, returnInfo := service.ContextService.InitSite(); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, returnInfo)
	}
}

// ContextAdd
// @Summary			添加配置
// @Description		权限：Admin
// @Tags			context
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string					true		"短token"
// @Param			lToken		header		string 					true		"长token"
// @Param			receiver	body		receiver.ContextAdd		true		"请求参数"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/context/contextAdd [post]
func ContextAdd(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Admin) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.ContextAdd
	if err := c.ShouldBindJSON(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 执行业务
	if err := service.ContextService.ContextAdd(data.Name, data.Content, data.Remarks); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, nil)
	}
}

// ContextUpdate
// @Summary			修改配置
// @Description		权限：Admin；若表中name不存在则添加配置，已存在则修改配置
// @Tags			context
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string		true			"短token"
// @Param			lToken		header		string 		true			"长token"
// @Param			receiver	body		receiver.ContextUpdate		true		"请求参数"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/context/contextUpdate [put]
func ContextUpdate(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Admin) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.ContextUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 执行业务
	if err := service.ContextService.ContextUpdate(data.Name, data.Content); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, nil)
	}
}
