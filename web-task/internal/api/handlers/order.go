package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	var req struct {
		AddressID uint              `json:"address_id" binding:"required"`
		Items     []models.OrderItem `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("orderService").(*service.OrderService)
	order, err := svc.CreateOrder(userID.(uint), req.Items, req.AddressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(order))
}

// GetOrder 获取订单详情
func GetOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid order ID"))
		return
	}

	svc := c.MustGet("orderService").(*service.OrderService)
	order, err := svc.GetOrder(uint(orderID), userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(404, "Order not found"))
		return
	}

	c.JSON(http.StatusOK, response.Success(order))
}

// ListOrders 获取用户订单列表
func ListOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	svc := c.MustGet("orderService").(*service.OrderService)
	orders, total, err := svc.ListUserOrders(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(gin.H{
		"orders":    orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}))
}

// AdminListOrders 管理员查看订单列表
func AdminListOrders(c *gin.Context) {
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	svc := c.MustGet("orderService").(*service.OrderService)
	orders, total, err := svc.ListOrdersByStatus(status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(gin.H{
		"orders":    orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}))
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(c *gin.Context) {
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid order ID"))
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("orderService").(*service.OrderService)
	if err := svc.UpdateOrderStatus(uint(orderID), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
}

// UpdateLogistics 更新物流信息
func UpdateLogistics(c *gin.Context) {
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	var logistics models.Logistics
	if err := c.ShouldBindJSON(&logistics); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("orderService").(*service.OrderService)
	if err := svc.UpdateLogistics(&logistics); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(logistics))
}

// GetLogistics 获取订单物流信息
func GetLogistics(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid order ID"))
		return
	}

	svc := c.MustGet("orderService").(*service.OrderService)
	logistics, err := svc.GetLogistics(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(404, "Logistics not found"))
		return
	}

	c.JSON(http.StatusOK, response.Success(logistics))
} 