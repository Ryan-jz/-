<template>
  <header class="header">
    <div class="top-bar">
      <div class="container">
        <router-link to="/" class="logo">
          <img src="@/assets/images/logo.png" alt="Brand Logo" />
        </router-link>
      </div>
    </div>

    <PrimaryNavigation 
      :is-nav-fixed="isNavFixed" 
      @scroll-to-section="scrollToSection"
    />
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import PrimaryNavigation from './PrimaryNavigation.vue'

const isNavFixed = ref(false)

const scrollToSection = (sectionId) => {
  const element = document.getElementById(sectionId)
  if (element) {
    const offset = 100
    const elementPosition = element.getBoundingClientRect().top
    const offsetPosition = elementPosition + window.pageYOffset - offset
    window.scrollTo({ top: offsetPosition, behavior: 'smooth' })
  }
}

const handleScroll = () => {
  isNavFixed.value = window.scrollY > 100
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style lang="scss" scoped>
.header {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url('@/assets/images/background.jpg');
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
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }

  .logo {
    height: 100%;

    img {
      height: 100%;
      display: block;
    }
  }
}

@media (max-width: 768px) {
  .top-bar {
    height: auto;
    padding: 10px 0;

    .logo {
      height: 60px;
    }
  }
}
</style>
