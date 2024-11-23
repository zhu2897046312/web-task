<template>
  <div class="product-card">
    <img :src="product.image" :alt="product.name">
    <div class="product-info">
      <h3>{{ product.name }}</h3>
      <p class="price">¥{{ product.price }}</p>
      <button @click="addToCart">加入购物车</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Product } from '../types'
import { useStore } from '../store'

const store = useStore()
const props = defineProps<{
  product: Product
}>()

const addToCart = () => {
  store.addToCart(props.product)
}
</script>

<style scoped>
.product-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  transition: transform 0.2s;
}

.product-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
}

.product-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.product-info {
  padding: 15px;
}

.price {
  color: #e53935;
  font-weight: bold;
  margin: 10px 0;
}

button {
  width: 100%;
  padding: 8px;
  background: #1976d2;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background: #1565c0;
}
</style> 