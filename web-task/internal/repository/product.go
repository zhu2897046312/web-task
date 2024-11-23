package repository

import (
    "web-task/internal/models"
    "gorm.io/gorm"
)

type ProductRepository struct {
    *BaseRepository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{
        BaseRepository: NewBaseRepository(db),
    }
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
    var product models.Product
    err := r.db.Preload("Reviews").First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) List(page, pageSize int) ([]models.Product, int64, error) {
    var products []models.Product
    var total int64

    query := r.db.Model(&models.Product{})
    
    // 获取总数
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 获取分页数据
    err := query.Offset((page - 1) * pageSize).
        Limit(pageSize).
        Find(&products).Error

    return products, total, err
}

func (r *ProductRepository) ListByCategory(category string) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Where("category = ?", category).Find(&products).Error
    return products, err
}

func (r *ProductRepository) UpdateStock(id uint, quantity int) error {
    return r.db.Model(&models.Product{}).
        Where("id = ? AND stock >= ?", id, quantity).
        UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *ProductRepository) UpdateSales(id uint, quantity int) error {
    return r.db.Model(&models.Product{}).
        Where("id = ?", id).
        UpdateColumn("sales", gorm.Expr("sales + ?", quantity)).Error
} 