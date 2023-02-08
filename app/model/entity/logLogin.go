package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/29 11:59
 */

type LogLogin struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Uid        uint32    `json:"uid" gorm:"not null;index"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:0->前往登录,1->退出登录"`
	Path       string    `json:"path"`
	Ip         string    `json:"ip"`
	OS         string    `json:"os"`
	Browser    string    `json:"browser"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
