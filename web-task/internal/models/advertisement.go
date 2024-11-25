package models

import (
	"time"

	"gorm.io/gorm"
)

type Advertisement struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`              // 广告的唯一标识符
	Title     string         `gorm:"type:varchar(100);not null" json:"title"`         // 广告标题
	Image     string         `gorm:"type:varchar(255);not null" json:"image"`         // 广告图片URL
	URL       string         `gorm:"type:varchar(255)" json:"url"`                    // 广告链接URL
	Position  string         `gorm:"type:varchar(50)" json:"position"`                // 广告位置
	StartTime time.Time      `json:"start_time"`                                      // 广告开始时间
	EndTime   time.Time      `json:"end_time"`                                        // 广告结束时间
	Status    string         `gorm:"type:varchar(20);default:'active'" json:"status"` // 广告状态，默认为active
	CreatedAt time.Time      `json:"created_at"`                                      // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                                      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                  // 删除时间，软删除
}