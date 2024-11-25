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
        // 公开路由 - 不需要认证
        public := v1.Group("")
        {
            // 用户认证
            public.POST("/users/register", handlers.RegisterUser)
            public.POST("/users/login", handlers.LoginUser)

            // 公开的商品接口
            public.GET("/products", handlers.ListProducts)
            public.GET("/products/:id", handlers.GetProduct)
            public.GET("/products/:id/reviews", handlers.GetProductReviews)

            // 广告
            public.GET("/advertisements", handlers.GetActiveAdvertisements)
        }

        // 需要认证的路由
        authorized := v1.Group("")
        authorized.Use(middleware.AuthMiddleware()) // 添加认证中间件
        {
            // 用户相关
            users := authorized.Group("/users")
            {
                users.GET("/profile", handlers.GetUserProfile)
                users.PUT("/profile", handlers.UpdateUserProfile)
                users.POST("/addresses", handlers.AddUserAddress)
                users.GET("/addresses", handlers.ListUserAddresses)
            }

            // 订单相关
            orders := authorized.Group("/orders")
            {
                orders.POST("", handlers.CreateOrder)
                orders.GET("", handlers.ListOrders)
                orders.GET("/:id", handlers.GetOrder)
            }

            // 购物车相关
            cart := authorized.Group("/cart")
            {
                cart.POST("/items", handlers.AddCartItem)
                cart.GET("/items", handlers.ListCartItems)
                cart.DELETE("/items/:id", handlers.RemoveCartItem)
                cart.PUT("/items/:id", handlers.UpdateCartItem)
            }

            // 评论相关
            reviews := authorized.Group("/reviews")
            {
                reviews.POST("", handlers.CreateReview)
                reviews.GET("/me", handlers.GetUserReviews)
                reviews.DELETE("/:id", handlers.DeleteReview)
            }

            // 需要认证的商品操作（如管理员操作）
            products := authorized.Group("/products")
            {
                products.POST("", handlers.CreateProduct)
                products.PUT("/:id", handlers.UpdateProduct)
                products.DELETE("/:id", handlers.DeleteProduct)
            }

            // 广告管理
            ads := authorized.Group("/advertisements")
            {
                ads.POST("", handlers.CreateAdvertisement)
                ads.PUT("/:id/status", handlers.UpdateAdvertisementStatus)
            }
        }
    }
} 