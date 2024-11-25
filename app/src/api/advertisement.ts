import request from './request'
import type { Advertisement } from '@/types'

export const advertisementApi = {
  // 获取广告列表
  getAdvertisements(position: string) {
    return request.get<Advertisement[]>('/advertisements', {
      params: { position }
    })
  }
} 