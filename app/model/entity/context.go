package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/31 14:00
 */

// Context 配置信息
type Context struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null;unique"`
	Remarks    string    `json:"remarks" gorm:"comment:备注"`
	Version    uint32    `json:"version" gorm:"not null;default:1;comment:当前最新版本"`
	Content    string    `json:"content" gorm:"not null;size:350;comment:当前最新内容"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime:milli"`
}
