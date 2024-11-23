package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(localhost:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库结构
	err = db.AutoMigrate(
		&User{},
		&Address{},
		&Product{},
		&Review{},
		&Order{},
		&OrderItem{},
		&Logistics{},
		&LogisticsTrace{},
		&CartItem{},
		&Advertisement{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
} 