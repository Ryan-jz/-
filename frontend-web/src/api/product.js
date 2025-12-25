import request from '@/utils/request'

/**
 * 获取产品分类列表
 * @param {Object} params - 查询参数
 * @returns {Promise}
 */
export function getCategoryList(params) {
  return request({
    url: '/product/category/list',
    method: 'get',
    params
  })
}

/**
 * 获取产品列表
 * @param {Object} params - 查询参数
 * @param {number} params.categoryId - 分类ID
 * @param {string} params.keyword - 关键词
 * @param {number} params.status - 状态
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 * @returns {Promise}
 */
export function getProductList(params) {
  return request({
    url: '/product/list',
    method: 'get',
    params
  })
}

/**
 * 获取产品详情
 * @param {number} id - 产品ID
 * @returns {Promise}
 */
export function getProductDetail(id) {
  return request({
    url: `/product/detail/${id}`,
    method: 'get'
  })
}
