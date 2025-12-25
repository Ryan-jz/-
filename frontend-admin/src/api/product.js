import request from '@/utils/request'

// 获取产品分类列表
export function getCategoryList(params) {
  return request({
    url: '/v1/product/category/list',
    method: 'get',
    params
  })
}

// 创建产品分类
export function createCategory(data) {
  return request({
    url: '/v1/product/category/create',
    method: 'post',
    data
  })
}

// 更新产品分类
export function updateCategory(data) {
  return request({
    url: '/v1/product/category/update',
    method: 'put',
    data
  })
}

// 删除产品分类
export function deleteCategory(id) {
  return request({
    url: `/v1/product/category/delete/${id}`,
    method: 'delete'
  })
}

// 获取产品列表
export function getProductList(params) {
  return request({
    url: '/v1/product/list',
    method: 'get',
    params
  })
}

// 获取产品详情
export function getProductDetail(id) {
  return request({
    url: `/v1/product/detail/${id}`,
    method: 'get'
  })
}

// 创建产品
export function createProduct(data) {
  return request({
    url: '/v1/product/create',
    method: 'post',
    data
  })
}

// 更新产品
export function updateProduct(data) {
  return request({
    url: '/v1/product/update',
    method: 'put',
    data
  })
}

// 删除产品
export function deleteProduct(id) {
  return request({
    url: `/v1/product/delete/${id}`,
    method: 'delete'
  })
}
