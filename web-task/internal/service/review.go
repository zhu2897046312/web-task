package service

import (
	"errors"
	"web-task/internal/models"
)

type ReviewService struct {
	*Service
}

func NewReviewService(base *Service) *ReviewService {
	return &ReviewService{Service: base}
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	// 检查订单是否存在
	order, err := s.repoFactory.GetOrderRepository().GetByID(review.OrderID)
	if err != nil {
		return errors.New("order not found")
	}

	// 检查订单是否属于该用户
	if order.UserID != review.UserID {
		return errors.New("order does not belong to this user")
	}

	// 检查评分范围
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// 创建评论
	if err := s.repoFactory.GetReviewRepository().Create(review); err != nil {
		return err
	}

	// 更新产品评分
	return s.repoFactory.GetReviewRepository().UpdateProductRating(review.ProductID)
}

func (s *ReviewService) GetProductReviews(productID uint) ([]models.Review, error) {
	return s.repoFactory.GetReviewRepository().ListByProduct(productID)
}

func (s *ReviewService) GetUserReviews(userID uint) ([]models.Review, error) {
	return s.repoFactory.GetReviewRepository().ListByUser(userID)
}

func (s *ReviewService) DeleteReview(reviewID, userID uint) error {
	// 检查评论是否存在
	review, err := s.repoFactory.GetReviewRepository().GetByID(reviewID)
	if err != nil {
		return errors.New("review not found")
	}

	// 检查评论是否属于该用户
	if review.UserID != userID {
		return errors.New("review does not belong to this user")
	}

	// 删除评论
	if err := s.repoFactory.GetReviewRepository().Delete(review); err != nil {
		return err
	}

	// 更新产品评分
	return s.repoFactory.GetReviewRepository().UpdateProductRating(review.ProductID)
} 