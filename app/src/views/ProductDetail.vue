<template>
  <div class="product-detail" v-if="product">
    <h2>{{ product.name }}</h2>
    <div class="content">
      <img :src="product.image" :alt="product.name">
      <div class="info">
        <p class="description">{{ product.description }}</p>
        <p class="price">{{ formatPrice(product.price) }}</p>
        <button @click="addToCart">加入购物车</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from '../store'
import { productApi } from '../api/product'
import { formatPrice } from '../utils/format'
import type { Product } from '../types'

const route = useRoute()
const store = useStore()
const product = ref<Product | null>(null)

onMounted(async () => {
  try {
    const id = Number(route.params.id)
    const data = await productApi.getProductById(id)
    product.value = data
  } catch (error) {
    console.error('获取商品详情失败:', error)
  }
})

const addToCart = () => {
  if (product.value) {
    store.addToCart(product.value)
  }
}
</script>

<style lang="scss" scoped>
.product-detail {
  padding: 20px;
  
  .content {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
    margin-top: 20px;
  }
  
  img {
    width: 100%;
    height: auto;
  }
  
  .info {
    padding: 20px;
  }
  
  button {
    margin-top: 20px;
  }
}
</style> 