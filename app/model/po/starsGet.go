package po

import (
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/1/30 14:43
 */

type StarsGet struct {
	Id         uint32
	Uid        uint64
	Content    string
	Anonymous  bool
	Name       string
	CreateTime time.Time
}
