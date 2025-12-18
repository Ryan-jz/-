/**
 * 用户状态管理
 * 管理用户登录状态、用户信息和权限
 */
import { defineStore } from 'pinia'
import { login, getInfo } from '@/api/auth'
import router from '@/router'

export const useUserStore = defineStore('user', {
  // 状态定义
  state: () => ({
    token: localStorage.getItem('token') || '',      // JWT Token
    userInfo: null,                                   // 用户信息
    roles: [],                                        // 用户角色
    permissions: []                                   // 用户权限
  }),

  // 计算属性
  getters: {
    // 是否已登录
    isLoggedIn: (state) => !!state.token,
    
    // 获取用户名
    username: (state) => state.userInfo?.userName || '',
    
    // 获取昵称
    nickname: (state) => state.userInfo?.nickName || ''
  },

  // 方法
  actions: {
    /**
     * 用户登录
     * @param {Object} loginForm - 登录表单数据
     * @returns {Promise}
     */
    async login(loginForm) {
      try {
        const res = await login(loginForm)
        
        // 保存 Token
        this.token = res.data.token
        localStorage.setItem('token', res.data.token)
        
        return res
      } catch (error) {
        throw error
      }
    },

    /**
     * 获取用户信息
     * @returns {Promise}
     */
    async getUserInfo() {
      try {
        const res = await getInfo()
        this.userInfo = res.data
        return res
      } catch (error) {
        throw error
      }
    },

    /**
     * 用户登出
     */
    logout() {
      // 清除状态
      this.token = ''
      this.userInfo = null
      this.roles = []
      this.permissions = []
      
      // 清除本地存储
      localStorage.removeItem('token')
      
      // 跳转到登录页
      router.push('/login')
    },

    /**
     * 重置 Token（用于 Token 过期等场景）
     */
    resetToken() {
      this.token = ''
      localStorage.removeItem('token')
    }
  }
})
