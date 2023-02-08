package validate

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 9:05
 */

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// GetErrorMsg request结构体需要实现Validator接口(为了避免性能损失，这里不做校验)
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		message := ""
		for _, v := range err.(validator.ValidationErrors) {
			if msg, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
				message += msg + ";"
			} else {
				message += v.Error() + ";"
			}
		}
		return strings.Trim(message, ";")
	}

	return "参数错误"
}
