/**
 * 食谱相关 API
 */
import request from '@/utils/request'

/**
 * 获取食谱列表
 * @param {Object} params - 查询参数
 * @returns {Promise}
 */
export function getRecipeList(params) {
  return request({
    url: '/v1/recipe/list',
    method: 'get',
    params
  })
}

/**
 * 获取食谱详情
 * @param {number} id - 食谱ID
 * @returns {Promise}
 */
export function getRecipeDetail(id) {
  return request({
    url: `/v1/recipe/detail/${id}`,
    method: 'get'
  })
}

/**
 * 创建食谱
 * @param {Object} data - 食谱数据
 * @returns {Promise}
 */
export function createRecipe(data) {
  return request({
    url: '/v1/recipe/create',
    method: 'post',
    data
  })
}

/**
 * 更新食谱
 * @param {Object} data - 食谱数据
 * @returns {Promise}
 */
export function updateRecipe(data) {
  return request({
    url: '/v1/recipe/update',
    method: 'put',
    data
  })
}

/**
 * 删除食谱
 * @param {number} id - 食谱ID
 * @returns {Promise}
 */
export function deleteRecipe(id) {
  return request({
    url: `/v1/recipe/delete/${id}`,
    method: 'delete'
  })
}
