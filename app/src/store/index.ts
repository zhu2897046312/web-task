import { defineStore } from 'pinia'
import { Product, CartItem } from '../types'

export const useStore = defineStore('main', {
  state: () => ({
    products: [] as Product[],
    cart: [] as CartItem[],
    user: null as any
  }),
  
  getters: {
    cartTotal: (state) => {
      return state.cart.reduce((total, item) => total + item.price * item.quantity, 0)
    }
  },
  
  actions: {
    addToCart(product: Product) {
      const existingItem = this.cart.find(item => item.id === product.id)
      if (existingItem) {
        existingItem.quantity++
      } else {
        this.cart.push({
          ...product,
          quantity: 1
        })
      }
    }
  }
}) 