package service

import (
	"errors"
	"web-task/internal/models"
)

type ProductService struct {
	*Service
}

func NewProductService(base *Service) *ProductService {
	return &ProductService{Service: base}
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.repoFactory.GetProductRepository().GetByID(id)
}

func (s *ProductService) ListProducts(page, pageSize int) ([]models.Product, int64, error) {
	return s.repoFactory.GetProductRepository().List(page, pageSize)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repoFactory.GetProductRepository().Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repoFactory.GetProductRepository().Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repoFactory.GetProductRepository().Delete(&models.Product{ID: id})
}

func (s *ProductService) CheckStock(productID uint, quantity int) error {
	product, err := s.GetProduct(productID)
	if err != nil {
		return err
	}
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}
	return nil
} 