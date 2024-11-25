import request from './request'
import type { Order, OrderStatus } from '@/types'

export const orderApi = {
  // 创建订单
  createOrder(data: {
    items: { productId: number; quantity: number }[]
    addressId: number
    paymentMethod: string
  }) {
    return request.post<Order>('/orders', data)
  },

  // 获取订单列表
  getOrders(params?: {
    status?: OrderStatus
    page?: number
    limit?: number
  }) {
    return request.get<{
      items: Order[]
      total: number
    }>('/orders', { params })
  },

  // 获取订单详情
  getOrderById(id: number) {
    return request.get<Order>(`/orders/${id}`)
  },

  // 取消订单
  cancelOrder(id: number) {
    return request.put(`/orders/${id}/cancel`)
  },

  // 支付订单
  payOrder(id: number, paymentMethod: string) {
    return request.post(`/orders/${id}/pay`, { paymentMethod })
  },

  // 确认收货
  confirmReceived(id: number) {
    return request.put(`/orders/${id}/confirm`)
  },

  // 获取物流信息
  getLogistics(id: number) {
    return request.get(`/orders/${id}/logistics`)
  }
} 