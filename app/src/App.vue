<template>
  <div class="app">
    <nav class="nav">
      <router-link to="/">首页</router-link>
      <router-link to="/products">商品列表</router-link>
      <router-link to="/cart">购物车</router-link>
      <router-link to="/user">个人中心</router-link>
    </nav>
    
    <main class="main">
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup lang="ts">
// 组件逻辑
</script>

<style lang="scss">
@import '@/assets/styles/reset.scss';
@import '@/assets/styles/variables.scss';
@import '@/assets/styles/mixins.scss';

.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.nav {
  @include flex(row, space-around, center);
  padding: $spacing-base;
  background-color: $primary-color;
  
  a {
    color: white;
    padding: $spacing-small $spacing-base;
    border-radius: 4px;
    
    &.router-link-active {
      background-color: rgba(255, 255, 255, 0.2);
    }
  }
}

.main {
  flex: 1;
  padding: $spacing-base;
}
</style> 