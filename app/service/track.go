package service

import (
	"encoding/json"
	"eyesStars/app/model/dto"
	"eyesStars/global"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/2/1 13:56
 */

type trackService struct {
}

var TrackService = new(trackService)

// TrackVisit 进站埋点
func (trackService trackService) TrackVisit(uid interface{}, ip string, pkg string) {
	// 处理json包
	pkgStruct := dto.TrackVisit{}
	err := json.Unmarshal([]byte(pkg), &pkgStruct)
	if err != nil {
		var uidStr string
		if uid == nil {
			uidStr = "nil"
		} else {
			uidStr = strconv.Itoa(uid.(int))
		}
		global.Log.Error("JSON解析错误！" + err.Error() + "uid:" + uidStr + "IP:" + ip + ";package:" + pkg)
		return
	}

	// 执行日志操作
	LogService.Visit(
		uid,
		pkgStruct.Path,
		ip,
		pkgStruct.OS,
		pkgStruct.Browser,
	)
}

// TrackLogin 登录埋点
func (trackService trackService) TrackLogin(uid interface{}, ip string, pkg string) {
	// 处理json包
	pkgStruct := dto.TrackLogin{}
	err := json.Unmarshal([]byte(pkg), &pkgStruct)
	if err != nil {
		var uidStr string
		if uid == nil {
			uidStr = "nil"
		} else {
			uidStr = strconv.Itoa(uid.(int))
		}
		global.Log.Error("JSON解析错误！" + err.Error() + "uid:" + uidStr + "IP:" + ip + ";package:" + pkg)
		return
	}

	// 执行日志操作
	LogService.Login(
		uid,
		pkgStruct.Pattern,
		pkgStruct.Path,
		ip,
		pkgStruct.OS,
		pkgStruct.Browser,
	)
}
