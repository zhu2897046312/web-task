package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
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

func LoginUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	svc := c.MustGet("userService").(*service.UserService)
	user, err := svc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error(401, err.Error()))
		return
	}

	// TODO: 生成 JWT token
	token := "dummy-token" // 这里需要实现真实的 JWT token 生成

	c.JSON(http.StatusOK, response.Success(gin.H{
		"token": token,
		"user":  user,
	}))
}

func GetUserProfile(c *gin.Context) {
	// TODO: 从 JWT 中获取用户 ID
	userID := uint(1) // 临时写死，实际应该从 JWT 中获取

	svc := c.MustGet("userService").(*service.UserService)
	user, err := svc.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(404, "User not found"))
		return
	}

	// 清除敏感信息
	user.Password = ""
	c.JSON(http.StatusOK, response.Success(user))
}

func UpdateUserProfile(c *gin.Context) {
	// TODO: 从 JWT 中获取用户 ID
	userID := uint(1)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}
	user.ID = userID

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	// 清除敏感信息
	user.Password = ""
	c.JSON(http.StatusOK, response.Success(user))
}

func AddUserAddress(c *gin.Context) {
	// TODO: 从 JWT 中获取用户 ID
	userID := uint(1)

	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}
	address.UserID = userID

	svc := c.MustGet("userService").(*service.UserService)
	if err := svc.AddAddress(&address); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(address))
}

func ListUserAddresses(c *gin.Context) {
	// TODO: 从 JWT 中获取用户 ID
	userID := uint(1)

	svc := c.MustGet("userService").(*service.UserService)
	addresses, err := svc.ListAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(addresses))
} 