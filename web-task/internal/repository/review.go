package repository

import (
	"web-task/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	*BaseRepository
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *ReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Preload("User").First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) ListByProduct(productID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("product_id = ?", productID).
		Preload("User").
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) ListByUser(userID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("user_id = ?", userID).
		Preload("Product").
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) UpdateProductRating(productID uint) error {
	// 计算产品的平均评分并更新
	var avgRating float32
	err := r.db.Model(&models.Review{}).
		Select("COALESCE(AVG(rating), 5)").
		Where("product_id = ?", productID).
		Scan(&avgRating).Error
	if err != nil {
		return err
	}

	return r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Update("rating", avgRating).Error
} 