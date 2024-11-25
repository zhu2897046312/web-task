import request from './request'
import type { Product, Review } from '@/types'

export const productApi = {
  // 获取商品列表
  getProducts(params?: {
    category?: string
    keyword?: string
    page?: number
    limit?: number
    sort?: string
  }) {
    return request.get<{
      items: Product[]
      total: number
    }>('/products', { params })
  },

  // 获取商品详情
  getProductById(id: number) {
    return request.get<Product>(`/products/${id}`)
  },

  // 获取推荐商品
  getFeaturedProducts() {
    return request.get<Product[]>('/products/featured')
  },

  // 获取商品评价
  getProductReviews(productId: number, params?: {
    page?: number
    limit?: number
    sort?: string
  }) {
    return request.get<{
      items: Review[]
      total: number
    }>(`/products/${productId}/reviews`, { params })
  },

  // 提交商品评价
  createReview(productId: number, data: {
    orderId: number
    rating: number
    content: string
    images?: string[]
  }) {
    return request.post<Review>(`/products/${productId}/reviews`, data)
  }
} 