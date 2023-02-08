package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/29 11:59
 */

// LogContext 历史配置信息
type LogContext struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Cid        uint32    `json:"cid" gorm:"not null"`
	Version    uint32    `json:"version" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null;type:TEXT"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
