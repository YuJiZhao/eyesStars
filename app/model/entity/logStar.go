package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/30 10:50
 */

type LogStar struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Uid        uint32    `json:"uid"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:0->创建,1->查询"`
	Sid        uint32    `json:"sid" gorm:"not null"`
	Flag       bool      `json:"flag" gorm:"not null"`
	Ip         string    `json:"ip"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
