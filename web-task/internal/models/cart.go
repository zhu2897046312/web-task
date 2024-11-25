package models

import (
    "time"
)

type CartItem struct {
	ID        uint      `gorm:"primarykey;autoIncrement" json:"id"`  // 购物车项的唯一标识符
	UserID    uint      `gorm:"not null" json:"user_id"`             // 关联的用户ID
	User      User      `gorm:"foreignKey:UserID" json:"-"`          // 关联的用户对象，JSON序列化时忽略
	ProductID uint      `gorm:"not null" json:"product_id"`          // 关联的产品ID
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"` // 关联的产品对象
	Quantity  int       `gorm:"not null" json:"quantity"`            // 产品数量
	Selected  bool      `gorm:"default:true" json:"selected"`        // 是否选中，默认为true
	CreatedAt time.Time `json:"created_at"`                          // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                          // 更新时间
}