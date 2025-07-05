package bootstrap

import (
	"eyesStars/global"
	"fmt"
	"github.com/spf13/viper"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 16:54
 */

func InitializeConfig() {
	// 初始化 viper
	v := viper.New()

	// 读取根配置文件
	v.SetConfigFile("config.json")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 根配置文件赋值
	if err := v.Unmarshal(&global.Config); err != nil {
		global.Log.Error(err.Error())
		fmt.Println(err)
	}
}
