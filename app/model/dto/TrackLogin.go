package dto

/**
 * @author eyesYeager
 * @date 2023/2/1 15:28
 */

type TrackLogin struct {
	Pattern uint8  `json:"pattern"` // 操作类型(0:前往登录，1:前往信息页)
	Path    string `json:"path"`    // 操作时所处页面路径
	OS      string `json:"os"`      // 客户端操作系统
	Browser string `json:"browser"` // 客户端浏览器
}
