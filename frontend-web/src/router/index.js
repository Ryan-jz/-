/**
 * 前台路由配置
 */
import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/home/index.vue'),
      //  component: () => import('@/views/under-construction/index.vue'),
    meta: { title: '网站建设中' }
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
  },
  {
    path: '/recipe',
    name: 'Recipe',
    component: () => import('@/views/recipe/index.vue'),
    meta: { title: '食谱列表' }
  },
  {
    path: '/recipe/:id',
    name: 'RecipeDetail',
    component: () => import('@/views/recipe/detail.vue'),
    meta: { title: '食谱详情' }
  },
  {
    path: '/product',
    name: 'Product',
    component: () => import('@/views/product/index.vue'),
    meta: { title: '产品列表' }
  },
  {
    path: '/nachhaltigkeit',
    name: 'Nachhaltigkeit',
    component: () => import('@/views/nachhaltigkeit/index.vue'),
    meta: { title: '可持续性' }
  },
    {
    path: '/dian',
    name: 'Dian',
    component: () => import('@/views/dian/index.vue'),
    meta: { title: '可持续性' }
  },
  {
    path: '/product/:id',
    name: 'ProductDetail',
    component: () => import('@/views/product/detail.vue'),
    meta: { title: '产品详情' }
  },
  {
    path: '/brand',
    name: 'Brand',
    component: () => import('@/views/brand/index.vue'),
    meta: { title: '品牌', layout: false }
  },
    {
    path: '/alpensalze',
    name: 'Alpensalze',
    component: () => import('@/views/alpensalze/index.vue'),
    meta: { title: '关于我们' }
  },
  {
    path: '/test-unocss',
    name: 'TestUnoCSS',
    component: () => import('@/views/test-unocss.vue'),
    meta: { title: 'UnoCSS 测试' }
  },
  {
    path: '/under-construction',
    name: 'UnderConstruction',
    component: () => import('@/views/under-construction/index.vue'),
    meta: { title: '网站建设中' }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} ` : '白金盐'
  next()
})

export default router
