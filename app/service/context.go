package service

import (
	"eyesStars/app/common"
	"eyesStars/app/model/entity"
	"eyesStars/app/model/po"
	"eyesStars/app/model/returnee"
	"eyesStars/global"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:52
 */

type contextService struct {
	// 应用初始化时需要的配置
	initSiteIds []uint32
}

var ContextService = contextService{
	initSiteIds: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
}

// InitSite 应用初始化配置信息
func (contextService contextService) InitSite() (err error, result returnee.ContextInitSite) {
	// 获取指定配置
	var initSitePOList []po.ContextInitSite
	err = global.DB.Table("context").Select([]string{"name", "content"}).Where("id in ?", contextService.initSiteIds).Find(&initSitePOList).Error
	if err != nil {
		global.Log.Error("查询应用初始化配置错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("程序异常"), result
	}

	// 结构转换
	var middleMap = make(map[string]string)
	for _, v := range initSitePOList {
		middleMap[v.Name] = v.Content
	}
	result = returnee.ContextInitSite{}
	if err = mapstructure.Decode(middleMap, &result); err != nil {
		global.Log.Error("map转结构体错误！" + err.Error())
		return common.CustomError{}.SetErrorMsg("程序异常"), result
	}

	return err, result
}

// ContextAdd 添加配置
func (contextService contextService) ContextAdd(name string, content string, remarks interface{}) error {
	// 判断配置是否已经存在
	var context = entity.Context{}
	if tx := global.DB.Where(entity.Context{Name: name}).First(&context); tx.RowsAffected != 0 {
		return common.CustomError{}.SetErrorMsg("该配置已存在！id:" + strconv.Itoa(int(context.Id)) + ";content:" + context.Content + ";remarks:" + context.Remarks)
	}

	// 添加配置
	context = entity.Context{
		Name:    name,
		Content: content,
	}
	if remarks != nil {
		context.Remarks = remarks.(string)
	}
	err := global.DB.Create(&context).Error
	if err != nil {
		global.Log.Error("配置添加错误！" + err.Error())
		return err
	}

	// 写入日志，这个日志比较重要，因此出现异常需要返回
	if err = LogService.Context(context); err != nil {
		global.Log.Error("历史配置插入错误！" + err.Error())
		return err
	}

	return err
}

// ContextUpdate 修改配置
func (contextService contextService) ContextUpdate(name string, content string) error {
	// 判断配置是否已经存在
	var context = entity.Context{}
	tx := global.DB.Where(entity.Context{Name: name}).First(&context)
	if err := tx.Error; err != nil {
		global.Log.Error("程序错误！" + err.Error())
		// 这是管理员接口，因此直接返回完整真实错误信息
		return err
	}
	if tx.RowsAffected == 0 {
		return common.CustomError{}.SetErrorMsg("该配置不存在！请调用添加配置接口")
	}

	// 判断新配置是否与旧配置相同
	if content == context.Content {
		return common.CustomError{}.SetErrorMsg("配置内容与原内容相同，本次操作无效")
	}

	// 修改配置
	err := global.DB.Model(&context).Updates(entity.Context{
		Version: context.Version + 1,
		Content: content,
	}).Error
	if err != nil {
		global.Log.Error("修改配置错误！" + err.Error())
		return err
	}

	fmt.Println(context)

	// 记录日志
	if err = LogService.Context(context); err != nil {
		global.Log.Error("历史配置插入错误！" + err.Error())
		return err
	}

	return err
}
