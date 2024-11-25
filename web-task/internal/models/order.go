package models

import (
    "time"

    "gorm.io/gorm"
    "github.com/shopspring/decimal" // 添加这个导入
)

type Order struct {
	ID            uint            `gorm:"primarykey;autoIncrement" json:"id"`                   // 订单的唯一标识符
	UserID        uint            `gorm:"not null" json:"user_id"`                              // 关联的用户ID
	User          User            `gorm:"foreignKey:UserID" json:"-"`                           // 关联的用户对象
	OrderNumber   string          `gorm:"type:varchar(50);unique;not null" json:"order_number"` // 订单编号
	Status        string          `gorm:"type:varchar(20);not null" json:"status"`              // 订单状态：待处理/已支付/已发货/已完成/已取消
	TotalAmount   decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"total_amount"`      // 订单总金额
	AddressID     uint            `gorm:"not null" json:"address_id"`                           // 关联的地址ID
	Address       Address         `gorm:"foreignKey:AddressID" json:"address"`                  // 关联的地址对象
	PaymentMethod string          `gorm:"type:varchar(20)" json:"payment_method"`               // 支付方式
	PaymentStatus string          `gorm:"type:varchar(20)" json:"payment_status"`               // 支付状态：未支付/已支付/已退款
	PaymentTime   *time.Time      `json:"payment_time"`                                         // 支付时间
	CreatedAt     time.Time       `json:"created_at"`                                           // 创建时间
	UpdatedAt     time.Time       `json:"updated_at"`                                           // 更新时间
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"`                                       // 删除时间（软删除）
}

type OrderItem struct {
	ID        uint            `gorm:"primarykey" json:"id"`                     // 订单项的唯一标识符
	OrderID   uint            `gorm:"not null" json:"order_id"`                 // 关联的订单ID
	Order     Order           `gorm:"foreignKey:OrderID" json:"-"`              // 关联的订单对象
	ProductID uint            `gorm:"not null" json:"product_id"`               // 关联的产品ID
	Product   Product         `gorm:"foreignKey:ProductID" json:"product"`      // 关联的产品对象
	Quantity  int             `gorm:"not null" json:"quantity"`                 // 产品数量
	Price     decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"price"` // 产品单价
	CreatedAt time.Time       `json:"created_at"`                               // 创建时间
	UpdatedAt time.Time       `json:"updated_at"`                               // 更新时间
}


// Logistics 物流信息表
type Logistics struct {
	ID            uint            `gorm:"primarykey;autoIncrement" json:"id"`     // 物流信息的唯一标识符
	OrderID       uint            `gorm:"not null" json:"order_id"`               // 关联的订单ID
	Order         Order           `gorm:"foreignKey:OrderID" json:"-"`            // 关联的订单对象，JSON序列化时忽略
	TrackingNo    string          `gorm:"type:varchar(50)" json:"tracking_no"`    // 物流追踪号
	Carrier       string          `gorm:"type:varchar(50)" json:"carrier"`        // 承运商名称
	Status        string          `gorm:"type:varchar(20)" json:"status"`         // 物流状态
	ShippingFee   decimal.Decimal `gorm:"type:decimal(10,2)" json:"shipping_fee"` // 运费
	ShippedTime   *time.Time      `json:"shipped_time"`                           // 发货时间
	DeliveredTime *time.Time      `json:"delivered_time"`                         // 送达时间
	CreatedAt     time.Time       `json:"created_at"`                             // 创建时间
	UpdatedAt     time.Time       `json:"updated_at"`                             // 更新时间
}

// LogisticsTrace 物流跟踪表
type LogisticsTrace struct {
	ID          uint      `gorm:"primarykey" json:"id"`              // 物流跟踪记录的唯一标识符
	LogisticsID uint      `gorm:"not null" json:"logistics_id"`      // 关联的物流信息ID
	Logistics   Logistics `gorm:"foreignKey:LogisticsID" json:"-"`   // 关联的物流信息对象，JSON序列化时忽略
	Location    string    `gorm:"type:varchar(100)" json:"location"` // 当前位置
	Status      string    `gorm:"type:varchar(50)" json:"status"`    // 当前状态
	Description string    `gorm:"type:text" json:"description"`      // 状态描述
	TraceTime   time.Time `gorm:"not null" json:"trace_time"`        // 跟踪时间
	CreatedAt   time.Time `json:"created_at"`                        // 创建时间
}
