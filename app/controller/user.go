package controller

import (
	"eyesStars/app/common"
	"eyesStars/app/common/result"
	"eyesStars/app/common/validate"
	"eyesStars/app/constant"
	"eyesStars/app/model/receiver"
	"eyesStars/app/service"
	"eyesStars/global"
	"github.com/gin-gonic/gin"
	"strings"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:51
 */

// UserInfoGet
// @Summary			获取用户基本信息
// @Description		权限：User
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header  	string				true		"短token"
// @Param			lToken		header  	string				true		"长token"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/user/userInfoGet [get]
func UserInfoGet(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.User) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 获取用户信息
	uid, _ := c.Get(constant.Auth.CUid)

	// 执行业务逻辑
	if err, resultInfo := service.UserService.GetInfo(c, uid.(uint32)); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, resultInfo)
	}
}

// UserInfoUpdate
// @Summary			修改用户基本信息
// @Description		权限：User
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header  	string					true		"短token"
// @Param			lToken		header  	string					true		"长token"
// @Param			receiver	body    	receiver.UserInfoUpdate	true		"请求参数"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/user/userInfoUpdate [put]
func UserInfoUpdate(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.User) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.UserInfoUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 获取用户信息
	uid, _ := c.Get(constant.Auth.CUid)

	// 执行业务逻辑
	if err := service.UserService.UpdateInfo(c, data, uid.(uint32)); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, nil)
	}
}

// UserAvatarUpdate
// @Summary			修改用户头像
// @Description		权限：User
// @Tags			user
// @Accept			multipart/form-data
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header  	string				true		"短token"
// @Param			lToken		header  	string				true		"长token"
// @Param			file		formData	file				true		"头像文件"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/user/userUpdateAvatar [post]
func UserAvatarUpdate(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.User) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	file, fileHeader, ValidateErr := c.Request.FormFile("file")
	if ValidateErr != nil {
		global.Log.Info("没有传入头像文件！" + ValidateErr.Error())
		result.FailAttachedMsg(c, "请传入头像文件")
		return
	}
	filename := fileHeader.Filename
	index := strings.LastIndex(filename, ".")
	if index == -1 || !common.FitImgType(filename[index:]) {
		global.Log.Warn("非法头像文件类型！" + filename)
		result.FailAttachedMsg(c, "头像文件类型不支持")
		return
	}

	// 获取用户信息
	uid, _ := c.Get(constant.Auth.CUid)

	// 执行业务逻辑
	if err, resultInfo := service.UserService.UpdateAvatar(c, file, fileHeader, uid.(uint32)); err != nil {
		result.FailAttachedMsg(c, err.Error())
		go service.LogService.File(uid.(uint32), constant.LogFile.UploadAvatar, false, nil)
	} else {
		result.SuccessDefault(c, resultInfo)
		go service.LogService.File(uid.(uint32), constant.LogFile.UploadAvatar, true, resultInfo)
	}
}
