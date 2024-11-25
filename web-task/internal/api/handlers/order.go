package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "web-task/internal/models"
    "web-task/internal/repository"
    "web-task/pkg/utils/response"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
        return
    }

    // TODO: 从上下文获取用户ID
    // userID := utils.GetUserIDFromContext(c)
    
    repo := c.MustGet("orderRepo").(*repository.OrderRepository)
    if err := repo.Create(&order); err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(order))
}

func ListOrders(c *gin.Context) {
    repo := c.MustGet("orderRepo").(*repository.OrderRepository)
    
    // TODO: 从上下文获取用户ID
    userID := uint(1) // 临时写死，实际应该从 JWT 中获取
    
    orders, err := repo.ListByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(orders))
}

func GetOrder(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, "Invalid order ID"))
        return
    }

    repo := c.MustGet("orderRepo").(*repository.OrderRepository)
    order, err := repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, response.Error(404, "Order not found"))
        return
    }

    c.JSON(http.StatusOK, response.Success(order))
} 