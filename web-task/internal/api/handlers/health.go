package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web-task/pkg/utils/response"
	"gorm.io/gorm"
)

func HealthCheck(c *gin.Context) {
	// 获取数据库状态
	db := c.MustGet("db").(*gorm.DB)
	sqlDB, err := db.DB()
	dbStatus := "connected"
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "disconnected"
	}

	c.JSON(http.StatusOK, response.Success(gin.H{
		"status": "ok",
		"time":   time.Now(),
		"db":     dbStatus,
	}))
} 