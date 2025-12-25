import request from '@/utils/request'

// 获取产品分类列表
export function getCategoryList(params) {
  return request({
    url: '/product/category/list',
    method: 'get',
    params
  })
}

// 获取产品列表
export function getProductList(params) {
  return request({
    url: '/product/list',
    method: 'get',
    params
  })
}

// 获取产品详情
export function getProductDetail(id) {
  return request({
    url: `/product/detail/${id}`,
    method: 'get'
  })
}

// 创建产品
export function createProduct(data) {
  return request({
    url: '/product/create',
    method: 'post',
    data
  })
}

// 更新产品
export function updateProduct(data) {
  return request({
    url: '/product/update',
    method: 'put',
    data
  })
}

// 删除产品
export function deleteProduct(id) {
  return request({
    url: `/product/delete/${id}`,
    method: 'delete'
  })
}
