package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username    string     `gorm:"size:50;not null;uniqueIndex" json:"username"`
    Password    string     `gorm:"size:100;not null" json:"-"`
    Email       string     `gorm:"size:100;not null;uniqueIndex" json:"email"`
    Phone       string     `gorm:"size:20" json:"phone"`
    Avatar      string     `json:"avatar"`
    Addresses   []Address  `json:"addresses,omitempty"`
    Orders      []Order    `json:"orders,omitempty"`
}

type Address struct {
    gorm.Model
    UserID      uint   `gorm:"not null" json:"userId"`
    Receiver    string `gorm:"size:50;not null" json:"receiver"`
    Phone       string `gorm:"size:20;not null" json:"phone"`
    Province    string `gorm:"size:50;not null" json:"province"`
    City        string `gorm:"size:50;not null" json:"city"`
    District    string `gorm:"size:50;not null" json:"district"`
    Detail      string `gorm:"size:200;not null" json:"detail"`
    IsDefault   bool   `gorm:"default:false" json:"isDefault"`
} 