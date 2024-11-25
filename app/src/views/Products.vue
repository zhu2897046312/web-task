<template>
  <div class="products">
    <h2>商品列表</h2>
    <div class="product-grid">
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product="product"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { productApi } from '../api/product'
import ProductCard from '../components/ProductCard.vue'
import type { Product } from '../types'

const products = ref<Product[]>([
  {
    id: 1,
    name: "测试商品1",
    price: 99.99,
    description: "这是一个测试商品",
    image: "https://via.placeholder.com/200",
    category: "测试分类"
  },
  {
    id: 2,
    name: "测试商品2",
    price: 199.99,
    description: "这是另一个测试商品",
    image: "https://via.placeholder.com/200",
    category: "测试分类"
  },
  {
    id: 3,
    name: "测试商品3",
    price: 299.99,
    description: "这是第三个测试商品",
    image: "https://via.placeholder.com/200",
    category: "测试分类"
  }
])

// 暂时注释掉 API 调用，使用模拟数据
// onMounted(async () => {
//   try {
//     const data = await productApi.getProducts()
//     products.value = data
//   } catch (error) {
//     console.error('获取商品列表失败:', error)
//   }
// })
</script>

<style lang="scss" scoped>
.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
  padding: 20px;
}
</style> 