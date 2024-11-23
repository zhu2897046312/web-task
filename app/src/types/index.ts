export interface Product {
  id: number
  name: string
  price: number
  description: string
  image: string
  category: string
  stock: number
  sales: number
  rating: number
  reviews: number
  createdAt: string
}

export interface CartItem extends Product {
  quantity: number
}

export interface User {
  id: number
  username: string
  email: string
  phone?: string
  avatar?: string
  createdAt: string
}

export interface LoginForm {
  username: string
  password: string
}

export interface RegisterForm extends LoginForm {
  email: string
  phone?: string
}

export interface Order {
  id: number
  userId: number
  items: OrderItem[]
  totalAmount: number
  status: OrderStatus
  paymentMethod: PaymentMethod
  address: Address
  logistics?: Logistics
  createdAt: string
}

export interface OrderItem {
  id: number
  productId: number
  productName: string
  price: number
  quantity: number
  subtotal: number
}

export enum OrderStatus {
  Pending = 'pending',
  Paid = 'paid',
  Shipped = 'shipped',
  Delivered = 'delivered',
  Cancelled = 'cancelled'
}

export enum PaymentMethod {
  Alipay = 'alipay',
  WechatPay = 'wechat',
  CreditCard = 'credit_card'
}

export interface Address {
  id: number
  userId: number
  receiver: string
  phone: string
  province: string
  city: string
  district: string
  detail: string
  isDefault: boolean
}

export interface Logistics {
  id: number
  orderId: number
  trackingNumber: string
  carrier: string
  status: LogisticsStatus
  traces: LogisticsTrace[]
}

export interface LogisticsTrace {
  time: string
  location: string
  description: string
}

export enum LogisticsStatus {
  Pending = 'pending',
  InTransit = 'in_transit',
  Delivered = 'delivered'
}

export interface Review {
  id: number
  userId: number
  productId: number
  orderId: number
  rating: number
  content: string
  images?: string[]
  createdAt: string
}

export interface Advertisement {
  id: number
  title: string
  image: string
  link: string
  position: string
  startTime: string
  endTime: string
  status: boolean
}

export interface ApiResponse<T> {
  code: number
  message: string
  data: T
} 