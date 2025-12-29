/**
 * Axios 请求封装
 * 统一处理请求和响应，包括请求拦截、响应拦截、错误处理
 */
import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import router from '@/router'

// 根据环境获取 API 基础路径
const getBaseURL = () => {
  // 生产环境使用公网 IP
  if (import.meta.env.MODE === 'production') {
    return 'http://8.133.175.112:8000/api'
  }
  // 开发环境使用代理
  return '/api'
}

// 创建 axios 实例
const service = axios.create({
  baseURL: getBaseURL(),     // API 基础路径（根据环境自动切换）
  timeout: 10000,            // 请求超时时间
  headers: {
    'Content-Type': 'application/json;charset=UTF-8'
  }
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    
    // 如果存在 token，在请求头中添加 Authorization
    if (userStore.token) {
      config.headers['Authorization'] = `Bearer ${userStore.token}`
    }
    
    // 添加语言参数到请求头（默认中文，管理后台主要用于编辑多语言内容）
    config.headers['Accept-Language'] = 'zh-CN'
    
    return config
  },
  error => {
    // 请求错误处理
    console.error('请求错误：', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    
    // 根据业务状态码判断请求是否成功
    // 后端返回 code: 0 表示成功
    if (res.code === 0) {
      return res
    } else {
      // 业务错误处理
      ElMessage.error(res.message || '请求失败')
      return Promise.reject(new Error(res.message || '请求失败'))
    }
  },
  error => {
    console.error('响应错误：', error)
    
    // HTTP 状态码错误处理
    if (error.response) {
      const status = error.response.status
      
      switch (status) {
        case 401:
          // 未授权，清除 token 并跳转到登录页
          ElMessage.error('登录已过期，请重新登录')
          const userStore = useUserStore()
          userStore.resetToken()
          router.push('/login')
          break
        case 403:
          ElMessage.error('没有权限访问')
          break
        case 404:
          ElMessage.error('请求的资源不存在')
          break
        case 500:
          ElMessage.error('服务器错误')
          break
        default:
          ElMessage.error(error.response.data.message || '请求失败')
      }
    } else if (error.request) {
      // 请求已发出但没有收到响应
      ElMessage.error('网络错误，请检查网络连接')
    } else {
      // 其他错误
      ElMessage.error(error.message || '请求失败')
    }
    
    return Promise.reject(error)
  }
)

export default service
