/**
 * 路由配置文件
 * 定义应用的所有路由规则和导航守卫
 */
import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

// 路由配置
const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { title: '登录', requiresAuth: false }
  },
  {
    path: '/',
    component: () => import('@/layout/index.vue'),
    redirect: '/dashboard',
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '首页', icon: 'HomeFilled' }
      },
      {
        path: 'product/list',
        name: 'ProductList',
        component: () => import('@/views/product/list.vue'),
        meta: { title: '产品管理', icon: 'Goods' }
      },
      {
        path: 'product/category',
        name: 'ProductCategory',
        component: () => import('@/views/product/category.vue'),
        meta: { title: '分类管理', icon: 'Menu' }
      },
      {
        path: 'recipe/list',
        name: 'RecipeList',
        component: () => import('@/views/recipe/list.vue'),
        meta: { title: '食谱管理', icon: 'Reading' }
      },
      {
        path: 'system/user',
        name: 'User',
        component: () => import('@/views/system/user/index.vue'),
        meta: { title: '用户管理', icon: 'User' }
      },
      {
        path: 'system/role',
        name: 'Role',
        component: () => import('@/views/system/role/index.vue'),
        meta: { title: '角色管理', icon: 'UserFilled' }
      },
      {
        path: 'system/menu',
        name: 'Menu',
        component: () => import('@/views/system/menu/index.vue'),
        meta: { title: '菜单管理', icon: 'Menu' }
      }
    ]
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),
  routes
})

// 全局前置守卫 - 权限验证
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  const token = userStore.token

  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - GF Admin` : 'GF Admin'

  // 判断是否需要登录
  if (to.meta.requiresAuth !== false) {
    if (!token) {
      // 未登录，跳转到登录页
      next('/login')
    } else {
      next()
    }
  } else {
    // 已登录用户访问登录页，跳转到首页
    if (to.path === '/login' && token) {
      next('/')
    } else {
      next()
    }
  }
})

export default router
