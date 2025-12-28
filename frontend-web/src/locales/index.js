import { createI18n } from 'vue-i18n'
import zhCN from './zh-CN'
import enUS from './en-US'
import deDE from './de-DE'

// 从localStorage获取保存的语言，如果没有则使用浏览器语言或默认中文
function getDefaultLocale() {
  const savedLocale = localStorage.getItem('locale')
  if (savedLocale) {
    return savedLocale
  }
  
  // 获取浏览器语言
  const browserLang = navigator.language || navigator.userLanguage
  
  // 匹配支持的语言
  if (browserLang.startsWith('zh')) {
    return 'zh-CN'
  } else if (browserLang.startsWith('en')) {
    return 'en-US'
  } else if (browserLang.startsWith('de')) {
    return 'de-DE'
  }
  
  // 默认中文
  return 'zh-CN'
}

const i18n = createI18n({
  legacy: false, // 使用 Composition API 模式
  locale: getDefaultLocale(),
  fallbackLocale: 'zh-CN', // 回退语言
  messages: {
    'zh-CN': zhCN,
    'en-US': enUS,
    'de-DE': deDE,
  },
  // 全局注入 $t 函数
  globalInjection: true,
})

// 导出切换语言的函数
export function setLocale(locale) {
  i18n.global.locale.value = locale
  localStorage.setItem('locale', locale)
  // 更新HTML lang属性
  document.querySelector('html').setAttribute('lang', locale)
}

// 导出获取当前语言的函数
export function getLocale() {
  return i18n.global.locale.value
}

export default i18n
