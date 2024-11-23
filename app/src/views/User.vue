<template>
  <div class="user">
    <h2>个人中心</h2>
    <div v-if="store.user">
      <p>用户名：{{ store.user.username }}</p>
      <p>邮箱：{{ store.user.email }}</p>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStore } from '../store'

const store = useStore()
</script>

<style lang="scss" scoped>
.user {
  padding: 20px;
}
</style> 