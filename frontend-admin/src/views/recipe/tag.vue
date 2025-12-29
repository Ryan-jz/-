<template>
  <div class="app-container">
    <el-card>
      <div class="filter-container">
        <el-input v-model="listQuery.keyword" placeholder="标签名称" style="width: 200px;" clearable @keyup.enter="handleFilter" />
        <el-select v-model="listQuery.status" placeholder="状态" clearable style="width: 120px; margin-left: 10px;">
          <el-option label="启用" :value="1" />
          <el-option label="禁用" :value="0" />
        </el-select>
        <el-button type="primary" icon="Search" style="margin-left: 10px;" @click="handleFilter">搜索</el-button>
        <el-button type="success" icon="Plus" @click="handleCreate">新增</el-button>
      </div>

      <el-table :data="list" border style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="标签名称" />
        <el-table-column prop="slug" label="标签标识" />
        <el-table-column prop="sortOrder" label="排序" width="100" />
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

    <el-dialog v-model="dialogVisible" :title="dialogType === 'create' ? '新增标签' : '编辑标签'" width="500px">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入标签名称" />
        </el-form-item>
        <el-form-item label="标签标识" prop="slug">
          <el-input v-model="formData.slug" placeholder="请输入标签标识(英文)" />
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
import { getTagList, createTag, updateTag, deleteTag } from '@/api/recipe'

const list = ref([])
const total = ref(0)
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
  slug: '',
  sortOrder: 0,
  status: 1
})

const rules = {
  name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }],
  slug: [{ required: true, message: '请输入标签标识', trigger: 'blur' }]
}

const getList = async () => {
  try {
    const res = await getTagList(listQuery)
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    ElMessage.error('获取列表失败')
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
    slug: '',
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

const handleSubmit = async () => {
  await formRef.value.validate()
  try {
    if (dialogType.value === 'create') {
      await createTag(formData)
      ElMessage.success('创建成功')
    } else {
      await updateTag(formData)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    getList()
  } catch (error) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定删除该标签吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteTag(row.id)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

onMounted(() => {
  getList()
})
</script>

<style scoped>
.filter-container {
  display: flex;
  align-items: center;
}
</style>
