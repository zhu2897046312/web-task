package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"web-task/internal/service"
)

func InjectServices(sf *service.ServiceFactory, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Set("userService", sf.GetUserService())
		c.Set("productService", sf.GetProductService())
		c.Set("orderService", sf.GetOrderService())
		c.Set("cartService", sf.GetCartService())
		c.Set("reviewService", sf.GetReviewService())
		c.Set("advertisementService", sf.GetAdvertisementService())
		c.Next()
	}
} 