package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"web-task/internal/models"
	"web-task/internal/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderService struct {
	*Service
}

func NewOrderService(base *Service) *OrderService {
	return &OrderService{Service: base}
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(userID uint, items []models.OrderItem, addressID uint) (*models.Order, error) {
	var result *models.Order

	err := s.repoFactory.GetDB().Transaction(func(tx *gorm.DB) error {
		txRepoFactory := repository.NewRepositoryFactory(tx)
		
		// 验证地址是否存在且属于该用户
		address, err := txRepoFactory.GetUserRepository().GetAddressByID(addressID)
		if err != nil || address.UserID != userID {
			return errors.New("invalid address")
		}

		totalAmount := decimal.NewFromFloat(0)
		
		// 验证商品并计算总金额
		for i := range items {
			product, err := txRepoFactory.GetProductRepository().GetByID(items[i].ProductID)
			if err != nil {
				return err
			}

			if product.Stock < items[i].Quantity {
				return fmt.Errorf("insufficient stock for product: %s", product.Name)
			}

			// 更新库存和销量
			if err := txRepoFactory.GetProductRepository().UpdateStock(product.ID, -items[i].Quantity); err != nil {
				return err
			}
			if err := txRepoFactory.GetProductRepository().UpdateSales(product.ID, items[i].Quantity); err != nil {
				return err
			}

			items[i].Price = product.Price
			itemTotal := product.Price.Mul(decimal.NewFromInt(int64(items[i].Quantity)))
			totalAmount = totalAmount.Add(itemTotal)
		}

		// 创建订单
		order := &models.Order{
			UserID:        userID,
			OrderNumber:   generateOrderNumber(),
			TotalAmount:   totalAmount,
			Status:        "pending",
			AddressID:     addressID,
			PaymentStatus: "unpaid",
		}

		if err := txRepoFactory.GetOrderRepository().Create(order); err != nil {
			return err
		}

		// 创建订单项
		for _, item := range items {
			item.OrderID = order.ID
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		// 清空购物车
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

// GetOrder 获取订单详情
func (s *OrderService) GetOrder(id uint, userID uint) (*models.Order, error) {
	order, err := s.repoFactory.GetOrderRepository().GetByID(id)
	if err != nil {
		return nil, err
	}
	
	// 验证订单所属权
	if order.UserID != userID {
		return nil, errors.New("permission denied")
	}
	
	return order, nil
}

// ListUserOrders 获取用户订单列表
func (s *OrderService) ListUserOrders(userID uint, page, pageSize int) ([]models.Order, int64, error) {
	return s.repoFactory.GetOrderRepository().ListByUserID(userID, page, pageSize)
}

// UpdateOrderStatus 更新订单状态
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.repoFactory.GetOrderRepository().UpdateStatus(orderID, status)
}

// UpdatePaymentStatus 更新支付状态
func (s *OrderService) UpdatePaymentStatus(orderID uint, status string) error {
	var paymentTime *time.Time
	if status == "paid" {
		now := time.Now()
		paymentTime = &now
	}
	return s.repoFactory.GetOrderRepository().UpdatePaymentStatus(orderID, status, paymentTime)
}

// CreateLogistics 创建物流信息
func (s *OrderService) CreateLogistics(logistics *models.Logistics) error {
	return s.repoFactory.GetOrderRepository().CreateLogistics(logistics)
}

// UpdateLogistics 更新物流信息
func (s *OrderService) UpdateLogistics(logistics *models.Logistics) error {
	return s.repoFactory.GetOrderRepository().UpdateLogistics(logistics)
}

// GetLogistics 获取订单物流信息
func (s *OrderService) GetLogistics(orderID uint) (*models.Logistics, error) {
	return s.repoFactory.GetOrderRepository().GetLogistics(orderID)
}

// AddLogisticsTrace 添加物流跟踪记录
func (s *OrderService) AddLogisticsTrace(trace *models.LogisticsTrace) error {
	return s.repoFactory.GetOrderRepository().AddLogisticsTrace(trace)
}

// ListOrdersByStatus 管理员按状态查询订单
func (s *OrderService) ListOrdersByStatus(status string, page, pageSize int) ([]models.Order, int64, error) {
	return s.repoFactory.GetOrderRepository().ListOrdersByStatus(status, page, pageSize)
}

// generateOrderNumber 生成订单号
func generateOrderNumber() string {
	timestamp := time.Now().Format("20060102150405")
	random := rand.Intn(1000)
	return fmt.Sprintf("ORD%s%03d", timestamp, random)
} 