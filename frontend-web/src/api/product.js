import request from '@/utils/request'

export function getCategoryList(params) {
  return request({
    url: '/product/category/list',
    method: 'get',
    params
  })
}

export function getProductList(params) {
  return request({
    url: '/product/list',
    method: 'get',
    params
  })
}

export function getProductDetail(id) {
  return request({
    url: `/product/detail/${id}`,
    method: 'get'
  })
}

export function getCategoryWithProducts(params) {
  return request({
    url: '/product/category/with-products',
    method: 'get',
    params
  })
}
