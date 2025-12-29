<template>
  <div class="banner-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>轮播图管理</span>
          <el-button type="primary" @click="handleAdd">新增轮播图</el-button>
        </div>
      </template>

      <!-- 搜索表单 -->
      <el-form :inline="true" :model="queryParams" class="search-form">
        <el-form-item label="位置">
          <el-select v-model="queryParams.position" placeholder="请选择位置" clearable>
            <el-option label="首页" value="home" />
            <el-option label="产品页" value="product" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="queryParams.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="停用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>

      <!-- 表格 -->
      <el-table :data="bannerList" v-loading="loading" border>
        <el-table-column label="ID" prop="bannerId" width="80" />
        <el-table-column label="预览" width="150">
          <template #default="{ row }">
            <img v-if="row.mediaType === 1" :src="row.mediaUrl" class="preview-image" />
            <video v-else :src="row.mediaUrl" class="preview-image" />
          </template>
        </el-table-column>
        <el-table-column label="标题" prop="title" />
        <el-table-column label="描述" prop="description" show-overflow-tooltip />
        <el-table-column label="类型" width="80">
          <template #default="{ row }">
            <el-tag :type="row.mediaType === 1 ? 'success' : 'warning'">
              {{ row.mediaType === 1 ? '图片' : '视频' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="位置" width="100">
          <template #default="{ row }">
            {{ row.position === 'home' ? '首页' : '产品页' }}
          </template>
        </el-table-column>
        <el-table-column label="排序" prop="sortOrder" width="80" />
        <el-table-column label="状态" width="80">
          <template #default="{ row }">
            <el-switch
              v-model="row.status"
              :active-value="1"
              :inactive-value="0"
              @change="handleStatusChange(row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        v-model:current-page="queryParams.page"
        v-model:page-size="queryParams.pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getList"
        @current-change="getList"
      />
    </el-card>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="handleDialogClose"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
      >
        <!-- 国际化内容 -->
        <el-divider content-position="left">多语言内容</el-divider>
        <BannerI18nEditor v-model="formData.i18n" />
        
        <el-divider content-position="left">媒体设置</el-divider>
        <el-form-item label="媒体类型" prop="mediaType">
          <el-radio-group v-model="formData.mediaType">
            <el-radio :label="1">图片</el-radio>
            <el-radio :label="2">视频</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="媒体文件" prop="mediaUrl">
          <el-upload
            class="upload-demo"
            :action="currentUploadUrl"
            :headers="uploadHeaders"
            :show-file-list="false"
            :on-success="handleUploadSuccess"
            :before-upload="beforeUpload"
          >
            <el-button type="primary">点击上传</el-button>
          </el-upload>
          <div v-if="formData.mediaUrl" class="preview-container">
            <img v-if="formData.mediaType === 1" :src="formData.mediaUrl" class="preview" />
            <video v-else :src="formData.mediaUrl" class="preview" controls />
          </div>
        </el-form-item>
        <el-form-item label="按钮链接" prop="buttonLink">
          <el-input v-model="formData.buttonLink" placeholder="请输入按钮链接" />
        </el-form-item>
        
        <el-divider content-position="left">其他设置</el-divider>
        <el-form-item label="位置" prop="position">
          <el-select v-model="formData.position" placeholder="请选择位置">
            <el-option label="首页" value="home" />
            <el-option label="产品页" value="product" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="formData.sortOrder" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">停用</el-radio>
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
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getBannerList, createBanner, updateBanner, deleteBanner, getBannerI18n } from '@/api/banner'
import BannerI18nEditor from '@/components/I18nEditor/BannerI18nEditor.vue'

const loading = ref(false)
const bannerList = ref([])
const total = ref(0)

const queryParams = reactive({
  position: '',
  status: null,
  page: 1,
  pageSize: 10
})

const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = reactive({
  bannerId: null,
  mediaType: 1,
  mediaUrl: '',
  buttonLink: '',
  position: 'home',
  sortOrder: 0,
  status: 1,
  i18n: {
    'zh-CN': { title: '', description: '', buttonText: '' },
    'en-US': { title: '', description: '', buttonText: '' },
    'de-DE': { title: '', description: '', buttonText: '' }
  }
})

const formRules = {
  mediaType: [{ required: true, message: '请选择媒体类型', trigger: 'change' }],
  mediaUrl: [{ required: true, message: '请上传媒体文件', trigger: 'change' }],
  position: [{ required: true, message: '请选择位置', trigger: 'change' }]
}

const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000'
const uploadHeaders = {
  Authorization: 'Bearer ' + localStorage.getItem('token')
}

// 根据媒体类型动态获取上传URL
const currentUploadUrl = computed(() => {
  return formData.mediaType === 2 
    ? `${baseUrl}/api/v1/upload/video`
    : `${baseUrl}/api/v1/upload/image`
})

// 获取列表
const getList = async () => {
  loading.value = true
  try {
    const res = await getBannerList(queryParams)
    bannerList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    ElMessage.error('获取列表失败')
  } finally {
    loading.value = false
  }
}

// 查询
const handleQuery = () => {
  queryParams.page = 1
  getList()
}

// 重置
const handleReset = () => {
  queryParams.position = ''
  queryParams.status = null
  queryParams.page = 1
  getList()
}

// 新增
const handleAdd = () => {
  dialogTitle.value = '新增轮播图'
  dialogVisible.value = true
}

// 编辑
const handleEdit = async (row) => {
  dialogTitle.value = '编辑轮播图'
  
  // 复制基本数据
  Object.assign(formData, {
    bannerId: row.bannerId,
    mediaType: row.mediaType,
    mediaUrl: row.mediaUrl,
    buttonLink: row.buttonLink,
    position: row.position,
    sortOrder: row.sortOrder,
    status: row.status
  })
  
  // 获取国际化数据
  try {
    const res = await getBannerI18n(row.bannerId)
    if (res.code === 0 && res.data) {
      formData.i18n = res.data
    }
  } catch (error) {
    console.error('获取国际化数据失败', error)
    // 如果获取失败，使用当前行的数据作为默认值
    formData.i18n = {
      'zh-CN': { title: row.title || '', description: row.description || '', buttonText: row.buttonText || '' },
      'en-US': { title: row.title || '', description: row.description || '', buttonText: row.buttonText || '' },
      'de-DE': { title: row.title || '', description: row.description || '', buttonText: row.buttonText || '' }
    }
  }
  
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该轮播图吗？', '提示', {
      type: 'warning'
    })
    await deleteBanner(row.bannerId)
    ElMessage.success('删除成功')
    getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 状态切换
const handleStatusChange = async (row) => {
  try {
    await updateBanner({
      bannerId: row.bannerId,
      mediaType: row.mediaType,
      mediaUrl: row.mediaUrl,
      buttonLink: row.buttonLink,
      position: row.position,
      sortOrder: row.sortOrder,
      status: row.status,
      i18n: row.i18n || {
        'zh-CN': { title: row.title, description: row.description, buttonText: row.buttonText },
        'en-US': { title: row.title, description: row.description, buttonText: row.buttonText },
        'de-DE': { title: row.title, description: row.description, buttonText: row.buttonText }
      }
    })
    ElMessage.success('状态更新成功')
  } catch (error) {
    ElMessage.error('状态更新失败')
    row.status = row.status === 1 ? 0 : 1
  }
}

// 上传成功
const handleUploadSuccess = (response) => {
  if (response.code === 0) {
    formData.mediaUrl = response.data.url
    ElMessage.success('上传成功')
  } else {
    ElMessage.error(response.message || '上传失败')
  }
}

// 上传前校验
const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isVideo = file.type.startsWith('video/')
  
  if (formData.mediaType === 1 && !isImage) {
    ElMessage.error('请上传图片文件')
    return false
  }
  
  if (formData.mediaType === 2 && !isVideo) {
    ElMessage.error('请上传视频文件')
    return false
  }
  
  const maxSize = formData.mediaType === 2 ? 50 : 5
  const isLtMaxSize = file.size / 1024 / 1024 < maxSize
  if (!isLtMaxSize) {
    ElMessage.error(`文件大小不能超过 ${maxSize}MB`)
    return false
  }
  
  return true
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    
    if (formData.bannerId) {
      await updateBanner(formData)
      ElMessage.success('更新成功')
    } else {
      await createBanner(formData)
      ElMessage.success('创建成功')
    }
    
    dialogVisible.value = false
    getList()
  } catch (error) {
    console.error(error)
  }
}

// 关闭对话框
const handleDialogClose = () => {
  formRef.value?.resetFields()
  Object.assign(formData, {
    bannerId: null,
    mediaType: 1,
    mediaUrl: '',
    buttonLink: '',
    position: 'home',
    sortOrder: 0,
    status: 1,
    i18n: {
      'zh-CN': { title: '', description: '', buttonText: '' },
      'en-US': { title: '', description: '', buttonText: '' },
      'de-DE': { title: '', description: '', buttonText: '' }
    }
  })
}

onMounted(() => {
  getList()
})
</script>

<style scoped>
.banner-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-form {
  margin-bottom: 20px;
}

.preview-image {
  width: 120px;
  height: 80px;
  object-fit: cover;
  border-radius: 4px;
}

.preview-container {
  margin-top: 10px;
}

.preview {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
}

.el-pagination {
  margin-top: 20px;
  justify-content: flex-end;
}
</style>
