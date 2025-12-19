<!--
  产品展示区块组件
-->
<template>
  <div class="product-section-wrapper" :style="{ maxWidth: maxWidth, width: '100%', marginBottom: spacing }">
    <!-- 标题区域 -->
    <div class="flex flex-col text-black justify-start items-start">
      <div class="category-label text-base text-gray-500 mb-2">{{ category }}</div>
      <h2 class="product-title text-2xl  mb-6">{{ title }}</h2>
    </div>
    
    <!-- 描述文字 -->
    <div class="product-description">
      <p class="text-lg leading-relaxed text-gray-700">
        {{ description }}
      </p>
    </div>

    <!-- 产品图片网格 -->
    <div v-if="products && products.length > 0" class="product-grid" :class="`grid-cols-${columns}`">
      <div 
        v-for="product in products" 
        :key="product.id"
        class="product-item"
      >
     
        <router-link :to="`/product/${product.id}`" class="product-link">
          <div class="product-image-wrapper">
         
            <img 
             src="@/assets/images/02.png" 
              :alt="product.name"
              loading="lazy"
              class="product-image"
            />
          </div>
          <div class="product-name">{{ product.name }}</div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  // 分类标签（如：经典）
  category: {
    type: String,
    default: '经典'
  },
  // 产品标题
  title: {
    type: String,
    required: true
  },
  // 产品描述
  description: {
    type: String,
    required: true
  },
  // 产品列表
  products: {
    type: Array,
    default: () => []
  },
  // 网格列数
  columns: {
    type: Number,
    default: 5
  },
  // 最大宽度
  maxWidth: {
    type: String,
    default: '1000px'
  },
  // 底部间距
  spacing: {
    type: String,
    default: '30px'
  }
})
</script>

<style scoped>
.product-section-wrapper {
  margin:0 80px;
  width: 100%;
}

.category-label {
  text-transform: uppercase;
  letter-spacing: 1px;
  font-weight: 400;
}

.product-title {
  line-height: 1.2;
  color: #333;
}

.product-description {
  margin-top: 20px;
  margin-bottom: 40px;
}

.product-description p {
  line-height: 1.8;
}

/* 产品网格 */
.product-grid {
  display: grid;
  gap: 30px;
  margin-top: 40px;
}

.product-grid.grid-cols-3 {
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
}

.product-grid.grid-cols-4 {
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
}

.product-grid.grid-cols-5 {
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
}

.product-grid.grid-cols-6 {
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
}

.product-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: transform 0.3s ease;
}

.product-item:hover {
  transform: translateY(-8px);
}

.product-link {
  text-decoration: none;
  color: inherit;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.product-image-wrapper {
  width: 100%;
  aspect-ratio: 1 / 1.2;
  display: flex;
  align-items: center;
  justify-content: center;

  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 15px;
  /* box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.3s ease; */
}

.product-item:hover .product-image-wrapper {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
  padding: 10px;
}

.product-name {
  text-align: center;
  font-size: 14px;
  color: #333;
  line-height: 1.4;
  font-weight: 500;
  transition: color 0.3s ease;
}

.product-item:hover .product-name {
  color: #c41e3a;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .product-section-wrapper {
    padding: 0 15px;
    margin-bottom: 20px !important;
  }

  .category-label {
    font-size: 12px;
    margin-bottom: 8px;
  }

  .product-title {
    font-size: 1.75rem;
    line-height: 1.3;
  }
  
  .product-description {
    margin-top: 15px;
    margin-bottom: 30px;
  }
  
  .product-description p {
    font-size: 14px;
    line-height: 1.6;
  }

  .product-grid {
    gap: 20px;
    margin-top: 30px;
  }

  .product-grid.grid-cols-3,
  .product-grid.grid-cols-4,
  .product-grid.grid-cols-5,
  .product-grid.grid-cols-6 {
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  }

  .product-name {
    font-size: 12px;
  }
}

@media (max-width: 480px) {
  .product-title {
    font-size: 1.5rem;
  }
  
  .product-description p {
    font-size: 13px;
  }

  .product-grid {
    gap: 15px;
  }

  .product-grid.grid-cols-3,
  .product-grid.grid-cols-4,
  .product-grid.grid-cols-5,
  .product-grid.grid-cols-6 {
    grid-template-columns: repeat(2, 1fr);
  }

  .product-image-wrapper {
    margin-bottom: 10px;
  }

  .product-name {
    font-size: 11px;
  }
}
</style>
