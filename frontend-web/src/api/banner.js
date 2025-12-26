import request from '@/utils/request'

/**
 * 获取轮播图列表
 */
export function getBannerList(params) {
  return request({
    url: '/banner/list',
    method: 'get',
    params
  })
}

/**
 * 获取轮播图详情
 */
export function getBannerDetail(bannerId) {
  return request({
    url: '/banner/detail',
    method: 'get',
    params: { bannerId }
  })
}
