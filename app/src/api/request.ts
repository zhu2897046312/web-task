import axios, { AxiosResponse } from 'axios'
import type { ApiResponse } from '@/types'
import { storage } from '@/utils/storage'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    const token = storage.get('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse<ApiResponse<any>>) => {
    const res = response.data
    if (res.code !== 0) {
      // 统一处理错误
      console.error(res.message)
      return Promise.reject(new Error(res.message))
    }
    return res.data
  },
  (error) => {
    if (error.response?.status === 401) {
      // 处理未授权
      storage.remove('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default request 