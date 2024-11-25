import request from './request'
import type { User, LoginForm, RegisterForm, Address } from '@/types'

export const userApi = {
  // 用户登录
  login(data: LoginForm) {
    return request.post<{ token: string; user: User }>('/auth/login', data)
  },

  // 用户注册
  register(data: RegisterForm) {
    return request.post<{ token: string; user: User }>('/auth/register', data)
  },

  // 获取用户信息
  getUserInfo() {
    return request.get<User>('/user/info')
  },

  // 更新用户信息
  updateUserInfo(data: Partial<User>) {
    return request.put<User>('/user/info', data)
  },

  // 获取用户地址列表
  getAddresses() {
    return request.get<Address[]>('/user/addresses')
  },

  // 添加地址
  addAddress(data: Omit<Address, 'id' | 'userId'>) {
    return request.post<Address>('/user/addresses', data)
  },

  // 更新地址
  updateAddress(id: number, data: Partial<Address>) {
    return request.put<Address>(`/user/addresses/${id}`, data)
  },

  // 删除地址
  deleteAddress(id: number) {
    return request.delete(`/user/addresses/${id}`)
  },

  // 设置默认地址
  setDefaultAddress(id: number) {
    return request.put(`/user/addresses/${id}/default`)
  }
} 