package models

import (
    "gorm.io/gorm"
)

type Order struct {
    gorm.Model
    UserID        uint          `gorm:"not null" json:"userId"`
    User          User          `json:"user"`
    Items         []OrderItem   `json:"items"`
    TotalAmount   float64       `gorm:"type:decimal(10,2);not null" json:"totalAmount"`
    Status        string        `gorm:"size:20;not null;default:'pending'" json:"status"`
    PaymentMethod string        `gorm:"size:20" json:"paymentMethod"`
    AddressID     uint          `gorm:"not null" json:"addressId"`
    Address       Address       `json:"address"`
    Logistics     *Logistics    `json:"logistics,omitempty"`
}

type OrderItem struct {
    gorm.Model
    OrderID     uint    `gorm:"not null" json:"orderId"`
    ProductID   uint    `gorm:"not null" json:"productId"`
    Product     Product `json:"product"`
    Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
    Quantity    int     `gorm:"not null" json:"quantity"`
    Subtotal    float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
}

type Logistics struct {
    gorm.Model
    OrderID         uint             `gorm:"not null;uniqueIndex" json:"orderId"`
    Order           Order            `json:"-"`
    TrackingNumber  string          `gorm:"size:50" json:"trackingNumber"`
    Carrier         string          `gorm:"size:50" json:"carrier"`
    Status          string          `gorm:"size:20;not null;default:'pending'" json:"status"`
    Traces          []LogisticsTrace `json:"traces"`
}

type LogisticsTrace struct {
    gorm.Model
    LogisticsID uint   `gorm:"not null" json:"logisticsId"`
    Time        string `gorm:"not null" json:"time"`
    Location    string `gorm:"size:100" json:"location"`
    Description string `gorm:"type:text" json:"description"`
} 