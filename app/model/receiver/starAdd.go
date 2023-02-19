package receiver

import (
	"eyesStars/app/common/validate"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 14:53
 */

type StarAdd struct {
	Content   string `json:"content" binding:"required,max=300"` // "提交内容"
	Anonymous bool   `json:"anonymous" default:"false"`          // "是否匿名"
	Name      string `json:"name" binding:"max=10"`              // "化名"
}

func (starAdd StarAdd) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"content.required": "提交内容不能为空",
		"content.max":      "提交内容长度不能大于300",
		"name.max":         "化名长度不能大于10",
	}
}
