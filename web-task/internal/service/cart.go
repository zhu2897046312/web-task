package service

import (
    "errors"
    "web-task/internal/models"
)

type CartService struct {
    *Service
}

func NewCartService(base *Service) *CartService {
    return &CartService{Service: base}
}

func (s *CartService) AddItem(userID, productID uint, quantity int) error {
    // 检查商品是否存在
    product, err := s.repoFactory.GetProductRepository().GetByID(productID)
    if err != nil {
        return errors.New("product not found")
    }

    // 检查库存
    if product.Stock < quantity {
        return errors.New("insufficient stock")
    }

    // 检查购物车是否已有该商品
    if existingItem, err := s.repoFactory.GetCartRepository().GetCartItem(userID, productID); err == nil {
        // 更新数量
        newQuantity := existingItem.Quantity + quantity
        return s.repoFactory.GetCartRepository().UpdateQuantity(existingItem.ID, newQuantity)
    }

    // 创建新的购物车项
    cartItem := &models.CartItem{
        UserID:    userID,
        ProductID: productID,
        Quantity:  quantity,
    }

    return s.repoFactory.GetCartRepository().Create(cartItem)
}

func (s *CartService) ListItems(userID uint) ([]models.CartItem, error) {
    return s.repoFactory.GetCartRepository().GetUserCart(userID)
}

func (s *CartService) UpdateQuantity(userID, itemID uint, quantity int) error {
    // 检查购物车项是否存在
    item, err := s.repoFactory.GetCartRepository().GetCartItem(userID, itemID)
    if err != nil {
        return errors.New("cart item not found")
    }

    // 检查库存
    product, err := s.repoFactory.GetProductRepository().GetByID(item.ProductID)
    if err != nil {
        return errors.New("product not found")
    }

    if product.Stock < quantity {
        return errors.New("insufficient stock")
    }

    return s.repoFactory.GetCartRepository().UpdateQuantity(itemID, quantity)
}

func (s *CartService) RemoveItem(userID, itemID uint) error {
    return s.repoFactory.GetCartRepository().Delete(&models.CartItem{ID: itemID})
}

func (s *CartService) ClearCart(userID uint) error {
    return s.repoFactory.GetCartRepository().ClearCart(userID)
} 