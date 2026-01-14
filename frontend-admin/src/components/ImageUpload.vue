<template>
  <div class="image-upload">
    <el-upload
      :action="uploadUrl"
      :headers="uploadHeaders"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleError"
      :before-upload="beforeUpload"
      :disabled="disabled"
      accept="image/*"
    >
      <div v-if="imageUrl" class="image-preview">
        <img :src="getImageUrl(imageUrl)" alt="预览图" />
        <div class="image-overlay">
          <el-icon class="preview-icon" @click.stop="handlePreview">
            <ZoomIn />
          </el-icon>
          <el-icon class="delete-icon" @click.stop="handleDelete">
            <Delete />
          </el-icon>
        </div>
      </div>
      <div v-else class="upload-placeholder">
        <el-icon class="upload-icon"><Plus /></el-icon>
        <div class="upload-text">{{ placeholder }}</div>
      </div>
    </el-upload>

    <!-- 图片预览对话框 -->
    <el-dialog v-model="previewVisible" title="图片预览" width="800px">
      <img :src="getImageUrl(imageUrl)" alt="预览图" style="width: 100%" />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus, ZoomIn, Delete } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: '点击上传图片'
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue'])

const userStore = useUserStore()
const imageUrl = ref(props.modelValue)
const previewVisible = ref(false)

// 上传地址
const uploadUrl = computed(() => {
  
  const baseURL = import.meta.env.VITE_API_BASE_URL+'/api'
  return `${baseURL}/v1/upload/image`
})

// 上传请求头
const uploadHeaders = computed(() => {
  return {
    Authorization: `Bearer ${userStore.token}`
  }
})

// 获取完整图片URL
const getImageUrl = (url) => {
  if (!url) return ''
  if (url.startsWith('http')) return url
  return `/api${url}`
}

// 上传前验证
const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt5M = file.size / 1024 / 1024 < 5

  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt5M) {
    ElMessage.error('图片大小不能超过 5MB!')
    return false
  }
  return true
}

// 上传成功
const handleSuccess = (response) => {
  if (response.code === 0) {
    imageUrl.value = response.data.url
    emit('update:modelValue', response.data.url)
    ElMessage.success('上传成功')
  } else {
    ElMessage.error(response.message || '上传失败')
  }
}

// 上传失败
const handleError = (error) => {
  console.error('上传失败:', error)
  ElMessage.error('上传失败，请重试')
}

// 预览图片
const handlePreview = () => {
  previewVisible.value = true
}

// 删除图片
const handleDelete = () => {
  imageUrl.value = ''
  emit('update:modelValue', '')
  ElMessage.success('已删除')
}

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  imageUrl.value = newVal
})
</script>

<style scoped lang="scss">
.image-upload {
  :deep(.el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);

    &:hover {
      border-color: var(--el-color-primary);
    }
  }

  .image-preview {
    width: 148px;
    height: 148px;
    position: relative;

    img {
      width: 100%;
      height: 100%;
      object-fit: fill;
    }

    .image-overlay {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: rgba(0, 0, 0, 0.5);
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 20px;
      opacity: 0;
      transition: opacity 0.3s;

      .el-icon {
        color: #fff;
        font-size: 20px;
        cursor: pointer;

        &:hover {
          transform: scale(1.2);
        }
      }
    }

    &:hover .image-overlay {
      opacity: 1;
    }
  }

  .upload-placeholder {
    width: 148px;
    height: 148px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    color: var(--el-text-color-secondary);

    .upload-icon {
      font-size: 28px;
      margin-bottom: 8px;
    }

    .upload-text {
      font-size: 14px;
    }
  }
}
</style>
