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
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="英文名称" prop="nameEn">
          <el-input v-model="formData.nameEn" placeholder="请输入英文名称" />
        </el-form-item>
        <el-form-item label="分类标识" prop="slug">
          <el-input v-model="formData.slug" placeholder="例如：alpine-salt" />
        </el-form-item>
        <el-form-item label="分类描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="4"
            placeholder="请输入分类描述"
          />
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
import { getCategoryList } from '@/api/product'

// 数据
const loading = ref(false)
const categoryList = ref([])

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新增分类')
const formRef = ref(null)
const formData = reactive({
  id: null,
  name: '',
  nameEn: '',
  slug: '',
  description: '',
  sortOrder: 0,
  status: 1
})

// 表单验证规则
const formRules = {
  name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
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
    name: '',
    nameEn: '',
    slug: '',
    description: '',
    sortOrder: 0,
    status: 1
  })
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row) => {
  dialogTitle.value = '编辑分类'
  Object.assign(formData, row)
  dialogVisible.value = true
}

// 提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    try {
      // 这里需要添加创建/更新分类的API
      ElMessage.success('操作成功')
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
      // 这里需要添加删除分类的API
      ElMessage.success('删除成功')
      loadCategoryList()
    } catch (error) {
      ElMessage.error('删除失败')
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
</style>
