package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 10:08
 */

// MobileValidate 校验手机号
func MobileValidate(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	return ok
}
