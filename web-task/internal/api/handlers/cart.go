package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

func AddCartItem(c *gin.Context) {
	var req struct {
		ProductID uint `json:"productId" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("cartService").(*service.CartService)
	if err := svc.AddItem(userID, req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
}

func ListCartItems(c *gin.Context) {
	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("cartService").(*service.CartService)
	items, err := svc.ListItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(items))
}

func UpdateCartItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid item ID"))
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("cartService").(*service.CartService)
	if err := svc.UpdateQuantity(userID, uint(itemID), req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
}

func RemoveCartItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid item ID"))
		return
	}

	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("cartService").(*service.CartService)
	if err := svc.RemoveItem(userID, uint(itemID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
} 