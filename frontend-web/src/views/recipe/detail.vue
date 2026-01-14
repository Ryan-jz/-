<template>
  <div class="recipe-detail-container">
    <header class="header">
      <div class="top-bar">
        <div class="container">
          <router-link to="/" class="logo">
            <img src="@/assets/images/logo.png" alt="Brand Logo" />
          </router-link>
        </div>
      </div>
    </header>

    <PrimaryNavigation :is-nav-fixed="isNavFixed" />

    <main class="main-content">
      <div class="container">
        <div class="recipe-detail" v-if="recipe">
          <div class="recipe-header fade-in">
            <h1>{{ recipe.name }}</h1>
            <p class="subtitle" v-if="recipe.subtitle">{{ recipe.subtitle }}</p>
          </div>

          <div class="recipe-content">
            <div class="recipe-main">
              <div class="recipe-image fade-in">
                <img :src="recipe.image" :alt="recipe.name" />
              </div>

              <div class="recipe-meta fade-in">
                <div class="meta-item">
                  <span class="label">烹饪时间</span>
                  <span class="value">{{ recipe.cookingTime }}分钟</span>
                </div>
                <div class="meta-item">
                  <span class="label">难度</span>
                  <span class="value">{{ getDifficultyText(recipe.difficulty) }}</span>
                </div>
                <div class="meta-item">
                  <span class="label">份量</span>
                  <span class="value">{{ recipe.servings }}人份</span>
                </div>
              </div>

              <div class="two-column-section">
                <div class="ingredients-section fade-in">
                  <h2>食材清单</h2>
                  <ul class="ingredients-list">
                    <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
                      <span class="name">{{ ingredient.name }}</span>
                      <span class="amount">{{ ingredient.amount }}</span>
                    </li>
                  </ul>
                </div>

                <div class="products-section fade-in" v-if="relatedProduct">
                  <h2>相关产品</h2>
                  <div 
                    class="product-card"
                    @click="goToProduct(relatedProduct.id)"
                  >
                    <img :src="relatedProduct.image" :alt="relatedProduct.name" />
                    <h4>{{ relatedProduct.name }}</h4>
                  </div>
                </div>
              </div>

              <div class="steps-section fade-in">
                <h2>制作步骤</h2>
                <div class="content" v-html="recipe.content"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import PrimaryNavigation from '@/components/PrimaryNavigation.vue'
import { getRecipeDetail } from '@/api/recipe'
import { getProductList } from '@/api/product'

const route = useRoute()
const router = useRouter()
const isNavFixed = ref(false)
const recipe = ref(null)
const relatedProduct = ref(null)

const getDifficultyText = (difficulty) => {
  const map = { 1: '简单', 2: '中等', 3: '困难' }
  return map[difficulty] || '简单'
}

const loadRecipeDetail = async () => {
  try {
    const res = await getRecipeDetail(route.params.id)
    if (res.data) {
      recipe.value = res.data
      if (recipe.value.productIds) {
        loadRelatedProducts(recipe.value.productIds)
      }
    }
  } catch (error) {
    console.error('加载食谱详情失败:', error)
  }
}

const loadRelatedProducts = async (productIds) => {
  try {
    const ids = productIds.split(',').map(id => parseInt(id))
    if (ids.length === 0) return
    
    const res = await getProductList({ status: 1, page: 1, pageSize: 100 })
    if (res.data && res.data.list) {
      const products = res.data.list.filter(p => ids.includes(p.id))
      relatedProduct.value = products[0] || null
    }
  } catch (error) {
    console.error('加载相关产品失败:', error)
  }
}

const goToProduct = (id) => {
  router.push(`/product/${id}`)
}

const observeElements = () => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible')
      }
    })
  }, { threshold: 0.1 })

  document.querySelectorAll('.fade-in').forEach(el => observer.observe(el))
}

onMounted(() => {
  loadRecipeDetail()
  setTimeout(observeElements, 100)
})
</script>

<style lang="scss" scoped>
.recipe-detail-container {
  background-image: url('@/assets/images/background.jpg');
  background-repeat: no-repeat;
  background-size: 100%;
  min-height: 100vh;
}

.fade-in {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.6s ease, transform 0.6s ease;

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}

.header {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.top-bar {
  padding: 2px 0;
  height: 112px;
  display: flex;
  align-items: center;

  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }

  .logo img {
    height: 100px;
  }
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.main-content {
  padding: 60px 0;
}

.recipe-detail {
  background: #fff;
  border-radius: 12px;
  padding: 40px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.recipe-header {
  text-align: center;
  margin-bottom: 40px;

  h1 {
    font-size: 42px;
    color: #333;
    margin: 0 0 10px 0;
  }

  .subtitle {
    font-size: 18px;
    color: #666;
  }
}

.recipe-content {
  width: 100%;
}

.recipe-main {
  .recipe-image {
    width: 100%;
    height: 400px;
    border-radius: 12px;
    overflow: hidden;
    margin-bottom: 30px;

    img {
      width: 100%;
      height: 100%;
      object-fit: fill;
    }
  }

  .recipe-meta {
    display: flex;
    gap: 30px;
    padding: 20px;
    background: #f8f8f8;
    border-radius: 8px;
    margin-bottom: 30px;

    .meta-item {
      display: flex;
      flex-direction: column;
      gap: 5px;

      .label {
        font-size: 14px;
        color: #999;
      }

      .value {
        font-size: 18px;
        color: #333;
        font-weight: 600;
      }
    }
  }

  .two-column-section {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 40px;
    margin-bottom: 40px;
  }

  .ingredients-section,
  .products-section,
  .steps-section {
    h2 {
      font-size: 28px;
      color: #333;
      margin-bottom: 20px;
      padding-bottom: 10px;
      border-bottom: 2px solid #c41e3a;
    }
  }

  .ingredients-list {
    list-style: none;
    padding: 0;
    margin: 0;

    li {
      display: flex;
      justify-content: space-between;
      padding: 12px 0;
      border-bottom: 1px solid #eee;

      .name {
        color: #333;
        font-size: 16px;
      }

      .amount {
        color: #666;
        font-size: 16px;
        font-weight: 500;
      }
    }
  }

  .product-card {
    background: #fff;
    border: 1px solid #eee;
    border-radius: 8px;
    overflow: hidden;
    cursor: pointer;
    transition: all 0.3s;
    display: flex;
    flex-direction: column;
    height: 100%;

    &:hover {
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      transform: translateY(-2px);
    }

    img {
      width: 100%;
      flex: 1;
      object-fit: fill;
    }

    h4 {
      font-size: 16px;
      color: #333;
      margin: 0;
      padding: 15px;
      text-align: center;
    }
  }

  .content {
    font-size: 16px;
    line-height: 1.8;
    color: #666;

    :deep(img) {
      max-width: 100%;
      border-radius: 8px;
      margin: 20px 0;
    }

    :deep(p) {
      margin-bottom: 15px;
    }
  }
}

@media (max-width: 768px) {
  .recipe-detail {
    padding: 20px;
  }

  .recipe-header h1 {
    font-size: 28px;
  }

  .recipe-main {
    .recipe-image {
      height: 250px;
    }

    .recipe-meta {
      flex-direction: column;
      gap: 15px;
    }

    .two-column-section {
      grid-template-columns: 1fr;
      gap: 30px;
    }
  }
}
</style>
