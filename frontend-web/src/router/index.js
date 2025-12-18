/**
 * 前台路由配置
 */
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/home/index.vue'),
    meta: { title: '首页' }
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/about/index.vue'),
    meta: { title: '关于我们' }
  },
  {
    path: '/contact',
    name: 'Contact',
    component: () => import('@/views/contact/index.vue'),
    meta: { title: '联系我们' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - GF Admin` : 'GF Admin'
  next()
})

export default router
