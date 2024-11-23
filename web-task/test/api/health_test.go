package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"os"
	"web-task/internal/config"

	"github.com/stretchr/testify/assert"
)

// 测试配置文件加载
func TestConfigLoading(t *testing.T) {
	// 切换到项目根目录
	err := os.Chdir("../..") // 根据你的项目目录结构调整路径
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	err = config.Init()
	assert.NoError(t, err, "配置文件应该成功加载")

	// 验证关键配置项
	assert.NotEmpty(t, config.GlobalConfig.Database.Host, "数据库主机配置不应为空")
	assert.NotEmpty(t, config.GlobalConfig.Database.Port, "数据库端口配置不应为空")
	assert.NotEmpty(t, config.GlobalConfig.Database.DBName, "数据库名称配置不应为空")
}

func TestHealthCheck(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Status string    `json:"status"`
			Time   time.Time `json:"time"`
		} `json:"data"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "success", response.Message)
	assert.Equal(t, "ok", response.Data.Status)
	assert.NotEmpty(t, response.Data.Time)
}

func TestDatabaseConnection(t *testing.T) {
	// 切换到项目根目录
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	// 确保配置已加载
	if err := config.Init(); err != nil {
		t.Fatalf("配置加载失败: %v", err)
	}

	// 在修改数据库名之前打印原始配置
	fmt.Printf("原始数据库配置:\n")
	fmt.Printf("Host: %s\n", config.GlobalConfig.Database.Host)
	fmt.Printf("DBName: %s\n", config.GlobalConfig.Database.DBName)
	fmt.Printf("Port: %d\n", config.GlobalConfig.Database.Port)

	db := setupTestDB()

	// 打印修改后的数据库名
	fmt.Printf("测试数据库名: %s\n", config.GlobalConfig.Database.DBName)

	// 测试数据库连接
	sqlDB, err := db.DB()
	assert.NoError(t, err)

	// 测试 ping
	err = sqlDB.Ping()
	assert.NoError(t, err)

	// 获取连接池统计信息
	stats := sqlDB.Stats()
	// 验证连接池是否正常工作
	assert.True(t, stats.MaxOpenConnections > 0)
	assert.True(t, stats.Idle >= 0)
	assert.True(t, stats.InUse >= 0)
}

func TestHealthCheckWithDB(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Status string    `json:"status"`
			Time   time.Time `json:"time"`
			DB     string    `json:"db"`
		} `json:"data"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "success", response.Message)
	assert.Equal(t, "ok", response.Data.Status)
	assert.Equal(t, "connected", response.Data.DB)
	assert.NotEmpty(t, response.Data.Time)
}
