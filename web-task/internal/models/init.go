package models

import (
	"fmt"
	"log"
	"time"
	"web-task/internal/config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func InitDB() (*gorm.DB, error) {
	dsn := config.GlobalConfig.Database.DSN()
	log.Printf("Connecting to database with DSN: %s", dsn)

	// 设置连接重试
	var db *gorm.DB
	var err error
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(time.Second * 2) // 等待2秒后重试
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database after %d attempts: %v", maxRetries, err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	// 自动迁移数据库表结构
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("auto migration failed: %v", err)
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	log.Println("Starting database migration...")
	err := db.AutoMigrate(
		&User{},
		&Address{},
		&Product{},
		&Order{},
		&OrderItem{},
		&CartItem{},
		&Review{},
		&Advertisement{},
		&Logistics{},
		&LogisticsTrace{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")
	return nil
} 