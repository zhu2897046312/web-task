package models

import (
    "gorm.io/gorm"
    "time"
)

type Advertisement struct {
    gorm.Model
    Title     string    `gorm:"size:100;not null" json:"title"`
    Image     string    `gorm:"not null" json:"image"`
    Link      string    `gorm:"not null" json:"link"`
    Position  string    `gorm:"size:50;not null" json:"position"`
    StartTime time.Time `gorm:"not null" json:"startTime"`
    EndTime   time.Time `gorm:"not null" json:"endTime"`
    Status    bool      `gorm:"default:true" json:"status"`
} 