<template>
  <div class="app-container">
    <el-card>
      <div class="filter-container">
        <el-input v-model="listQuery.keyword" placeholder="食谱名称" style="width: 200px;" clearable @keyup.enter="handleFilter" />
        <el-select v-model="listQuery.status" placeholder="状态" clearable style="width: 120px; margin-left: 10px;">
          <el-option label="启用" :value="1" />
          <el-option label="禁用" :value="0" />
        </el-select>
        <el-button type="primary" icon="Search" style="margin-left: 10px;" @click="handleFilter">搜索</el-button>
        <el-button type="success" icon="Plus" @click="handleCreate">新增</el-button>
        <el-button type="warning" icon="PriceTag" @click="$router.push('/recipe/tag')">标签管理</el-button>
      </div>

      <el-table :data="list" border style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="image" label="封面" width="100">
          <template #default="{ row }">
            <el-image v-if="row.image" :src="row.image" style="width: 60px; height: 60px;" fit="cover" />
          </template>
        </el-table-column>
        <el-table-column prop="name" label="食谱名称" />
        <el-table-column prop="cookingTime" label="烹饪时间" width="100">
          <template #default="{ row }">{{ row.cookingTime }}分钟</template>
        </el-table-column>
        <el-table-column prop="difficulty" label="难度" width="100">
          <template #default="{ row }">
            <el-tag :type="row.difficulty === 1 ? 'success' : row.difficulty === 2 ? 'warning' : 'danger'">
              {{ ['', '简单', '中等', '困难'][row.difficulty] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="listQuery.page"
        v-model:page-size="listQuery.pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        style="margin-top: 20px;"
        @size-change="getList"
        @current-change="getList"
      />
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogType === 'create' ? '新增食谱' : '编辑食谱'" width="800px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="食谱名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入食谱名称" />
        </el-form-item>
        <el-form-item label="副标题" prop="subtitle">
          <el-input v-model="formData.subtitle" placeholder="请输入副标题" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="formData.description" type="textarea" :rows="3" />
        </el-form-item>
        <el-form-item label="封面图片" prop="image">
          <el-upload
            class="avatar-uploader"
            action="/api/v1/upload/image"
            :show-file-list="false"
            :on-success="handleImageSuccess"
          >
            <img v-if="formData.image" :src="formData.image" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="烹饪时间" prop="cookingTime">
          <el-input-number v-model="formData.cookingTime" :min="0" />
          <span style="margin-left: 10px;">分钟</span>
        </el-form-item>
        <el-form-item label="难度" prop="difficulty">
          <el-radio-group v-model="formData.difficulty">
            <el-radio :label="1">简单</el-radio>
            <el-radio :label="2">中等</el-radio>
            <el-radio :label="3">困难</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="份数" prop="servings">
          <el-input-number v-model="formData.servings" :min="1" />
        </el-form-item>
        <el-form-item label="标签" prop="tagIds">
          <el-select v-model="formData.tagIds" multiple placeholder="请选择标签" style="width: 100%;">
            <el-option v-for="tag in tags" :key="tag.id" :label="tag.name" :value="tag.id" />
          </el-select>
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
import { getRecipeList, createRecipe, updateRecipe, deleteRecipe, getAllTags } from '@/api/recipe'

const list = ref([])
const total = ref(0)
const tags = ref([])
const listQuery = reactive({
  keyword: '',
  status: null,
  page: 1,
  pageSize: 10
})

const dialogVisible = ref(false)
const dialogType = ref('create')
const formRef = ref(null)
const formData = reactive({
  id: null,
  name: '',
  subtitle: '',
  description: '',
  image: '',
  cookingTime: 30,
  difficulty: 1,
  servings: 2,
  tagIds: [],
  sortOrder: 0,
  status: 1
})

const rules = {
  name: [{ required: true, message: '请输入食谱名称', trigger: 'blur' }]
}

const getList = async () => {
  try {
    const res = await getRecipeList(listQuery)
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    ElMessage.error('获取列表失败')
  }
}

const loadTags = async () => {
  try {
    const res = await getAllTags()
    tags.value = res.data || []
  } catch (error) {
    console.error('获取标签失败')
  }
}

const handleFilter = () => {
  listQuery.page = 1
  getList()
}

const handleCreate = () => {
  dialogType.value = 'create'
  Object.assign(formData, {
    id: null,
    name: '',
    subtitle: '',
    description: '',
    image: '',
    cookingTime: 30,
    difficulty: 1,
    servings: 2,
    tagIds: [],
    sortOrder: 0,
    status: 1
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogType.value = 'edit'
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleImageSuccess = (response) => {
  if (response.code === 0) {
    formData.image = response.data.url
  }
}

const handleSubmit = async () => {
  await formRef.value.validate()
  try {
    if (dialogType.value === 'create') {
      await createRecipe(formData)
      ElMessage.success('创建成功')
    } else {
      await updateRecipe(formData)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    getList()
  } catch (error) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定删除该食谱吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteRecipe(row.id)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

onMounted(() => {
  getList()
  loadTags()
})
</script>

<style scoped>
.filter-container {
  display: flex;
  align-items: center;
}

.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
