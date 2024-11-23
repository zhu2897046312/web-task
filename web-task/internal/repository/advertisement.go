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

func (r *AdvertisementRepository) GetActiveByPosition(position string) ([]models.Advertisement, error) {
    var ads []models.Advertisement
    now := time.Now()
    
    err := r.db.Where("position = ? AND status = ? AND start_time <= ? AND end_time >= ?",
        position, true, now, now).
        Find(&ads).Error
    
    return ads, err
}

func (r *AdvertisementRepository) ListActive() ([]models.Advertisement, error) {
    var ads []models.Advertisement
    now := time.Now()
    
    err := r.db.Where("status = ? AND start_time <= ? AND end_time >= ?",
        true, now, now).
        Find(&ads).Error
    
    return ads, err
}

func (r *AdvertisementRepository) UpdateStatus(id uint, status bool) error {
    return r.db.Model(&models.Advertisement{}).
        Where("id = ?", id).
        Update("status", status).Error
} 