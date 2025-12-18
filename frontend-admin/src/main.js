/**
 * 应用程序入口文件
 * 负责初始化 Vue 应用、注册插件和挂载应用
 */
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import './styles/index.scss'

// 创建 Vue 应用实例
const app = createApp(App)

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 注册插件
app.use(createPinia())      // 状态管理
app.use(router)             // 路由
app.use(ElementPlus)        // UI 组件库

// 挂载应用
app.mount('#app')
