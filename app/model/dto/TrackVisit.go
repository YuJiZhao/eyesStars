package dto

/**
 * @author eyesYeager
 * @date 2023/2/1 13:31
 */

type TrackVisit struct {
	path    string // 进站路径
	os      string // 客户端操作系统
	browser string // 客户端浏览器
}

func (trackVisit TrackVisit) GetPath() string {
	return trackVisit.path
}

func (trackVisit TrackVisit) GetOS() string {
	return trackVisit.os
}

func (trackVisit TrackVisit) GetBrowser() string {
	return trackVisit.browser
}
