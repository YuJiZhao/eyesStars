package service

import (
	"eyesStars/app/common"
	"eyesStars/app/model/receiver"
	"eyesStars/app/rpc/generate/userThrift"
	"eyesStars/app/rpc/rpc"
	"eyesStars/global"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:52
 */

type userService struct {
}

var UserService = new(userService)

// GetInfo 获取用户基本信息
func (userService userService) GetInfo(c *gin.Context, uid uint32) (err error, result *userThrift.UserInfoReturnee) {
	// 连接耶瞳用户中心
	err, client, transport := rpc.User()
	if err != nil {
		return common.CustomError{}.SetErrorMsg("用户中心服务似乎宕机了"), result
	}
	defer func(transport thrift.TTransport) {
		closeErr := transport.Close()
		if closeErr != nil {
			global.Log.Error("UserService:thrift连接关闭异常！" + closeErr.Error())
		}
	}(transport)

	// 获取用户基本信息
	result, err = client.GetUserInfo(c, int64(uid))
	if err != nil {
		global.Log.Error("UserService.GetUserInfo:thrift调用错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("用户信息获取错误"), result
	}
	return err, result
}

// UpdateInfo 修改用户基本信息
func (userService userService) UpdateInfo(c *gin.Context, data receiver.UserInfoUpdate, uid uint32) error {
	// 连接耶瞳用户中心
	err, client, transport := rpc.User()
	if err != nil {
		return common.CustomError{}.SetErrorMsg("用户中心服务似乎宕机了")
	}
	defer func(transport thrift.TTransport) {
		closeErr := transport.Close()
		if closeErr != nil {
			global.Log.Error("UserService:thrift连接关闭异常！" + closeErr.Error())
		}
	}(transport)

	// 获取用户基本信息
	userUpdateInfoBO := userThrift.UserUpdateInfoReceiver{
		UID:  int64(uid),
		Name: data.Name,
	}
	err = client.UpdateUserInfo(c, &userUpdateInfoBO)
	if err != nil {
		global.Log.Error("UserService.UpdateUserInfo:thrift调用错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("用户信息修改失败")
	}
	return err
}

// UpdateAvatar 修改用户头像
func (userService userService) UpdateAvatar(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader, uid uint32) (err error, result string) {
	// 连接耶瞳用户中心
	linkErr, client, transport := rpc.User()
	if linkErr != nil {
		return common.CustomError{}.SetErrorMsg("用户中心服务似乎宕机了"), result
	}
	defer func(transport thrift.TTransport) {
		closeErr := transport.Close()
		if closeErr != nil {
			global.Log.Error("UserService:thrift连接关闭异常！" + closeErr.Error())
		}
	}(transport)

	// 文件处理：获取[]byte格式文件
	byteData := make([]byte, fileHeader.Size)
	if _, err = file.Read(byteData); err != nil {
		global.Log.Error("multipart.File转[]byte错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("文件处理错误"), result
	}

	// 修改用户头像
	result, err = client.UpdateUserAvatar(c, int64(uid), byteData)
	if err != nil {
		global.Log.Error("UserService.UpdateUserAvatar: thrift调用错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("头像修改错误"), result
	}

	return err, result
}
