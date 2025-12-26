<template>
  <nav class="primary-navigation" :class="{ affixed: isNavFixed }">
    <div class="container">
      <button class="menu-toggle" @click="mobileMenuOpen = !mobileMenuOpen">
        <span></span>
        <span></span>
        <span></span>
      </button>
      
      <ul class="nav-menu" :class="{ open: mobileMenuOpen }">
        <li class="nav-item has-dropdown">
          <span class="nav-link">
            产品
            <span class="arrow">▼</span>
          </span>
          <div class="dropdown-menu-wrapper">
            <div class="dropdown-menu">
              <div class="dropdown-container">
                <div class="dropdown-content">
                  <div 
                    v-for="category in productCategories" 
                    :key="category.id" 
                    class="dropdown-column"
                  >
                    <h4 class="dropdown-title">{{ category.title }}</h4>
                    <ul>
                      <li v-for="item in category.items" :key="item.id">
                        <a 
                          :href="`#${category.anchor}`" 
                          @click="handleScrollToSection(category.anchor)"
                        >
                          {{ item.name }}
                        </a>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </li>
        <li class="nav-item">
          <a href="#sustainability" @click="handleScrollToSection('sustainability')">可持续发展</a>
        </li>
        <li class="nav-item">
          <router-link to="/about">品牌故事</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/contact">联系我们</router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  isNavFixed: {
    type: Boolean,
    default: false
  },
  productCategories: {
    type: Array,
    default: () => [
      {
        id: 'seasoned-salt',
        title: '调味盐系列',
        anchor: 'seasoned-salt',
        items: [
          { id: 1, name: 'BBQ调味盐' },
          { id: 2, name: '香草调味盐' },
          { id: 3, name: '大蒜调味盐' }
        ]
      },
      {
        id: 'alpine-salt',
        title: '阿尔卑斯盐',
        anchor: 'alpine-salt',
        items: [
          { id: 1, name: '粗盐' },
          { id: 2, name: '细盐' },
          { id: 3, name: '研磨盐' }
        ]
      },
      {
        id: 'table-salt',
        title: '食用盐',
        anchor: 'table-salt',
        items: [
          { id: 1, name: '精制食用盐' },
          { id: 2, name: '加碘盐' },
          { id: 3, name: '海盐' }
        ]
      },
      {
        id: 'specialty-salt',
        title: '特色盐',
        anchor: 'specialty-salt',
        items: [
          { id: 1, name: '喜马拉雅粉盐' },
          { id: 2, name: '黑盐' },
          { id: 3, name: '烟熏盐' }
        ]
      }
    ]
  }
})

const emit = defineEmits(['scroll-to-section'])

const mobileMenuOpen = ref(false)

const handleScrollToSection = (section) => {
  mobileMenuOpen.value = false
  emit('scroll-to-section', section)
}

// 点击外部关闭移动菜单
const handleClickOutside = (event) => {
  const nav = event.target.closest('.primary-navigation')
  if (!nav && mobileMenuOpen.value) {
    mobileMenuOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.primary-navigation {
  color: #fff;
  background: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 100;
  transition: all 0.3s ease;
}

.primary-navigation.affixed {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  position: relative;
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
  background: #333;
  transition: all 0.3s ease;
}

.nav-menu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 40px;
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
}

.dropdown-menu {
  background: white;
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
  color: #333;
  margin: 0 0 15px 0;
  padding-bottom: 10px;
  border-bottom: 2px solid #e31e24;
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
  color: #666;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s ease;
  display: block;
  padding: 5px 0;
}

.dropdown-column a:hover {
  color: #e31e24;
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
    background: white;
    flex-direction: column;
    gap: 0;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
    max-height: 0;
    overflow: hidden;
    transition: max-height 0.3s ease;
  }

  .nav-menu.open {
    max-height: 600px;
  }

  .nav-item > a,
  .nav-link {
    padding: 15px 20px;
    border-bottom: 1px solid #eee;
  }

  .dropdown-menu-wrapper {
    position: static;
    transform: none;
    opacity: 1;
    visibility: visible;
    padding: 0;
  }

  .dropdown-menu {
    box-shadow: none;
    border-radius: 0;
  }

  .dropdown-container {
    padding: 0;
  }

  .dropdown-content {
    grid-template-columns: 1fr;
    gap: 0;
    min-width: auto;
  }

  .dropdown-column {
    padding: 15px 20px;
    border-bottom: 1px solid #eee;
  }

  .dropdown-title {
    font-size: 14px;
    margin-bottom: 10px;
  }

  .has-dropdown .dropdown-menu-wrapper {
    display: none;
  }

  .has-dropdown:hover .dropdown-menu-wrapper,
  .has-dropdown.active .dropdown-menu-wrapper {
    display: block;
  }
}
</style>
