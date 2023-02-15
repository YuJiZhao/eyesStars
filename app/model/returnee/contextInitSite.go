package returnee

/**
 * @author eyesYeager
 * @date 2023/2/1 16:06
 */

type ContextInitSite struct {
	AppName          string `json:"appName"`          // 网站名称
	AppNameEN        string `json:"appNameEN"`        // 网站名称(英文)
	AppVersion       string `json:"appVersion"`       // 网站版本
	AppTitle         string `json:"appTitle"`         // 网站标题
	SiteUrl          string `json:"siteUrl"`          // 网站地址
	MetaDesc         string `json:"metaDesc"`         // meta数据：description
	MetaKeyWord      string `json:"metaKeyWord"`      // meta数据：keyword
	Favicon          string `json:"favicon"`          // 网站图标地址
	Slogan           string `json:"slogan"`           // 页角标语
	SloganFontFamily string `json:"sloganFontFamily"` // 页脚标语字体名称
	SloganFontUrl    string `json:"sloganFontUrl"`    // 页脚标语字体文件地址
	AesKey           string `json:"sesKey"`           // aes密钥
	MoonUrl          string `json:"moonUrl"`          // 月亮图片
}
