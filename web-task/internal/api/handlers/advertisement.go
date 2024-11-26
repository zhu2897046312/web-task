package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

// CreateAdvertisement 创建广告
func CreateAdvertisement(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	var ad models.Advertisement
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.CreateAd(&ad); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(ad))
}

// GetAdvertisement 获取广告详情
func GetAdvertisement(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid advertisement ID"))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	ad, err := svc.GetAd(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(404, "Advertisement not found"))
		return
	}

	c.JSON(http.StatusOK, response.Success(ad))
}

// UpdateAdvertisement 更新广告信息
func UpdateAdvertisement(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid advertisement ID"))
		return
	}

	var ad models.Advertisement
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	ad.ID = uint(id)

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.UpdateAd(&ad); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(ad))
}

// DeleteAdvertisement 删除广告
func DeleteAdvertisement(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid advertisement ID"))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.DeleteAd(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
}

// ListAdvertisements 获取广告列表
func ListAdvertisements(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	
	var ads []models.Advertisement
	var total int64
	var err error

	if status != "" {
		ads, total, err = svc.ListAdsByStatus(status, page, pageSize)
	} else {
		ads, total, err = svc.ListAllAds(page, pageSize)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(gin.H{
		"advertisements": ads,
		"total":         total,
		"page":          page,
		"page_size":     pageSize,
	}))
}

// GetActiveAdvertisements 获取有效广告
func GetActiveAdvertisements(c *gin.Context) {
	position := c.Query("position")
	svc := c.MustGet("advertisementService").(*service.AdvertisementService)

	var ads []models.Advertisement
	var err error

	if position != "" {
		ads, err = svc.GetActiveAds(position)
	} else {
		ads, err = svc.ListActiveAds()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(ads))
}

// UpdateAdvertisementStatus 更新广告状态
func UpdateAdvertisementStatus(c *gin.Context) {
	// 检查是否是管理员
	role, exists := c.Get("userRole")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, response.Error(403, "Permission denied"))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid advertisement ID"))
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid request parameters"))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.UpdateAdStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
} 