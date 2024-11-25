package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-task/internal/models"
	"web-task/internal/service"
	"web-task/pkg/utils/response"
)

func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
		return
	}

	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)
	review.UserID = userID

	svc := c.MustGet("reviewService").(*service.ReviewService)
	if err := svc.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(review))
}

func GetProductReviews(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid product ID"))
		return
	}

	svc := c.MustGet("reviewService").(*service.ReviewService)
	reviews, err := svc.GetProductReviews(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(reviews))
}

func GetUserReviews(c *gin.Context) {
	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("reviewService").(*service.ReviewService)
	reviews, err := svc.GetUserReviews(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(reviews))
}

func DeleteReview(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Invalid review ID"))
		return
	}

	// TODO: 从 JWT 获取用户 ID
	userID := uint(1)

	svc := c.MustGet("reviewService").(*service.ReviewService)
	if err := svc.DeleteReview(uint(reviewID), userID); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(nil))
} 