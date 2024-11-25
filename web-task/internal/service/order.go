package service

import (
	"errors"
	"web-task/internal/models"
	"web-task/internal/repository"
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
	"fmt"
	"math/rand"
	"time"
)

type OrderService struct {
	*Service
}

func NewOrderService(base *Service) *OrderService {
	return &OrderService{Service: base}
}

func (s *OrderService) CreateOrder(userID uint, items []models.OrderItem, addressID uint) (*models.Order, error) {
	var result *models.Order

	err := s.repoFactory.GetDB().Transaction(func(tx *gorm.DB) error {
		txRepoFactory := repository.NewRepositoryFactory(tx)
		
		totalAmount := decimal.NewFromFloat(0)
		
		for i := range items {
			product, err := txRepoFactory.GetProductRepository().GetByID(items[i].ProductID)
			if err != nil {
				return err
			}

			if product.Stock < items[i].Quantity {
				return errors.New("insufficient stock")
			}

			if err := txRepoFactory.GetProductRepository().UpdateStock(product.ID, items[i].Quantity); err != nil {
				return err
			}

			if err := txRepoFactory.GetProductRepository().UpdateSales(product.ID, items[i].Quantity); err != nil {
				return err
			}

			items[i].Price = product.Price
			itemTotal := product.Price.Mul(decimal.NewFromInt(int64(items[i].Quantity)))
			totalAmount = totalAmount.Add(itemTotal)
		}

		orderNumber := generateOrderNumber()

		order := &models.Order{
			UserID:      userID,
			OrderNumber: orderNumber,
			TotalAmount: totalAmount,
			Status:      "pending",
			AddressID:   addressID,
		}

		if err := txRepoFactory.GetOrderRepository().Create(order); err != nil {
			return err
		}

		for _, item := range items {
			item.OrderID = order.ID
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		if err := txRepoFactory.GetCartRepository().ClearCart(userID); err != nil {
			return err
		}

		result = order
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	return s.repoFactory.GetOrderRepository().GetByID(id)
}

func (s *OrderService) ListUserOrders(userID uint) ([]models.Order, error) {
	return s.repoFactory.GetOrderRepository().ListByUserID(userID)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.repoFactory.GetOrderRepository().UpdateStatus(orderID, status)
}

func generateOrderNumber() string {
	timestamp := time.Now().Format("20060102150405")
	random := rand.Intn(1000)
	return fmt.Sprintf("ORD%s%03d", timestamp, random)
} 