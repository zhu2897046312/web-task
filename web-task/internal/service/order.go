package service

import (
	"errors"
	"web-task/internal/models"
	"web-task/internal/repository"
	"gorm.io/gorm"
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
		// 创建临时的仓储工厂，使用事务连接
		txRepoFactory := repository.NewRepositoryFactory(tx)
		
		// 检查并扣减库存
		var totalAmount float64
		for _, item := range items {
			product, err := txRepoFactory.GetProductRepository().GetByID(item.ProductID)
			if err != nil {
				return err
			}

			if product.Stock < item.Quantity {
				return errors.New("insufficient stock")
			}

			// 扣减库存
			if err := txRepoFactory.GetProductRepository().UpdateStock(product.ID, item.Quantity); err != nil {
				return err
			}

			// 更新销量
			if err := txRepoFactory.GetProductRepository().UpdateSales(product.ID, item.Quantity); err != nil {
				return err
			}

			item.Price = product.Price
			item.Subtotal = product.Price * float64(item.Quantity)
			totalAmount += item.Subtotal
		}

		// 创建订单
		order := &models.Order{
			UserID:      userID,
			Items:       items,
			TotalAmount: totalAmount,
			Status:      "pending",
			AddressID:   addressID,
		}

		if err := txRepoFactory.GetOrderRepository().Create(order); err != nil {
			return err
		}

		// 清空购物车
		if err := txRepoFactory.GetCartRepository().ClearCart(userID); err != nil {
			return err
		}

		// 保存结果
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