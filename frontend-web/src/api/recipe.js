import request from '@/utils/request'

export function getRecipeList(params) {
  return request({
    url: '/v1/recipe/list',
    method: 'get',
    params
  })
}

export function getRecipeDetail(id) {
  return request({
    url: `/v1/recipe/detail/${id}`,
    method: 'get'
  })
}
