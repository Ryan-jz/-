import axios from 'axios'
import { getLocale } from '@/locales'

// 创建axios实例
const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api/v1',
  timeout: 15000
})

console.log('API Base URL:', import.meta.env.VITE_API_BASE_URL)
console.log('Environment:', import.meta.env.MODE)

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 添加语言参数到请求头
    const locale = getLocale()
    config.headers['Accept-Language'] = locale
    
    // 也可以作为查询参数（可选）
    // config.params = { ...config.params, lang: locale }
    
    // 可以在这里添加token等
    return config
  },
  error => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    
    // 根据后端返回的数据结构调整
    if (res.code !== undefined && res.code !== 0) {
      console.error('接口错误:', res.message || 'Error')
      return Promise.reject(new Error(res.message || 'Error'))
    }
    
    return res
  },
  error => {
    console.error('响应错误:', error)
    return Promise.reject(error)
  }
)

export default service
