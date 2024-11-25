package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

func CreateAdvertisement(c *gin.Context) {
	var ad models.Advertisement
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.CreateAd(&ad); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(ad))
}

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

func UpdateAdvertisementStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid advertisement ID"))
		return
	}

	var req struct {
		Status bool `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	svc := c.MustGet("advertisementService").(*service.AdvertisementService)
	if err := svc.UpdateAdStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
} 