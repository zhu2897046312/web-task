package models

import (
    "gorm.io/gorm"
)

type CartItem struct {
    gorm.Model
    UserID    uint    `gorm:"not null" json:"userId"`
    User      User    `json:"-"`
    ProductID uint    `gorm:"not null" json:"productId"`
    Product   Product `json:"product"`
    Quantity  int     `gorm:"not null" json:"quantity"`
} 