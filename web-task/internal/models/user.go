package models

import (
	"time"

	"gorm.io/gorm"
)

// User 定义了电商平台的用户模型
type User struct {
	ID            uint           `gorm:"primarykey;autoIncrement" json:"id"`             // 用户ID，主键
	Email         string         `gorm:"type:varchar(100);unique;not null" json:"email"` // 用户邮箱，唯一，不能为空
	Password      string         `gorm:"type:varchar(100);not null" json:"password"`     // 用户密码，不能为空，存储时加密，前端不显示
	Nickname      string         `gorm:"type:varchar(50)" json:"nickname"`               // 用户昵称，最大50个字符，可选
	Avatar        string         `gorm:"type:varchar(255)" json:"avatar"`                // 用户头像URL，最大255个字符，可选
	Role          string         `gorm:"type:varchar(20);default:'user'" json:"role"`    // 用户角色，默认为'user'，可为'user'或'admin'
	EmailVerified bool           `gorm:"default:false" json:"email_verified"`            // 邮箱是否验证，默认false，表示未验证
	VerifyToken   string         `gorm:"type:varchar(100)" json:"-"`                     // 用于邮箱验证的token，前端不显示
	TokenExpiry   *time.Time     `json:"-"`                                              // 验证token的过期时间，前端不显示
	CreatedAt     time.Time      `json:"created_at"`                                     // 用户账户创建时间
	UpdatedAt     time.Time      `json:"updated_at"`                                     // 用户信息最后更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                                 // 软删除字段，表示删除用户的时间，前端不显示
}

type Address struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`        // 地址的唯一标识符
	UserID    uint           `gorm:"not null;" json:"user_id"`                  // 关联的用户ID
	User      User           `gorm:"foreignKey:UserID;" json:"-"`               // 关联的用户对象 Address 中的 UserID 关联 User 中的 ID自增主键
	Name      string         `gorm:"type:varchar(50);not null" json:"name"`     // 收件人姓名
	Phone     string         `gorm:"type:varchar(20);not null" json:"phone"`    // 收件人电话
	Province  string         `gorm:"type:varchar(50);not null" json:"province"` // 省份
	City      string         `gorm:"type:varchar(50);not null" json:"city"`     // 城市
	District  string         `gorm:"type:varchar(50);not null" json:"district"` // 区/县
	Street    string         `gorm:"type:varchar(100);not null" json:"street"`  // 街道地址
	PostCode  string         `gorm:"type:varchar(10)" json:"post_code"`         // 邮政编码
	IsDefault bool           `gorm:"default:false" json:"is_default"`           // 是否为默认地址
	CreatedAt time.Time      `json:"created_at"`                                // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                                // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                            // 删除时间（软删除）
}
