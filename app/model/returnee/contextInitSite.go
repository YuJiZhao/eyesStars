package returnee

/**
 * @author eyesYeager
 * @date 2023/2/1 16:06
 */

type ContextInitSite struct {
	AppName          string // 网站名称
	AppNameEN        string // 网站名称(英文)
	AppVersion       string // 网站版本
	AppTitle         string // 网站标题
	SiteUrl          string // 网站地址
	MetaDesc         string // meta数据：description
	MetaKeyWord      string // meta数据：keyword
	Favicon          string // 网站图标地址
	Slogan           string // 页角标语
	SloganFontFamily string // 页脚标语字体名称
	SloganFontUrl    string // 页脚标语字体文件地址
	AesKey           string // aes密钥
}
