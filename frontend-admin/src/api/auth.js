/**
 * 认证相关 API
 * 包括登录、获取用户信息等接口
 */
import request from '@/utils/request'

/**
 * 用户登录
 * @param {Object} data - 登录表单数据
 * @param {string} data.username - 用户名
 * @param {string} data.password - 密码
 * @returns {Promise}
 */
export function login(data) {
  return request({
    url: '/v1/login',
    method: 'post',
    data
  })
}

/**
 * 获取当前登录用户信息
 * @returns {Promise}
 */
export function getInfo() {
  return request({
    url: '/getInfo',
    method: 'get'
  })
}

/**
 * 用户登出
 * @returns {Promise}
 */
export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}
