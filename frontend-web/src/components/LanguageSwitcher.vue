<template>
  <div class="language-switcher">
    <select 
      v-model="currentLocale" 
      @change="handleChange"
      class="language-select"
    >
      <option value="zh-CN">中文</option>
      <option value="en-US">English</option>
      <option value="de-DE">Deutsch</option>
    </select>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { setLocale, getLocale } from '@/locales'
import { useRouter } from 'vue-router'

const { locale } = useI18n()
const router = useRouter()
const currentLocale = ref(getLocale())

onMounted(() => {
  currentLocale.value = getLocale()
})

const handleChange = () => {
  setLocale(currentLocale.value)
  locale.value = currentLocale.value
  
  // 刷新当前页面数据
  router.go(0)
}
</script>

<style scoped>
.language-switcher {
  display: inline-block;
}

.language-select {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: white;
  cursor: pointer;
  font-size: 14px;
  outline: none;
  transition: all 0.3s;
}

.language-select:hover {
  border-color: #409eff;
}

.language-select:focus {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
}
</style>
