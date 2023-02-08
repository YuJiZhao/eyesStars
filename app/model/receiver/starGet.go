package receiver

import (
	"eyesStars/app/common/validate"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 15:52
 */

type StarGet struct {
	Id uint32 `uri:"id" binding:"gte=1"` // 星星id
}

func (starGet StarGet) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"Id.gte": "星星id必须大于等于1",
	}
}
