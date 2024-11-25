package service

import (
	"errors"
	"web-task/internal/models"
)

type AdvertisementService struct {
	*Service
}

func NewAdvertisementService(base *Service) *AdvertisementService {
	return &AdvertisementService{Service: base}
}

func (s *AdvertisementService) CreateAd(ad *models.Advertisement) error {
	// 验证时间
	if ad.StartTime.After(ad.EndTime) {
		return errors.New("start time cannot be after end time")
	}
	return s.repoFactory.GetAdvertisementRepository().Create(ad)
}

func (s *AdvertisementService) GetActiveAds(position string) ([]models.Advertisement, error) {
	return s.repoFactory.GetAdvertisementRepository().GetActiveByPosition(position)
}

func (s *AdvertisementService) ListActiveAds() ([]models.Advertisement, error) {
	return s.repoFactory.GetAdvertisementRepository().ListActive()
}

func (s *AdvertisementService) UpdateAdStatus(id uint, status bool) error {
	return s.repoFactory.GetAdvertisementRepository().UpdateStatus(id, status)
}

func (s *AdvertisementService) DeleteAd(id uint) error {
	return s.repoFactory.GetAdvertisementRepository().Delete(&models.Advertisement{ID: id})
} 
