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

export function createRecipe(data) {
  return request({
    url: '/v1/recipe/create',
    method: 'post',
    data
  })
}

export function updateRecipe(data) {
  return request({
    url: '/v1/recipe/update',
    method: 'put',
    data
  })
}

export function deleteRecipe(id) {
  return request({
    url: `/v1/recipe/delete/${id}`,
    method: 'delete'
  })
}

export function getTagList(params) {
  return request({
    url: '/v1/recipe/tag/list',
    method: 'get',
    params
  })
}

export function getAllTags() {
  return request({
    url: '/v1/recipe/tag/all',
    method: 'get'
  })
}

export function createTag(data) {
  return request({
    url: '/v1/recipe/tag/create',
    method: 'post',
    data
  })
}

export function updateTag(data) {
  return request({
    url: '/v1/recipe/tag/update',
    method: 'put',
    data
  })
}

export function deleteTag(id) {
  return request({
    url: `/v1/recipe/tag/delete/${id}`,
    method: 'delete'
  })
}
