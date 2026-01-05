<template>
    <PageHeader></PageHeader>
  <main>
    <HeroSection />
    <div class="animate-section"><PackagingSection /></div>
    <div class="animate-section"><MissionSection /></div>
    <div class="animate-section"><PackagingGoals /></div>
    <div class="animate-section"><CurrentFutureGoals /></div>
    <div class="animate-section"><RecyclingTips /></div>
    <div class="animate-section"><PartnershipsSection /></div>
    <div class="animate-section"><NationalParkSection /></div>
    <div class="animate-section"><CertificatesSection /></div>
  </main>
</template>

<script setup>
import { onMounted, onUnmounted } from 'vue'
import HeroSection from '@/components/sustainability/HeroSection.vue'
import PackagingSection from '@/components/sustainability/PackagingSection.vue'
import MissionSection from '@/components/sustainability/MissionSection.vue'
import PackagingGoals from '@/components/sustainability/PackagingGoals.vue'
import CurrentFutureGoals from '@/components/sustainability/CurrentFutureGoals.vue'
import RecyclingTips from '@/components/sustainability/RecyclingTips.vue'
import PartnershipsSection from '@/components/sustainability/PartnershipsSection.vue'
import NationalParkSection from '@/components/sustainability/NationalParkSection.vue'
import CertificatesSection from '@/components/sustainability/CertificatesSection.vue'
import PageHeader from '@/components/PageHeader.vue'
const handleScroll = () => {
  const sections = document.querySelectorAll('.animate-section')
  sections.forEach(section => {
    const rect = section.getBoundingClientRect()
    const windowHeight = window.innerHeight
    if (rect.top < windowHeight * 0.85) {
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
main {
  background: #fff;

  scroll-behavior: smooth;
}

.animate-section {
  opacity: 0;
  transform: translateY(80px) scale(0.95);
  transition: all 1s cubic-bezier(0.4, 0, 0.2, 1);

  &.visible {
    opacity: 1;
    transform: translateY(0) scale(1);
  }

  :deep(img) {
    transition: transform 0.6s ease-out;
  }

  &.visible :deep(img) {
    transform: scale(1);
  }

  :deep(img) {
    transform: scale(0.9);
  }
}

@media (max-width: 768px) {
  main {
    padding-top: 60px;
  }

  .animate-section {
    transform: translateY(40px) scale(0.95);
  }
}
</style>
