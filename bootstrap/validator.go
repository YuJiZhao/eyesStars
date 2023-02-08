package bootstrap

import (
	"eyesStars/app/common/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 10:09
 */

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", validate.MobileValidate)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
