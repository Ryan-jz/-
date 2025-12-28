/**
 * 前台应用入口文件
 */
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import i18n from './locales'
import './styles/index.scss'

// UnoCSS
import 'virtual:uno.css'
import '@unocss/reset/tailwind.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(i18n)

app.mount('#app')
