/**
 * 文件上传相关 API
 */
import request from '@/utils/request'

/**
 * 上传图片
 * @param {File} file - 图片文件
 * @returns {Promise}
 */
export function uploadImage(file) {
  const formData = new FormData()
  formData.append('file', file)
  
  return request({
    url: '/v1/upload/image',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

/**
 * 删除文件
 * @param {string} url - 文件URL
 * @returns {Promise}
 */
export function deleteFile(url) {
  return request({
    url: '/v1/upload/delete',
    method: 'post',
    data: { url }
  })
}
