package repository

import (
	"web-task/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	*BaseRepository
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").Preload("Address").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) ListByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).
		Preload("Items").
		Preload("Address").
		Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}

func (r *OrderRepository) GetOrderItems(orderID uint) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := r.db.Where("order_id = ?", orderID).
		Preload("Product").
		Find(&items).Error
	return items, err
} 