<template>
  <div class="product-detail">
    <!-- 语言切换器 -->
    <div class="header">
      <LanguageSwitcher />
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading">
      {{ $t('common.loading') }}
    </div>

    <!-- 产品详情 -->
    <div v-else-if="product" class="content">
      <!-- 产品图片 -->
      <div class="product-image">
        <img :src="product.image" :alt="product.name" />
      </div>

      <!-- 产品信息 -->
      <div class="product-info">
        <h1>{{ product.name }}</h1>
        <p class="subtitle">{{ product.subtitle }}</p>

        <!-- 价格和规格 -->
        <div class="price-info">
          <span class="label">{{ $t('product.price') }}:</span>
          <span class="price">€{{ product.price }}</span>
        </div>

        <div class="weight-info">
          <span class="label">{{ $t('product.weight') }}:</span>
          <span>{{ product.weight }}</span>
        </div>

        <!-- 库存状态 -->
        <div class="stock-info">
          <span class="label">{{ $t('product.stock') }}:</span>
          <span :class="product.stock > 0 ? 'in-stock' : 'out-stock'">
            {{ product.stock > 0 ? $t('product.inStock') : $t('product.outOfStock') }}
          </span>
        </div>

        <!-- 操作按钮 -->
        <div class="actions">
          <button class="btn-primary">{{ $t('product.addToCart') }}</button>
          <button class="btn-secondary">{{ $t('product.buyNow') }}</button>
        </div>
      </div>

      <!-- 产品描述 -->
      <div class="product-description">
        <h2>{{ $t('product.detail') }}</h2>
        <p>{{ product.description }}</p>
      </div>

      <!-- 成分说明 -->
      <div v-if="product.ingredients" class="product-ingredients">
        <h2>{{ $t('product.ingredients') }}</h2>
        <p>{{ product.ingredients }}</p>
      </div>

      <!-- 使用方法 -->
      <div v-if="product.usage" class="product-usage">
        <h2>{{ $t('product.usage') }}</h2>
        <p>{{ product.usage }}</p>
      </div>

      <!-- 产品特点 -->
      <div v-if="product.features && product.features.length" class="product-features">
        <h2>{{ $t('product.features') }}</h2>
        <ul>
          <li v-for="(feature, index) in product.features" :key="index">
            {{ feature }}
          </li>
        </ul>
      </div>
    </div>

    <!-- 无数据 -->
    <div v-else class="no-data">
      {{ $t('common.noData') }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import request from '@/utils/request'
import LanguageSwitcher from '@/components/LanguageSwitcher.vue'

const route = useRoute()
const { t } = useI18n()

const loading = ref(true)
const product = ref(null)

// 获取产品详情
const fetchProductDetail = async () => {
  try {
    loading.value = true
    const res = await request.get(`/products/${route.params.id}`)
    
    // 后端会根据Accept-Language头返回对应语言的数据
    product.value = res.data
  } catch (error) {
    console.error('获取产品详情失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProductDetail()
})
</script>

<style scoped>
.product-detail {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 20px;
}

.loading,
.no-data {
  text-align: center;
  padding: 40px;
  font-size: 18px;
  color: #666;
}

.content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
}

.product-image img {
  width: 100%;
  border-radius: 8px;
}

.product-info h1 {
  font-size: 32px;
  margin-bottom: 10px;
}

.subtitle {
  color: #666;
  font-size: 16px;
  margin-bottom: 20px;
}

.price-info,
.weight-info,
.stock-info {
  margin: 15px 0;
  font-size: 16px;
}

.label {
  font-weight: bold;
  margin-right: 10px;
}

.price {
  color: #e74c3c;
  font-size: 24px;
  font-weight: bold;
}

.in-stock {
  color: #27ae60;
}

.out-stock {
  color: #e74c3c;
}

.actions {
  margin-top: 30px;
  display: flex;
  gap: 15px;
}

.btn-primary,
.btn-secondary {
  padding: 12px 30px;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover {
  background-color: #2980b9;
}

.btn-secondary {
  background-color: #95a5a6;
  color: white;
}

.btn-secondary:hover {
  background-color: #7f8c8d;
}

.product-description,
.product-ingredients,
.product-usage,
.product-features {
  grid-column: 1 / -1;
  margin-top: 40px;
}

.product-description h2,
.product-ingredients h2,
.product-usage h2,
.product-features h2 {
  font-size: 24px;
  margin-bottom: 15px;
  border-bottom: 2px solid #3498db;
  padding-bottom: 10px;
}

.product-features ul {
  list-style: none;
  padding: 0;
}

.product-features li {
  padding: 10px 0;
  padding-left: 25px;
  position: relative;
}

.product-features li:before {
  content: "✓";
  position: absolute;
  left: 0;
  color: #27ae60;
  font-weight: bold;
}
</style>
