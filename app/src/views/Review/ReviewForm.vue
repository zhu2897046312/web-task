<template>
  <div class="review-form">
    <h3>商品评价</h3>
    <div class="rating">
      <span>评分：</span>
      <div class="stars">
        <i 
          v-for="n in 5" 
          :key="n"
          :class="['star', { active: n <= rating }]"
          @click="rating = n"
        ></i>
      </div>
    </div>
    
    <div class="content">
      <textarea
        v-model="content"
        placeholder="请输入您的评价内容..."
        rows="4"
      ></textarea>
    </div>
    
    <div class="images">
      <div 
        v-for="(img, index) in images" 
        :key="index" 
        class="image-item"
      >
        <img :src="img" alt="评价图片">
        <button @click="removeImage(index)" class="remove-btn">×</button>
      </div>
      
      <div v-if="images.length < 5" class="upload-btn" @click="uploadImage">
        <i class="plus">+</i>
        <span>上传图片</span>
      </div>
    </div>
    
    <div class="actions">
      <button @click="cancel">取消</button>
      <button @click="submit" :disabled="!isValid">提交评价</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Review } from '@/types'

const props = defineProps<{
  orderId: number
  productId: number
}>()

const emit = defineEmits<{
  (e: 'success'): void
  (e: 'cancel'): void
}>()

const rating = ref(5)
const content = ref('')
const images = ref<string[]>([])

const isValid = computed(() => {
  return rating.value > 0 && content.value.trim().length > 0
})

const uploadImage = () => {
  // 实现图片上传逻辑
}

const removeImage = (index: number) => {
  images.value.splice(index, 1)
}

const submit = async () => {
  if (!isValid.value) return
  
  try {
    // 提交评价逻辑
    emit('success')
  } catch (error) {
    console.error('提交评价失败:', error)
  }
}

const cancel = () => {
  emit('cancel')
}
</script>

<style lang="scss" scoped>
.review-form {
  padding: 20px;
}

.rating {
  margin: 20px 0;
  @include flex(row, flex-start, center);
  
  .stars {
    margin-left: 10px;
    
    .star {
      cursor: pointer;
      color: #ddd;
      font-size: 24px;
      
      &.active {
        color: #ffd700;
      }
    }
  }
}

.content {
  textarea {
    width: 100%;
    padding: 10px;
    border: 1px solid $border-color;
    border-radius: 4px;
    resize: vertical;
  }
}

.images {
  @include flex(row, flex-start);
  flex-wrap: wrap;
  margin: 20px 0;
  
  .image-item {
    position: relative;
    width: 100px;
    height: 100px;
    margin: 0 10px 10px 0;
    
    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    
    .remove-btn {
      position: absolute;
      top: -10px;
      right: -10px;
      width: 20px;
      height: 20px;
      border-radius: 50%;
      background: rgba(0, 0, 0, 0.5);
      color: white;
      border: none;
      cursor: pointer;
    }
  }
  
  .upload-btn {
    @include flex(column, center, center);
    width: 100px;
    height: 100px;
    border: 1px dashed $border-color;
    cursor: pointer;
    
    &:hover {
      border-color: $primary-color;
      color: $primary-color;
    }
  }
}

.actions {
  @include flex(row, flex-end);
  gap: 10px;
  
  button {
    padding: 8px 20px;
    border-radius: 4px;
    cursor: pointer;
    
    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
  }
}
</style> 