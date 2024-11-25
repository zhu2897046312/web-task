import request from './request'
import type { CartItem } from '@/types'

export const cartApi = {
  // 获取购物车列表
  getCartItems() {
    return request.get<CartItem[]>('/cart')
  },

  // 添加商品到购物车
  addToCart(data: { productId: number; quantity: number }) {
    return request.post<CartItem>('/cart', data)
  },

  // 更新购物车商品数量
  updateQuantity(productId: number, quantity: number) {
    return request.put(`/cart/${productId}`, { quantity })
  },

  // 删除购物车商品
  removeFromCart(productId: number) {
    return request.delete(`/cart/${productId}`)
  },

  // 清空购物车
  clearCart() {
    return request.delete('/cart')
  }
} 