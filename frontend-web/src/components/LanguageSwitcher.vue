<template>
  <div class="language-switcher">
    <button class="language-button" @click="toggleDropdown">
      <span class="flag-icon">{{ currentLanguage.flag }}</span>
      <span class="language-text">{{ currentLanguage.label }}</span>
      <span class="arrow" :class="{ open: isOpen }">▼</span>
    </button>
    
    <transition name="dropdown">
      <div v-if="isOpen" class="language-dropdown">
        <button
          v-for="lang in languages"
          :key="lang.code"
          class="language-option"
          :class="{ active: currentLocale === lang.code }"
          @click="changeLanguage(lang.code)"
        >
          <span class="flag-icon">{{ lang.flag }}</span>
          <span class="language-name">{{ lang.label }}</span>
          <span v-if="currentLocale === lang.code" class="check-icon">✓</span>
        </button>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { setLocale } from '@/locales'

const { locale } = useI18n()
const isOpen = ref(false)

const languages = [
  { code: 'zh-CN', label: '中文', flag: '🇨🇳' },
  // { code: 'en-US', label: 'English', flag: '🇺🇸' },
  { code: 'de-DE', label: 'Deutsch', flag: '🇩🇪' }
]

const currentLocale = computed(() => locale.value)

const currentLanguage = computed(() => {
  return languages.find(lang => lang.code === currentLocale.value) || languages[0]
})

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const changeLanguage = (langCode) => {
  setLocale(langCode)
  isOpen.value = false
  
  // 刷新页面以重新加载数据
  window.location.reload()
}

// 点击外部关闭下拉菜单
const handleClickOutside = (event) => {
  const switcher = event.target.closest('.language-switcher')
  if (!switcher && isOpen.value) {
    isOpen.value = false
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
.language-switcher {
  position: relative;
  display: inline-block;
}

.language-button {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  background: transparent;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 6px;
  color: #fff;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s ease;
}

.language-button:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.5);
}

.flag-icon {
  font-size: 18px;
  line-height: 1;
}

.language-text {
  font-weight: 500;
}

.arrow {
  font-size: 10px;
  transition: transform 0.3s ease;
}

.arrow.open {
  transform: rotate(180deg);
}

.language-dropdown {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  background: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  overflow: hidden;
  min-width: 160px;
  z-index: 1000;
}

.language-option {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  padding: 12px 16px;
  background: white;
  border: none;
  color: #333;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.2s ease;
  text-align: left;
}

.language-option:hover {
  background: #f5f5f5;
}

.language-option.active {
  background: #fff5f5;
  color: #e31e24;
}

.language-name {
  flex: 1;
  font-weight: 500;
}

.check-icon {
  color: #e31e24;
  font-weight: bold;
}

/* 下拉动画 */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.3s ease;
}

.dropdown-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* 响应式 */
@media (max-width: 768px) {
  .language-button {
    padding: 6px 12px;
    font-size: 13px;
  }
  
  .language-text {
    display: none;
  }
  
  .language-dropdown {
    right: auto;
    left: 0;
  }
}
</style>
