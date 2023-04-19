package main

import (
	"eyesStars/bootstrap"
	"eyesStars/global"
)

/**
 * @author eyesYeager
 * @date 2023/1/19 13:45
 */

// @title 耶瞳星空
// @version 1.0
// @description 耶瞳星空，来自星星的寄语
// @contact.name   耶瞳
// @contact.url    http://space.eyesspace.top
// @contact.email  eyesyeager@163.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	bootstrap.InitNacos()
	bootstrap.InitializeValidator()
	global.Log = bootstrap.InitializeLog()
	global.DB = bootstrap.InitializeDB()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			_ = db.Close()
		}
	}()

	// 启动程序
	bootstrap.RunServer()
}
