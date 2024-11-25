package main

import (
    "log"
	"fmt"
    "github.com/gin-gonic/gin"
    "web-task/internal/config"
    "web-task/internal/api"
    "web-task/internal/middleware"
    "web-task/internal/repository"
    "web-task/internal/service"
    "web-task/internal/models"
)

func main() {
    // 初始化配置
    if err := config.Init(); err != nil {
        log.Fatalf("配置初始化失败: %v", err)
    }

    // 初始化数据库连接
    db, err := models.InitDB()
    if err != nil {
        log.Fatalf("数据库初始化失败: %v", err)
    }

    // 创建仓储工厂
    repoFactory := repository.NewRepositoryFactory(db)

    // 创建服务工厂
    baseService := service.NewService(repoFactory)
    serviceFactory := service.NewServiceFactory(baseService)

    // 创建 Gin 引擎
    r := gin.Default()

    // 添加中间件
    r.Use(middleware.Cors())
    r.Use(middleware.Logger())

    // 初始化路由，传入数据库连接
    api.RegisterRoutes(r, serviceFactory, db)

    // 启动服务器
    port := config.GlobalConfig.Server.Port
    log.Printf("服务器启动在 :%d", port)
    log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
} 