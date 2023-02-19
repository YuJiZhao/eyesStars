package controller

import (
	"eyesStars/app/common"
	"eyesStars/app/common/result"
	"eyesStars/app/common/validate"
	"eyesStars/app/constant"
	"eyesStars/app/model/receiver"
	"eyesStars/app/service"
	"eyesStars/app/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:51
 */

// StarAdd
// @Summary			添加星星
// @Description		权限：User
// @Tags			star
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header  	string				true		"短token"
// @Param			lToken		header  	string				true		"长token"
// @Param			receiver	body    	receiver.StarAdd	true		"请求参数"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/star/starAdd [post]
func StarAdd(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.User) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.StarAdd
	if err := c.ShouldBindJSON(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 获取用户信息
	uid, _ := c.Get(constant.Auth.CUid)

	// 执行业务
	if err, resultInfo := service.StarService.AddStar(data, uid.(uint32)); err != nil {
		result.FailAttachedMsg(c, err.Error())
		go service.LogService.Star(uid, nil, constant.LogStar.Search, false, utils.GetIP(c))
	} else {
		result.SuccessDefault(c, resultInfo)
		go service.LogService.Star(uid, resultInfo.Id, constant.LogStar.Search, true, utils.GetIP(c))
	}
}

// StarGetById
// @Summary			通过id获取星星
// @Description		权限：Admin
// @Tags			star
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string				false		"短token"
// @Param			lToken		header		string				false		"长token"
// @Param			id			path		integer				true		"寄语id"		minimum(1)
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/star/starGetById/{id} [get]
func StarGetById(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Admin) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.StarGet
	if err := c.ShouldBindUri(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 执行业务
	var flag bool
	if err, resultInfo := service.StarService.GetStarById(data.Id); err != nil {
		flag = false
		result.FailAttachedMsg(c, err.Error())
	} else {
		flag = true
		result.SuccessDefault(c, resultInfo)
	}

	// 异步写入日志
	uid, _ := c.Get(constant.Auth.CUid)
	go service.LogService.Star(uid, data.Id, constant.LogStar.Search, flag, utils.GetIP(c))
}

// StarsGet
// @Summary			批量查询星星
// @Description		权限：Visitor
// @Tags			star
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string		false		"短token"
// @Param			lToken		header		string 		false		"长token"
// @Param			ids			query		string		false		"已读id密串"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/star/starsGet [get]
func StarsGet(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Visitor) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验(ids真实结构示例： "1,5,10")
	ids := c.Query("ids")
	if ids == "" {
		// ids要放到sql: not in(ids) 里执行，如果为空会报错，因此给它默认值
		// 数据库中star id为无符号整形，因此取0不影响业务
		ids = "0"
	} else {
		err, decrypt := utils.AesDecrypt(ids)
		if err != nil {
			result.FailAttachedMsg(c, "猜猜是怎么加密的")
			return
		}
		ids = decrypt
	}

	// 执行业务
	if err, resultInfo := service.StarService.GetStars(c, ids); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, resultInfo)
	}
}
