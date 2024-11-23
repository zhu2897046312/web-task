package api

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "web-task/internal/api/handlers"
    "web-task/internal/middleware"
    "web-task/internal/service"
)

func RegisterRoutes(r *gin.Engine, sf *service.ServiceFactory, db *gorm.DB) {
    // 注入服务和数据库连接
    r.Use(middleware.InjectServices(sf, db))

    // 健康检查
    r.GET("/health", handlers.HealthCheck)

    // API v1 分组
    v1 := r.Group("/api/v1")
    {
        // 用户相关路由
        user := v1.Group("/users")
        {
            user.POST("/register", handlers.RegisterUser)
            user.POST("/login", handlers.LoginUser)
            user.GET("/profile", handlers.GetUserProfile)
            user.PUT("/profile", handlers.UpdateUserProfile)
            user.POST("/addresses", handlers.AddUserAddress)
            user.GET("/addresses", handlers.ListUserAddresses)
        }

        // 商品相关路由
        product := v1.Group("/products")
        {
            product.GET("", handlers.ListProducts)
            product.GET("/:id", handlers.GetProduct)
            product.POST("", handlers.CreateProduct)
            product.PUT("/:id", handlers.UpdateProduct)
            product.DELETE("/:id", handlers.DeleteProduct)
        }

        // 订单相关路由
        order := v1.Group("/orders")
        {
            order.POST("", handlers.CreateOrder)
            order.GET("", handlers.ListOrders)
            order.GET("/:id", handlers.GetOrder)
        }

        // 购物车相关路由
        cart := v1.Group("/cart")
        {
            cart.POST("/items", handlers.AddCartItem)
            cart.GET("/items", handlers.ListCartItems)
            cart.DELETE("/items/:id", handlers.RemoveCartItem)
            cart.PUT("/items/:id", handlers.UpdateCartItem)
        }

        // 评论相关路由
        review := v1.Group("/reviews")
        {
            review.POST("", handlers.CreateReview)
            review.GET("/products/:productId", handlers.GetProductReviews)
            review.GET("/users", handlers.GetUserReviews)
            review.DELETE("/:id", handlers.DeleteReview)
        }

        // 广告相关路由
        ad := v1.Group("/advertisements")
        {
            ad.POST("", handlers.CreateAdvertisement)
            ad.GET("", handlers.GetActiveAdvertisements)
            ad.PUT("/:id/status", handlers.UpdateAdvertisementStatus)
        }
    }
} 