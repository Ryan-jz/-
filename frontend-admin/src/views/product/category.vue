<template>
  <div class="category-container">
    <el-card shadow="never">
      <template #header>
        <div class="card-header">
          <span>产品分类管理</span>
          <el-button type="primary" @click="handleCreate">新增分类</el-button>
        </div>
      </template>

      <el-table
        v-loading="loading"
        :data="categoryList"
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="分类名称" min-width="150" />
        <el-table-column prop="nameEn" label="英文名称" min-width="150" />
        <el-table-column prop="slug" label="标识" width="150" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sortOrder" label="排序" width="80" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <!-- 国际化内容 -->
        <el-divider content-position="left">多语言内容</el-divider>
        <CategoryI18nEditor v-model="formData.i18n" />
        
        <el-divider content-position="left">其他设置</el-divider>
        <el-form-item label="分类图片" prop="image">
          <el-upload
            class="avatar-uploader"
            :action="uploadUrl"
            :headers="uploadHeaders"
            :show-file-list="false"
            :on-success="handleImageSuccess"
            :before-upload="beforeImageUpload"
          >
            <img v-if="formData.image" :src="formData.image" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="分类标识" prop="slug">
          <el-input v-model="formData.slug" placeholder="例如：alpine-salt" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="formData.sortOrder" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import {
  getCategoryList,
  createCategory,
  updateCategory,
  deleteCategory,
  getCategoryI18n
} from '@/api/product'
import CategoryI18nEditor from '@/components/I18nEditor/CategoryI18nEditor.vue'

const uploadUrl = import.meta.env.VITE_API_BASE_URL + '/api/upload/image'
const uploadHeaders = {
  Authorization: 'Bearer ' + localStorage.getItem('token')
}

// 数据
const loading = ref(false)
const categoryList = ref([])

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新增分类')
const formRef = ref(null)
const formData = reactive({
  id: null,
  slug: '',
  image: '',
  sortOrder: 0,
  status: 1,
  i18n: {
    'zh-CN': { name: '', description: '' },
    'en-US': { name: '', description: '' },
    'de-DE': { name: '', description: '' }
  }
})

// 表单验证规则
const formRules = {
  slug: [{ required: true, message: '请输入分类标识', trigger: 'blur' }]
}

// 获取分类列表
const loadCategoryList = async () => {
  loading.value = true
  try {
    const res = await getCategoryList()
    categoryList.value = res.data.list || []
  } catch (error) {
    ElMessage.error('获取分类列表失败')
  } finally {
    loading.value = false
  }
}

// 新增
const handleCreate = () => {
  dialogTitle.value = '新增分类'
  Object.assign(formData, {
    id: null,
    slug: '',
    image: '',
    sortOrder: 0,
    status: 1,
    i18n: {
      'zh-CN': { name: '', description: '' },
      'en-US': { name: '', description: '' },
      'de-DE': { name: '', description: '' }
    }
  })
  dialogVisible.value = true
}

// 编辑
const handleEdit = async (row) => {
  dialogTitle.value = '编辑分类'
  
  Object.assign(formData, {
    id: row.id,
    slug: row.slug,
    image: row.image || '',
    sortOrder: row.sortOrder,
    status: row.status
  })
  
  try {
    const res = await getCategoryI18n(row.id)
    if (res.code === 0 && res.data) {
      formData.i18n = res.data
    }
  } catch (error) {
    console.error('获取国际化数据失败', error)
    formData.i18n = {
      'zh-CN': { name: row.name || '', description: row.description || '' },
      'en-US': { name: row.nameEn || '', description: row.description || '' },
      'de-DE': { name: row.name || '', description: row.description || '' }
    }
  }
  
  dialogVisible.value = true
}

const handleImageSuccess = (response) => {
  if (response.code === 0) {
    formData.image = response.data.url
    ElMessage.success('上传成功')
  } else {
    ElMessage.error(response.message || '上传失败')
  }
}

const beforeImageUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt5M = file.size / 1024 / 1024 < 5

  if (!isImage) {
    ElMessage.error('只能上传图片文件')
    return false
  }
  if (!isLt5M) {
    ElMessage.error('图片大小不能超过 5MB')
    return false
  }
  return true
}

// 提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    try {
      if (formData.id) {
        await updateCategory(formData)
        ElMessage.success('更新成功')
      } else {
        await createCategory(formData)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadCategoryList()
    } catch (error) {
      ElMessage.error(error.message || '操作失败')
    }
  })
}

// 删除
const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除分类"${row.name}"吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteCategory(row.id)
      ElMessage.success('删除成功')
      loadCategoryList()
    } catch (error) {
      ElMessage.error(error.message || '删除失败')
    }
  })
}

// 初始化
onMounted(() => {
  loadCategoryList()
})
</script>

<style lang="scss" scoped>
.category-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.avatar-uploader {
  :deep(.el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }

  :deep(.el-upload:hover) {
    border-color: var(--el-color-primary);
  }
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
  line-height: 178px;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
  object-fit: cover;
}
</style>
