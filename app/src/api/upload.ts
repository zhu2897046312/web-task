import request from './request'

export const uploadApi = {
  // 上传图片
  uploadImage(file: File) {
    const formData = new FormData()
    formData.append('file', file)
    return request.post<{ url: string }>('/upload/image', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
} 