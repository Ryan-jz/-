/**
 * 前台应用入口文件
 */
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import './styles/index.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
