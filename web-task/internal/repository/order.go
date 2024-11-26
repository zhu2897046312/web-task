package repository

import (
	"web-task/internal/models"
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	*BaseRepository
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// Create 创建订单
func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

// GetByID 获取订单详情
func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items.Product").
		Preload("Address").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, nickname, email") // 只选择需要的用户字段
		}).
		First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// ListByUserID 获取用户订单列表
func (r *OrderRepository) ListByUserID(userID uint, page, pageSize int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64
	
	offset := (page - 1) * pageSize

	// 获取总数
	if err := r.db.Model(&models.Order{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取订单列表
	err := r.db.Where("user_id = ?", userID).
		Preload("Items.Product").
		Preload("Address").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, total, err
}

// UpdateStatus 更新订单状态
func (r *OrderRepository) UpdateStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}

// UpdatePaymentStatus 更新支付状态
func (r *OrderRepository) UpdatePaymentStatus(orderID uint, status string, paymentTime *time.Time) error {
	updates := map[string]interface{}{
		"payment_status": status,
		"payment_time":   paymentTime,
	}
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates).Error
}

// CreateLogistics 创建物流信息
func (r *OrderRepository) CreateLogistics(logistics *models.Logistics) error {
	return r.db.Create(logistics).Error
}

// UpdateLogistics 更新物流信息
func (r *OrderRepository) UpdateLogistics(logistics *models.Logistics) error {
	return r.db.Save(logistics).Error
}

// GetLogistics 获取订单物流信息
func (r *OrderRepository) GetLogistics(orderID uint) (*models.Logistics, error) {
	var logistics models.Logistics
	err := r.db.Where("order_id = ?", orderID).
		Preload("LogisticsTraces", func(db *gorm.DB) *gorm.DB {
			return db.Order("trace_time DESC")
		}).
		First(&logistics).Error
	if err != nil {
		return nil, err
	}
	return &logistics, nil
}

// AddLogisticsTrace 添加物流跟踪记录
func (r *OrderRepository) AddLogisticsTrace(trace *models.LogisticsTrace) error {
	return r.db.Create(trace).Error
}

// ListOrdersByStatus 管理员按状态查询订单
func (r *OrderRepository) ListOrdersByStatus(status string, page, pageSize int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64
	
	offset := (page - 1) * pageSize

	query := r.db.Model(&models.Order{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取订单列表
	err := query.Preload("Items.Product").
		Preload("Address").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, nickname, email")
		}).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, total, err
} 