package models

import (
	"time"

	"github.com/shopspring/decimal" // 添加这个导入
	"gorm.io/gorm"
)

type Product struct {
	ID          uint            `gorm:"primarykey;autoIncrement" json:"id"`              // 产品的唯一标识符
	Name        string          `gorm:"type:varchar(100);not null" json:"name"`          // 产品名称
	Description string          `gorm:"type:text" json:"description"`                    // 产品描述
	Price       decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"price"`        // 产品价格
	Stock       int             `gorm:"not null" json:"stock"`                           // 库存数量
	Status      string          `gorm:"type:varchar(20);default:'active'" json:"status"` // 产品状态：活跃/不活跃
	Images      []string        `gorm:"type:json" json:"images"`                         // 产品图片列表
	Category    string          `gorm:"type:varchar(50)" json:"category"`                // 产品类别
	Tags        []string        `gorm:"type:json" json:"tags"`                           // 产品标签
	CreatedAt   time.Time       `json:"created_at"`                                      // 创建时间
	UpdatedAt   time.Time       `json:"updated_at"`                                      // 更新时间
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`                                  // 删除时间（软删除）
}


// Review 商品评价表
type Review struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`  // 评价的唯一标识符
	UserID    uint           `gorm:"not null" json:"user_id"`             // 关联的用户ID
	User      User           `gorm:"foreignKey:UserID" json:"user"`       // 关联的用户对象
	ProductID uint           `gorm:"not null" json:"product_id"`          // 关联的产品ID
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"` // 关联的产品对象
	OrderID   uint           `gorm:"not null" json:"order_id"`            // 关联的订单ID
	Order     Order          `gorm:"foreignKey:OrderID" json:"-"`         // 关联的订单对象，JSON序列化时忽略
	Rating    int            `gorm:"not null" json:"rating"`              // 评分，1-5星
	Content   string         `gorm:"type:text" json:"content"`            // 评价内容
	Images    []string       `gorm:"type:json" json:"images"`             // 评价图片，JSON格式
	CreatedAt time.Time      `json:"created_at"`                          // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                      // 删除时间，软删除
}
