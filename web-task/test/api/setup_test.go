package api_test

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"web-task/internal/api"
	"web-task/internal/config"
	"web-task/internal/models"
	"web-task/internal/repository"
	"web-task/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// 获取项目根目录
func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func setupTestDB() *gorm.DB {
	if testDB != nil {
		return testDB
	}

	// 切换到项目根目录（go.mod所在目录）
	projectRoot := getProjectRoot()
	if err := os.Chdir(projectRoot); err != nil {
		log.Fatalf("切换到项目根目录失败: %v", err)
	}

	// 加载测试配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 初始化测试数据库
	db, err := models.InitDB()
	if err != nil {
		log.Printf("数据库配置: %+v", config.GlobalConfig.Database)
		log.Fatalf("测试数据库初始化失败: %v", err)
	}

	// 验证数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	testDB = db
	return testDB
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// 创建测试用的服务工厂
	db := setupTestDB()
	repoFactory := repository.NewRepositoryFactory(db)
	baseService := service.NewService(repoFactory)
	serviceFactory := service.NewServiceFactory(baseService)

	// 传入数据库连接
	api.RegisterRoutes(r, serviceFactory, db)
	return r
}
