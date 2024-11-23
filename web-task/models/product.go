package models

import (
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Name        string    `gorm:"size:100;not null" json:"name"`
    Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
    Description string    `gorm:"type:text" json:"description"`
    Image       string    `json:"image"`
    Category    string    `gorm:"size:50;index" json:"category"`
    Stock       int       `gorm:"not null" json:"stock"`
    Sales       int       `gorm:"default:0" json:"sales"`
    Rating      float32   `gorm:"default:5" json:"rating"`
    Reviews     []Review  `json:"reviews,omitempty"`
}

type Review struct {
    gorm.Model
    UserID      uint      `gorm:"not null" json:"userId"`
    User        User      `json:"user"`
    ProductID   uint      `gorm:"not null" json:"productId"`
    Product     Product   `json:"-"`
    OrderID     uint      `gorm:"not null" json:"orderId"`
    Rating      int       `gorm:"not null" json:"rating"`
    Content     string    `gorm:"type:text" json:"content"`
    Images      []string  `gorm:"type:json" json:"images"`
} 