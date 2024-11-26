package repository

import (
    "time"
    "web-task/internal/models"
    "gorm.io/gorm"
)

type AdvertisementRepository struct {
    *BaseRepository
}

func NewAdvertisementRepository(db *gorm.DB) *AdvertisementRepository {
    return &AdvertisementRepository{
        BaseRepository: NewBaseRepository(db),
    }
}

// Create 创建广告
func (r *AdvertisementRepository) Create(ad *models.Advertisement) error {
    return r.db.Create(ad).Error
}

// GetByID 获取广告详情
func (r *AdvertisementRepository) GetByID(id uint) (*models.Advertisement, error) {
    var ad models.Advertisement
    err := r.db.First(&ad, id).Error
    if err != nil {
        return nil, err
    }
    return &ad, nil
}

// Update 更新广告信息
func (r *AdvertisementRepository) Update(ad *models.Advertisement) error {
    return r.db.Save(ad).Error
}

// Delete 删除广告(软删除)
func (r *AdvertisementRepository) Delete(id uint) error {
    return r.db.Delete(&models.Advertisement{}, id).Error
}

// ListAll 获取所有广告(支持分页)
func (r *AdvertisementRepository) ListAll(page, pageSize int) ([]models.Advertisement, int64, error) {
    var ads []models.Advertisement
    var total int64
    
    offset := (page - 1) * pageSize

    // 获取总数
    if err := r.db.Model(&models.Advertisement{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 获取分页数据
    err := r.db.Offset(offset).
        Limit(pageSize).
        Order("created_at DESC").
        Find(&ads).Error

    return ads, total, err
}

// GetActiveByPosition 获取指定位置的有效广告
func (r *AdvertisementRepository) GetActiveByPosition(position string) ([]models.Advertisement, error) {
    var ads []models.Advertisement
    now := time.Now()
    
    err := r.db.Where("position = ? AND status = ? AND start_time <= ? AND end_time >= ?",
        position, "active", now, now).
        Order("created_at DESC").
        Find(&ads).Error
    
    return ads, err
}

// ListActive 获取所有有效广告
func (r *AdvertisementRepository) ListActive() ([]models.Advertisement, error) {
    var ads []models.Advertisement
    now := time.Now()
    
    err := r.db.Where("status = ? AND start_time <= ? AND end_time >= ?",
        "active", now, now).
        Order("created_at DESC").
        Find(&ads).Error
    
    return ads, err
}

// UpdateStatus 更新广告状态
func (r *AdvertisementRepository) UpdateStatus(id uint, status string) error {
    return r.db.Model(&models.Advertisement{}).
        Where("id = ?", id).
        Update("status", status).Error
}

// ListByStatus 根据状态获取广告列表
func (r *AdvertisementRepository) ListByStatus(status string, page, pageSize int) ([]models.Advertisement, int64, error) {
    var ads []models.Advertisement
    var total int64
    
    offset := (page - 1) * pageSize

    query := r.db.Model(&models.Advertisement{})
    if status != "" {
        query = query.Where("status = ?", status)
    }

    // 获取总数
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 获取分页数据
    err := query.Offset(offset).
        Limit(pageSize).
        Order("created_at DESC").
        Find(&ads).Error

    return ads, total, err
} 