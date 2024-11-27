package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.Register(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	// 清除敏感信息
	user.Password = ""
	c.JSON(http.StatusOK, response.Success(user))
}

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	loginResp, err := svc.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error(401, err.Error()))
		return
	}

	// 清除敏感信息
	loginResp.User.Password = ""
	c.JSON(http.StatusOK, response.Success(loginResp))
}

// GetUserProfile 获取用户信息
func GetUserProfile(c *gin.Context) {
	// 从上下文中获取用户ID（由 AuthMiddleware 设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	user, err := svc.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(404, "User not found"))
		return
	}

	// 清除敏感信息
	user.Password = ""
	c.JSON(http.StatusOK, response.Success(user))
}

// UpdateUserProfile 更新用户信息
func UpdateUserProfile(c *gin.Context) {
	// 从上下文中获取用户ID（由 AuthMiddleware 设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	var updateData struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	// 创建用户对象，只包含要更新的字段
	user := &models.User{
		ID:       userID.(uint),
		Nickname: updateData.Nickname,
		Avatar:   updateData.Avatar,
	}

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	// 获取更新后的用户信息
	updatedUser, err := svc.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	// 清除敏感信息
	updatedUser.Password = ""
	c.JSON(http.StatusOK, response.Success(updatedUser))
}

// AddUserAddress 添加用户地址
func AddUserAddress(c *gin.Context) {
	// 从上下文中获取用户ID（由 AuthMiddleware 设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	// 确保只能为自己添加地址
	address.UserID = userID.(uint)

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.AddAddress(&address); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(address))
}

// ListUserAddresses 获取用户地址列表
func ListUserAddresses(c *gin.Context) {
	// 从上下文中获取用户ID（由 AuthMiddleware 设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error(401, "Unauthorized"))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	addresses, err := svc.ListAddresses(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(addresses))
}

// AdminListUsers 管理员获取用户列表
func AdminListUsers(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	// 获取分页参数
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	
	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)

	svc := c.MustGet("userService").(*service.UserService)
	users, total, err := svc.ListUsers(pageNum, pageSizeNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	// 清除敏感信息
	for i := range users {
		users[i].Password = ""
		users[i].VerifyToken = ""
	}

	c.JSON(http.StatusOK, response.Success(gin.H{
		"users": users,
		"total": total,
		"page": pageNum,
		"page_size": pageSizeNum,
	}))
}

// AdminUpdateUser 管理员更新用户信息
func AdminUpdateUser(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid user ID"))
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	user.ID = uint(userID)

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.UpdateUserByAdmin(&user); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, response.Success(user))
}

// AdminDeleteUser 管理员删除用户
func AdminDeleteUser(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid user ID"))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.DeleteUser(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
} 