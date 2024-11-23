<template>
  <div class="home">
    <header class="hero">
      <h1>欢迎来到我们的商城</h1>
      <p>发现优质好物</p>
    </header>
    
    <section class="featured-products">
      <h2>精选商品</h2>
      <div class="product-grid">
        <ProductCard
          v-for="product in featuredProducts"
          :key="product.id"
          :product="product"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ProductCard from '../components/ProductCard.vue'
import { Product } from '../types'

const featuredProducts = ref<Product[]>([
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
  }
])
</script>

<style scoped>
.home {
  padding: 20px;
}

.hero {
  text-align: center;
  padding: 40px 0;
  background: #f5f5f5;
  margin-bottom: 40px;
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
  padding: 20px;
}
</style> 