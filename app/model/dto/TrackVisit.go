package dto

/**
 * @author eyesYeager
 * @date 2023/2/1 13:31
 */

type TrackVisit struct {
	Path    string `json:"path"`    // 进站路径
	OS      string `json:"os"`      // 客户端操作系统
	Browser string `json:"Browser"` // 客户端浏览器
}
