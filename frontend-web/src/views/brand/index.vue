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
@import url('https://fonts.googleapis.com/css2?family=Cinzel:wght@400;700&family=Lato:wght@300;400;700&family=Playfair+Display:ital,wght@0,400;0,700;1,400&display=swap');

.brand-page {
  /* 品牌设计规范：沉浸式 / 奢华感 / 极简 */
  --brand-gold: #D4AF37;
  --brand-gold-hover: #B59328;
  --brand-dark: #0F172A;
  --brand-gray: #334155;
  --brand-light: #F8F9FA;
  --brand-white: #FFFFFF;
  
  --font-serif: 'Playfair Display', 'Cinzel', serif;
  --font-sans: 'Lato', system-ui, -apple-system, sans-serif;

  --shadow-soft: 0 10px 40px -10px rgba(0, 0, 0, 0.08);
  --shadow-card: 0 20px 60px -15px rgba(0, 0, 0, 0.1);

  font-family: var(--font-sans);
  color: var(--brand-gray);
  background-color: var(--brand-white);
  overflow-x: hidden;
  padding-top: 0;

  /* 全局排版优化 - 通过 Deep Selectors 渗透到子组件 */
  :deep(h1), :deep(h2), :deep(h3), :deep(.section-title) {
    font-family: var(--font-serif) !important;
    font-weight: 700;
    letter-spacing: -0.02em;
    color: var(--brand-dark) !important;
  }

  :deep(.section-title) {
    position: relative;
    display: inline-block;
    margin-bottom: 1.5rem;
    
    &::after {
      content: '';
      display: block;
      width: 60px;
      height: 3px;
      background: var(--brand-gold);
      margin-top: 1rem;
      margin-left: auto;
      margin-right: auto; /* 默认居中，具体组件可能需要调整 */
    }
  }

  :deep(.section-subtitle) {
    font-family: var(--font-sans) !important;
    font-weight: 400;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    color: var(--brand-gold) !important;
    font-size: 0.9rem;
    margin-bottom: 1rem;
  }

  :deep(p) {
    line-height: 1.8;
    color: var(--brand-gray);
    font-size: 1.05rem;
    font-family: var(--font-sans) !important;
  }

  /* 统一 Section 间距与呼吸感 */
  :deep(section) {
    padding: 8rem 0; /* 增加留白 */
    position: relative;
  }
}

.animate-section {
  opacity: 0;
  transform: translateY(60px);
  transition: all 1.2s cubic-bezier(0.22, 1, 0.36, 1);
  will-change: opacity, transform;

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}

.back-to-top {
  position: fixed;
  right: 32px;
  bottom: 32px;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  border: 1px solid rgba(212, 175, 55, 0.3);
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  color: var(--brand-gold);
  font-family: var(--font-serif);
  font-size: 24px;
  cursor: pointer;
  box-shadow: var(--shadow-soft);
  opacity: 0;
  pointer-events: none;
  transform: translateY(20px) scale(0.9);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  z-index: 999;

  &.visible {
    opacity: 1;
    pointer-events: auto;
    transform: translateY(0) scale(1);
  }

  &:hover {
    background: var(--brand-gold);
    color: #fff;
    box-shadow: var(--shadow-card);
    transform: translateY(-4px) scale(1.05);
  }
}

@media (max-width: 768px) {
  .brand-page {
    :deep(section) {
      padding: 4rem 0;
    }
  }
}

@media (prefers-reduced-motion: reduce) {
  .animate-section, .back-to-top {
    transition: none;
    opacity: 1;
    transform: none;
  }
}
</style>