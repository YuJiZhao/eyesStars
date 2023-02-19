package entity

import "time"

/**
 * @author eyesYeager
 * @date 2023/1/29 11:59
 */

type Star struct {
	Id         uint32    `json:"id" gorm:"primaryKey"`
	Uid        uint32    `json:"uid" gorm:"not null;index"`
	Name       string    `json:"name" gorm:"size:10"`
	Anonymous  bool      `json:"anonymous" gorm:"default:false;comment:是否匿名"`
	Content    string    `json:"content" gorm:"not null;size:300"`
	EmailNeed  bool      `json:"email_need" gorm:"default:false;comment:是否需要邮件通知"`
	EmailFlag  bool      `json:"email_flag" gorm:"default:false;comment:是否已经邮件通知"`
	Status     uint8     `json:"status" gorm:"not null;default:0;comment:0->正常,1->冻结,2->删除"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	DeleteTime time.Time `json:"delete_time"`
}
