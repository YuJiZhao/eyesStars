package service

import (
	"eyesStars/app/model/entity"
	"eyesStars/global"
)

/**
 * @author eyesYeager
 * @date 2023/1/31 14:16
 */

type logService struct {
}

var LogService = new(logService)

// Visit 进站日志
func (logService logService) Visit(uid interface{}, path string, ip string, os string, browser string) {
	log := entity.LogVisit{
		Path:    path,
		Ip:      ip,
		OS:      os,
		Browser: browser,
	}
	if uid != nil {
		log.Uid = uid.(uint32)
	}

	// 插入数据库
	if err := global.DB.Create(&log).Error; err != nil {
		global.Log.Error("进站日志插入失败！" + err.Error())
	}
}

// Login 登录日志
func (logService logService) Login(uid interface{}, pattern uint8, path string, ip string, os string, browser string) {
	log := entity.LogLogin{
		Pattern: pattern,
		Path:    path,
		Ip:      ip,
		OS:      os,
		Browser: browser,
	}
	if uid != nil {
		log.Uid = uid.(uint32)
	}

	// 插入数据库
	if err := global.DB.Create(&log).Error; err != nil {
		global.Log.Error("登录日志插入失败！" + err.Error())
	}
}

// Star 星星操作日志
func (logService logService) Star(uid interface{}, sid interface{}, pattern uint8, flag bool, ip string) {
	log := entity.LogStar{
		Pattern: pattern,
		Flag:    flag,
		Ip:      ip,
	}
	if uid != nil {
		log.Uid = uid.(uint32)
	}
	if sid != nil {
		log.Sid = sid.(uint32)
	}

	// 插入数据库
	if err := global.DB.Create(&log).Error; err != nil {
		global.Log.Error("星星操作日志插入失败！" + err.Error())
	}
}

// File 文件操作日志
func (logService logService) File(uid uint32, pattern uint8, flag bool, remarks interface{}) {
	log := entity.LogFile{
		Uid:     uid,
		Pattern: pattern,
		Flag:    flag,
	}
	if remarks != nil {
		log.Remarks = remarks.(string)
	}

	// 插入数据库
	if err := global.DB.Create(&log).Error; err != nil {
		global.Log.Error("文件操作日志插入失败！" + err.Error())
	}
}

// Context 配置操作日志
func (logService logService) Context(context entity.Context) (err error) {
	log := entity.LogContext{
		Cid:     context.Id,
		Version: context.Version,
		Content: context.Content,
	}
	if err = global.DB.Create(&log).Error; err != nil {
		global.Log.Error("配置操作日志插入失败！" + err.Error())
	}

	return err
}
