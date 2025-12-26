import request from '@/utils/request'

/**
 * 获取轮播图列表
 */
export function getBannerList(params) {
  return request({
    url: '/v1/banner/list',
    method: 'get',
    params
  })
}

/**
 * 获取轮播图详情
 */
export function getBannerDetail(bannerId) {
  return request({
    url: '/v1/banner/detail',
    method: 'get',
    params: { bannerId }
  })
}

/**
 * 创建轮播图
 */
export function createBanner(data) {
  return request({
    url: '/v1/banner/create',
    method: 'post',
    data
  })
}

/**
 * 更新轮播图
 */
export function updateBanner(data) {
  return request({
    url: '/v1/banner/update',
    method: 'put',
    data
  })
}

/**
 * 删除轮播图
 */
export function deleteBanner(bannerId) {
  return request({
    url: '/v1/banner/delete',
    method: 'delete',
    data: { bannerId }
  })
}
