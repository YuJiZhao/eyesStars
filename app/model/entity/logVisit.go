package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/29 12:00
 */

type LogVisit struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Uid        uint32    `json:"uid"`
	Path       string    `json:"path"`
	Ip         string    `json:"ip"`
	OS         string    `json:"os"`
	Browser    string    `json:"browser"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
