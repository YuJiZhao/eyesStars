package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/29 12:00
 */

type LogFile struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Uid        uint32    `json:"uid" gorm:"not null"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:操作类型0->上传头像"`
	Flag       bool      `json:"flag" gorm:"not null;comment:是否操作成功"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:备注"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
