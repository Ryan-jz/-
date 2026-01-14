import request from '@/utils/request'

export function getRecipeList(params) {
  return request({
    url: '/recipe/list',
    method: 'get',
    params
  })
}

export function getRecipeDetail(id) {
  return request({
    url: `/recipe/detail/${id}`,
    method: 'get'
  })
}
