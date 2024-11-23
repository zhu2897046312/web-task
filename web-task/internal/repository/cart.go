package repository

import (
    "web-task/internal/models"
    "gorm.io/gorm"
)

type CartRepository struct {
    *BaseRepository
}

func NewCartRepository(db *gorm.DB) *CartRepository {
    return &CartRepository{
        BaseRepository: NewBaseRepository(db),
    }
}

func (r *CartRepository) GetUserCart(userID uint) ([]models.CartItem, error) {
    var items []models.CartItem
    err := r.db.Where("user_id = ?", userID).
        Preload("Product").
        Find(&items).Error
    return items, err
}

func (r *CartRepository) GetCartItem(userID, productID uint) (*models.CartItem, error) {
    var item models.CartItem
    err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).
        First(&item).Error
    if err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *CartRepository) UpdateQuantity(itemID uint, quantity int) error {
    return r.db.Model(&models.CartItem{}).
        Where("id = ?", itemID).
        Update("quantity", quantity).Error
}

func (r *CartRepository) ClearCart(userID uint) error {
    return r.db.Where("user_id = ?", userID).
        Delete(&models.CartItem{}).Error
} 