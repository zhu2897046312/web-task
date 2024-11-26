package service

import (
	"errors"
	"time"
	"web-task/internal/models"
)

type AdvertisementService struct {
	*Service
}

func NewAdvertisementService(base *Service) *AdvertisementService {
	return &AdvertisementService{Service: base}
}

// CreateAd 创建广告
func (s *AdvertisementService) CreateAd(ad *models.Advertisement) error {
	// 验证时间
	if ad.StartTime.After(ad.EndTime) {
		return errors.New("start time cannot be after end time")
	}

	// 验证开始时间不能早于当前时间
	if ad.StartTime.Before(time.Now()) {
		return errors.New("start time cannot be earlier than current time")
	}

	// 设置默认状态
	if ad.Status == "" {
		ad.Status = "active"
	}

	return s.repoFactory.GetAdvertisementRepository().Create(ad)
}

// GetAd 获取广告详情
func (s *AdvertisementService) GetAd(id uint) (*models.Advertisement, error) {
	return s.repoFactory.GetAdvertisementRepository().GetByID(id)
}

// UpdateAd 更新广告信息
func (s *AdvertisementService) UpdateAd(ad *models.Advertisement) error {
	// 验证时间
	if ad.StartTime.After(ad.EndTime) {
		return errors.New("start time cannot be after end time")
	}

	return s.repoFactory.GetAdvertisementRepository().Update(ad)
}

// DeleteAd 删除广告
func (s *AdvertisementService) DeleteAd(id uint) error {
	return s.repoFactory.GetAdvertisementRepository().Delete(id)
}

// ListAllAds 获取所有广告
func (s *AdvertisementService) ListAllAds(page, pageSize int) ([]models.Advertisement, int64, error) {
	return s.repoFactory.GetAdvertisementRepository().ListAll(page, pageSize)
}

// GetActiveAds 获取指定位置的有效广告
func (s *AdvertisementService) GetActiveAds(position string) ([]models.Advertisement, error) {
	return s.repoFactory.GetAdvertisementRepository().GetActiveByPosition(position)
}

// ListActiveAds 获取所有有效广告
func (s *AdvertisementService) ListActiveAds() ([]models.Advertisement, error) {
	return s.repoFactory.GetAdvertisementRepository().ListActive()
}

// UpdateAdStatus 更新广告状态
func (s *AdvertisementService) UpdateAdStatus(id uint, status string) error {
	// 验证状态值
	validStatuses := map[string]bool{
		"active":   true,
		"inactive": true,
	}
	if !validStatuses[status] {
		return errors.New("invalid status value")
	}

	return s.repoFactory.GetAdvertisementRepository().UpdateStatus(id, status)
}

// ListAdsByStatus 根据状态获取广告列表
func (s *AdvertisementService) ListAdsByStatus(status string, page, pageSize int) ([]models.Advertisement, int64, error) {
	return s.repoFactory.GetAdvertisementRepository().ListByStatus(status, page, pageSize)
} 
