<template>
  <main class="brand-page" data-brand-page>
    <PageHeader />

    <BrandHeroPanels />

    <section id="royal" class="animate-section">
      <RoyalStorySection />
    </section>
    <section id="scene" class="animate-section">
      <SceneSection />
    </section>
    <section id="quality" class="animate-section">
      <QualitySection />
    </section>
    <section id="qa" class="animate-section">
      <QASection />
    </section>
    <section id="timeline" class="animate-section">
      <TimelineSection />
    </section>
    <section id="sustain" class="animate-section">
      <SustainabilitySection />
    </section>

    <button
      class="back-to-top"
      type="button"
      :class="{ visible: showBackToTop }"
      aria-label="回到顶部"
      @click="scrollToTop"
    >
      ↑
    </button>
  </main>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import PageHeader from '@/components/PageHeader.vue'
import BrandHeroPanels from '@/components/brand/BrandHeroPanels.vue'
import RoyalStorySection from '@/components/brand/RoyalStorySection.vue'
import SceneSection from '@/components/brand/SceneSection.vue'
import QualitySection from '@/components/brand/QualitySection.vue'
import QASection from '@/components/brand/QASection.vue'
import TimelineSection from '@/components/brand/TimelineSection.vue'
import SustainabilitySection from '@/components/brand/SustainabilitySection.vue'

const showBackToTop = ref(false)
let sectionObserver = null

const handleChrome = () => {
  const y = window.scrollY || 0
  showBackToTop.value = y > 500
}

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const setupSectionObserver = () => {
  const sections = document.querySelectorAll('.animate-section')
  if (!sections.length) return

  // 兼容极老浏览器：退化为直接全部展示
  if (!('IntersectionObserver' in window)) {
    sections.forEach((s) => s.classList.add('visible'))
    return
  }

  sectionObserver = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible')
          sectionObserver?.unobserve(entry.target)
        }
      })
    },
    { root: null, threshold: 0.12, rootMargin: '0px 0px -12% 0px' }
  )

  sections.forEach((s) => sectionObserver.observe(s))
}

onMounted(() => {
  window.addEventListener('scroll', handleChrome, { passive: true })
  handleChrome()
  setupSectionObserver()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleChrome)
  sectionObserver?.disconnect()
  sectionObserver = null
})
</script>

<style scoped lang="scss">
.brand-page {
  /* 设计规范：色彩 / 圆角 / 阴影 / 间距（页面级变量，供本页承载） */
  --gold: #d4af37;
  --snow: #f8f9fa;
  --rock: #6c757d;
  --blue: #e6f3ff;
  --green: #28a745;

  --radius-8: 8px;
  --radius-12: 12px;

  --shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  --shadow-hover: 0 8px 20px rgba(0, 0, 0, 0.12);

  --space-24: 24px;
  --block-gap: 64px;

  background: #fff;
  overflow-x: hidden;
  padding-top: 0;
}

.animate-section {
  opacity: 0;
  transform: translateY(50px);
  transition: opacity 0.8s ease, transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}

.back-to-top {
  position: fixed;
  right: 24px;
  bottom: 24px;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.25);
  background: rgba(212, 175, 55, 0.92);
  color: #fff;
  font-size: 22px;
  cursor: pointer;
  box-shadow: var(--shadow);
  opacity: 0;
  pointer-events: none;
  transform: translateY(10px);
  transition: opacity 0.25s ease, transform 0.25s ease, box-shadow 0.25s ease;

  &.visible {
    opacity: 1;
    pointer-events: auto;
    transform: translateY(0);
  }

  &:hover {
    box-shadow: var(--shadow-hover);
  }
}

@media (prefers-reduced-motion: reduce) {
  .animate-section {
    transition: none;
    opacity: 1;
    transform: none;
  }

  .brand-header,
  .back-to-top {
    transition: none;
  }
}
</style>