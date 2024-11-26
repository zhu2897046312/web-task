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

            // 公开的广告接口
            public.GET("/advertisements", handlers.GetActiveAdvertisements)
            public.GET("/advertisements/position/:position", handlers.GetActiveAdvertisements)
        }

        // 需要认证的路由
        authorized := v1.Group("")
        authorized.Use(middleware.AuthMiddleware())
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
                orders.GET("/:id/logistics", handlers.GetLogistics)
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

            // 管理员路由组
            admin := authorized.Group("/admin")
            {
                // 用户管理
                admin.GET("/users", handlers.AdminListUsers)
                admin.PUT("/users/:id", handlers.AdminUpdateUser)
                admin.DELETE("/users/:id", handlers.AdminDeleteUser)

                // 商品管理
                admin.POST("/products", handlers.CreateProduct)
                admin.PUT("/products/:id", handlers.UpdateProduct)
                admin.DELETE("/products/:id", handlers.DeleteProduct)

                // 订单管理
                admin.GET("/orders", handlers.AdminListOrders)
                admin.PUT("/orders/:id/status", handlers.UpdateOrderStatus)
                admin.POST("/orders/:id/logistics", handlers.UpdateLogistics)

                // 广告管理
                admin.POST("/advertisements", handlers.CreateAdvertisement)
                admin.GET("/advertisements", handlers.ListAdvertisements)
                admin.GET("/advertisements/:id", handlers.GetAdvertisement)
                admin.PUT("/advertisements/:id", handlers.UpdateAdvertisement)
                admin.DELETE("/advertisements/:id", handlers.DeleteAdvertisement)
                admin.PUT("/advertisements/:id/status", handlers.UpdateAdvertisementStatus)
            }
        }
    }
} 