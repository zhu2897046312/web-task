<template>
  <div class="cart">
    <h2>购物车</h2>
    <div class="cart-items" v-if="store.cart.length">
      <div v-for="item in store.cart" :key="item.id" class="cart-item">
        <img :src="item.image" :alt="item.name">
        <div class="item-info">
          <h3>{{ item.name }}</h3>
          <p>{{ formatPrice(item.price) }}</p>
          <div class="quantity">
            数量: {{ item.quantity }}
          </div>
        </div>
      </div>
      <div class="cart-total">
        总计: {{ formatPrice(store.cartTotal) }}
      </div>
    </div>
    <div v-else class="empty-cart">
      购物车是空的
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStore } from '../store'
import { formatPrice } from '../utils/format'

const store = useStore()
</script>

<style lang="scss" scoped>
.cart {
  padding: 20px;
}

.cart-item {
  display: grid;
  grid-template-columns: 100px 1fr;
  gap: 20px;
  padding: 20px;
  border-bottom: 1px solid #eee;
}

.cart-total {
  margin-top: 20px;
  text-align: right;
  font-weight: bold;
}

.empty-cart {
  text-align: center;
  padding: 40px;
  color: #999;
}
</style> 