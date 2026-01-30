<template>
  <main class="brand-page" :class="{ 'lock-scroll': isCarouselActive }">
    <HeroSection @carousel-complete="isCarouselActive = false" />
    <div class="animate-section"><RoyalStorySection /></div>
    <div class="animate-section"><SceneSection /></div>
    <div class="animate-section"><QualitySection /></div>
    <div class="animate-section"><QASection /></div>
    <div class="animate-section"><TimelineSection /></div>
    <div class="animate-section"><SustainabilitySection /></div>
  </main>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import HeroSection from '@/components/brand/HeroSection.vue'
import RoyalStorySection from '@/components/brand/RoyalStorySection.vue'
import SceneSection from '@/components/brand/SceneSection.vue'
import QualitySection from '@/components/brand/QualitySection.vue'
import QASection from '@/components/brand/QASection.vue'
import TimelineSection from '@/components/brand/TimelineSection.vue'
import SustainabilitySection from '@/components/brand/SustainabilitySection.vue'

const isCarouselActive = ref(true)

const handleScroll = () => {
  const sections = document.querySelectorAll('.animate-section')
  sections.forEach(section => {
    const rect = section.getBoundingClientRect()
    if (rect.top < window.innerHeight * 0.85) {
      section.classList.add('visible')
    }
  })
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  handleScroll()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped lang="scss">
.brand-page {
  background: #fff;
  overflow-x: hidden;
  padding-top: 0;

  &.lock-scroll {
    overflow: hidden;
    height: 100vh;
  }
}

.animate-section {
  opacity: 0;
  transform: translateY(80px);
  transition: all 1s cubic-bezier(0.4, 0, 0.2, 1);

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>