<template>
  <nav class="primary-navigation" :class="{ affixed: isNavFixed }">
    <div class="container">
      <button class="menu-toggle" @click="mobileMenuOpen = !mobileMenuOpen">
        <span></span>
        <span></span>
        <span></span>
      </button>
      
      <ul class="nav-menu" :class="{ open: mobileMenuOpen }">
        <li class="nav-item has-dropdown" :class="{ active: productMenuOpen }">
          <span class="nav-link" @click="toggleProductMenu">
            产品
            <span class="arrow">▼</span>
          </span>
          <div class="dropdown-menu-wrapper">
            <div class="dropdown-menu">
              <div class="dropdown-container">
                <div class="dropdown-content">
                  <div 
                    v-for="category in categories" 
                    :key="category.id" 
                    class="dropdown-column"
                  >
                    <router-link :to="`/product?category=${category.id}`" @click="closeMobileMenu" v-if="Array.isArray(category?.products) &&category?.products?.length>0">
                      <h4 class="dropdown-title">{{ category.name }}</h4>
                    </router-link>
                    <ul v-if="Array.isArray(category?.products) && category?.products?.length>0">
                      <li v-for="product in category.products" :key="product.id">
                        <router-link :to="`/product/${product.id}`" @click="closeMobileMenu">
                          {{ product.name }}
                        </router-link>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </li>
        <li class="nav-item">
           <router-link to="/recipe" @click="closeMobileMenu">食谱</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/nachhaltigkeit" @click="closeMobileMenu">可持续性</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/brand" @click="closeMobileMenu">品牌</router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { getCategoryWithProducts } from '@/api/product'

defineProps({
  isNavFixed: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['scroll-to-section'])

const mobileMenuOpen = ref(false)
const productMenuOpen = ref(false)
const categories = ref([])

const toggleProductMenu = () => {
  if (window.innerWidth <= 768) {
    productMenuOpen.value = !productMenuOpen.value
  }
}

const closeMobileMenu = () => {
  mobileMenuOpen.value = false
  productMenuOpen.value = false
}

const loadCategories = async () => {
  try {
    const res = await getCategoryWithProducts({ status: 1 })
    if (res.code === 0) {
      categories.value = res.data.list
    }
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
}

const handleScrollToSection = (section) => {
  mobileMenuOpen.value = false
  emit('scroll-to-section', section)
}

const handleClickOutside = (event) => {
  const nav = event.target.closest('.primary-navigation')
  if (!nav && mobileMenuOpen.value) {
    mobileMenuOpen.value = false
    productMenuOpen.value = false
  }
}

onMounted(() => {
  loadCategories()
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.primary-navigation {
  position: relative;
  z-index: 999;
  transition: all 0.3s;
  height: 68px;
  border-top: 2px solid #fff;
  border-bottom: 3px solid #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(90deg, #92121B 0%, #D5061C 25%, #D5061C 75%, #92121B 100%);
  width: 100%;
}

.primary-navigation.affixed {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.menu-toggle {
  display: none;
  flex-direction: column;
  gap: 4px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 10px;
}

.menu-toggle span {
  width: 25px;
  height: 3px;
  background: #fff;
  transition: all 0.3s ease;
}

.nav-menu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 40px;
  flex: 1;
  justify-content: center;
}

.nav-item {
  position: relative;
}

.nav-item > a,
.nav-link {
  display: block;
  padding: 20px 0;
  color: #fff;
  text-decoration: none;
  font-weight: 500;
  cursor: pointer;
  transition: color 0.3s ease;
}

.nav-item > a:hover,
.nav-link:hover {
  opacity: 0.8;
}

.arrow {
  font-size: 12px;
  margin-left: 5px;
  display: inline-block;
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.has-dropdown:hover .arrow {
  transform: rotate(180deg);
}

.dropdown-menu-wrapper {
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  opacity: 0;
  visibility: hidden;
  padding-top: 10px;
  pointer-events: none;
  z-index: 9999999;
}

.dropdown-menu {
  background: linear-gradient(90deg, #92121B 0%, #D5061C 25%, #D5061C 75%, #92121B 100%);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
  overflow: hidden;
  transform: translateY(-20px);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.has-dropdown:hover .dropdown-menu-wrapper {
  opacity: 1;
  visibility: visible;
  pointer-events: auto;
}

.has-dropdown:hover .dropdown-menu {
  transform: translateY(0);
}

.dropdown-container {
  padding: 30px;
}

.dropdown-content {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 40px;
  min-width: 800px;
}

.dropdown-column {
  min-width: 150px;
  opacity: 0;
  transform: translateY(10px);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.dropdown-column:nth-child(1) {
  transition-delay: 0.05s;
}

.dropdown-column:nth-child(2) {
  transition-delay: 0.1s;
}

.dropdown-column:nth-child(3) {
  transition-delay: 0.15s;
}

.dropdown-column:nth-child(4) {
  transition-delay: 0.2s;
}

.has-dropdown:hover .dropdown-column {
  opacity: 1;
  transform: translateY(0);
}

.dropdown-title {
  font-size: 16px;
  font-weight: 600;
  color: #fff;
  margin: 0 0 15px 0;
  padding-bottom: 10px;
  border-bottom: 2px solid #fff;
}

.dropdown-column ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.dropdown-column li {
  margin-bottom: 10px;
}

.dropdown-column a {
  color: #fff;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s ease;
  display: block;
  padding: 5px 0;
}

.dropdown-column a:hover {
  color: rgba(255, 255, 255, 0.7);
}

@media (max-width: 768px) {
  .menu-toggle {
    display: flex;
  }

  .nav-menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: linear-gradient(90deg, #92121B 0%, #D5061C 25%, #D5061C 75%, #92121B 100%);
    flex-direction: column;
    gap: 0;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
    max-height: 0;
    overflow: hidden;
    transition: max-height 0.3s ease;
    width: 100%;
    z-index: 9999;
  }

  .nav-menu.open {
    max-height: 600px;
  }

  .nav-item > a,
  .nav-link {
    padding: 15px 20px;
    border-bottom: 1px solid #fff;
    color: #fff;
  }

  .has-dropdown .dropdown-menu-wrapper {
    position: static;
    transform: none;
    padding: 0;
    max-height: 0;
    overflow: hidden;
    transition: max-height 0.3s ease;
    opacity: 1;
    visibility: visible;
  }

  .has-dropdown.active .dropdown-menu-wrapper {
    max-height: 2000px;
  }

  .dropdown-menu {
    box-shadow: none;
    border-radius: 0;
    transform: none;
    background: linear-gradient(90deg, #92121B 0%, #D5061C 25%, #D5061C 75%, #92121B 100%);
  }

  .dropdown-container {
    padding: 0;
    padding-top: 10px;
    max-height: 500px;
    overflow-y: auto;
  }

  .dropdown-content {
    grid-template-columns: 1fr;
    gap: 0;
    min-width: auto;
  }

  .dropdown-column {
    padding: 15px 20px;
    opacity: 1;
    transform: none;
    border-bottom: 1px solid #fff;
  }

  .dropdown-title {
    font-size: 14px;
    margin-bottom: 10px;
    color: #fff;
  }
}
</style>
