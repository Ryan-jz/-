/**
 * 用户管理 API
 * 用户的增删改查接口
 */
import request from '@/utils/request'

/**
 * 获取用户列表
 * @param {Object} params - 查询参数
 * @returns {Promise}
 */
export function getUserList(params) {
  return request({
    url: '/admin/user/list',
    method: 'get',
    params
  })
}

/**
 * 获取用户详情
 * @param {number} userId - 用户ID
 * @returns {Promise}
 */
export function getUserDetail(userId) {
  return request({
    url: '/admin/user/detail',
    method: 'get',
    params: { userId }
  })
}

/**
 * 创建用户
 * @param {Object} data - 用户数据
 * @returns {Promise}
 */
export function createUser(data) {
  return request({
    url: '/admin/user/create',
    method: 'post',
    data
  })
}

/**
 * 更新用户
 * @param {Object} data - 用户数据
 * @returns {Promise}
 */
export function updateUser(data) {
  return request({
    url: '/admin/user/update',
    method: 'put',
    data
  })
}

/**
 * 删除用户
 * @param {number} userId - 用户ID
 * @returns {Promise}
 */
export function deleteUser(userId) {
  return request({
    url: '/admin/user/delete',
    method: 'delete',
    data: { userId }
  })
}
