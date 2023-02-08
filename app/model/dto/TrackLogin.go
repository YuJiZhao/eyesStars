package dto

/**
 * @author eyesYeager
 * @date 2023/2/1 15:28
 */

type TrackLogin struct {
	pattern uint8  // 操作类型(0:前往登录，1:退出登录)
	path    string // 操作时所处页面路径
	os      string // 客户端操作系统
	browser string // 客户端浏览器
}

func (trackLogin TrackLogin) GetPattern() uint8 {
	return trackLogin.pattern
}

func (trackLogin TrackLogin) GetPath() string {
	return trackLogin.path
}

func (trackLogin TrackLogin) GetOS() string {
	return trackLogin.os
}

func (trackLogin TrackLogin) GetBrowser() string {
	return trackLogin.browser
}
