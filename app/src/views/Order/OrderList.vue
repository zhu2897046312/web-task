<template>
  <div class="order-list">
    <h2>我的订单</h2>
    <div class="order-tabs">
      <button 
        v-for="status in orderStatuses" 
        :key="status"
        :class="{ active: currentStatus === status }"
        @click="currentStatus = status"
      >
        {{ getStatusText(status) }}
      </button>
    </div>
    
    <div class="orders">
      <div v-for="order in filteredOrders" :key="order.id" class="order-card">
        <div class="order-header">
          <span>订单号：{{ order.id }}</span>
          <span>{{ getStatusText(order.status) }}</span>
        </div>
        
        <div class="order-items">
          <div v-for="item in order.items" :key="item.id" class="order-item">
            <img :src="item.image" :alt="item.productName">
            <div class="item-info">
              <h4>{{ item.productName }}</h4>
              <p>{{ formatPrice(item.price) }} x {{ item.quantity }}</p>
            </div>
          </div>
        </div>
        
        <div class="order-footer">
          <p>总计：{{ formatPrice(order.totalAmount) }}</p>
          <div class="order-actions">
            <button v-if="order.status === 'pending'" @click="payOrder(order.id)">
              立即支付
            </button>
            <button v-if="order.status === 'shipped'" @click="confirmReceived(order.id)">
              确认收货
            </button>
            <button v-if="order.status === 'delivered'" @click="writeReview(order.id)">
              评价
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Order, OrderStatus } from '@/types'
import { formatPrice } from '@/utils/format'

const currentStatus = ref<OrderStatus | 'all'>('all')
const orders = ref<Order[]>([])

const orderStatuses = ['all', ...Object.values(OrderStatus)]

const getStatusText = (status: string) => {
  const statusMap = {
    all: '全部',
    pending: '待付款',
    paid: '待发货',
    shipped: '待收货',
    delivered: '已完成',
    cancelled: '已取消'
  }
  return statusMap[status] || status
}

const filteredOrders = computed(() => {
  if (currentStatus.value === 'all') {
    return orders.value
  }
  return orders.value.filter(order => order.status === currentStatus.value)
})

const payOrder = (orderId: number) => {
  // 实现支付逻辑
}

const confirmReceived = (orderId: number) => {
  // 实现确认收货逻辑
}

const writeReview = (orderId: number) => {
  // 实现评价逻辑
}
</script>

<style lang="scss" scoped>
.order-list {
  padding: 20px;
}

.order-tabs {
  margin: 20px 0;
  button {
    margin-right: 10px;
    padding: 8px 16px;
    border: none;
    background: none;
    cursor: pointer;
    
    &.active {
      color: $primary-color;
      border-bottom: 2px solid $primary-color;
    }
  }
}

.order-card {
  border: 1px solid $border-color;
  border-radius: 8px;
  margin-bottom: 20px;
  overflow: hidden;
}

.order-header {
  @include flex(row, space-between);
  padding: 15px;
  background: #f8f9fa;
}

.order-items {
  padding: 15px;
}

.order-item {
  @include flex(row, flex-start);
  margin-bottom: 10px;
  
  img {
    width: 80px;
    height: 80px;
    object-fit: cover;
    margin-right: 15px;
  }
}

.order-footer {
  @include flex(row, space-between, center);
  padding: 15px;
  background: #f8f9fa;
  
  .order-actions {
    button {
      margin-left: 10px;
    }
  }
}
</style> 