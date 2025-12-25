<template>
  <div class="product-list-container">
    <!-- 搜索栏 -->
    <el-card class="search-card" shadow="never">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="产品分类">
          <el-select v-model="searchForm.categoryId" placeholder="请选择分类" clearable style="width: 200px">
            <el-option
              v-for="item in categoryList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入产品名称" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable style="width: 120px">
            <el-option label="上架" :value="1" />
            <el-option label="下架" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleCreate">新增产品</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 产品列表 -->
    <el-card class="table-card" shadow="never">
      <el-table
        v-loading="loading"
        :data="productList"
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column label="产品图片" width="100">
          <template #default="{ row }">
            <el-image
              :src="row.image"
              fit="cover"
              style="width: 60px; height: 60px; border-radius: 4px"
              :preview-src-list="[row.image]"
            />
          </template>
        </el-table-column>
        <el-table-column prop="name" label="产品名称" min-width="200" />
        <el-table-column prop="nameEn" label="英文名称" min-width="180" />
        <el-table-column label="分类" width="120">
          <template #default="{ row }">
            {{ getCategoryName(row.categoryId) }}
          </template>
        </el-table-column>
        <el-table-column label="价格" width="100">
          <template #default="{ row }">
            ¥{{ row.price }}
          </template>
        </el-table-column>
        <el-table-column prop="stock" label="库存" width="80" />
        <el-table-column prop="weight" label="规格" width="100" />
        <el-table-column prop="viewCount" label="浏览量" width="100" />
        <el-table-column label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '上架' : '下架' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="排序" width="80" prop="sortOrder" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="searchForm.page"
          v-model:page-size="searchForm.pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSearch"
          @current-change="handleSearch"
        />
      </div>
    </el-card>

    <!-- 编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="800px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <el-form-item label="产品分类" prop="categoryId">
          <el-select v-model="formData.categoryId" placeholder="请选择分类" style="width: 100%">
            <el-option
              v-for="item in categoryList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="产品名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入产品名称" />
        </el-form-item>
        <el-form-item label="英文名称" prop="nameEn">
          <el-input v-model="formData.nameEn" placeholder="请输入英文名称" />
        </el-form-item>
        <el-form-item label="副标题" prop="subtitle">
          <el-input v-model="formData.subtitle" placeholder="请输入副标题" />
        </el-form-item>
        <el-form-item label="产品描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="4"
            placeholder="请输入产品描述"
          />
        </el-form-item>
        <el-form-item label="主图片" prop="image">
          <el-input v-model="formData.image" placeholder="请输入图片URL" />
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input-number v-model="formData.price" :min="0" :precision="2" :step="0.1" />
        </el-form-item>
        <el-form-item label="库存" prop="stock">
          <el-input-number v-model="formData.stock" :min="0" />
        </el-form-item>
        <el-form-item label="规格/重量" prop="weight">
          <el-input v-model="formData.weight" placeholder="例如：500g" />
        </el-form-item>
        <el-form-item label="成分说明" prop="ingredients">
          <el-input
            v-model="formData.ingredients"
            type="textarea"
            :rows="3"
            placeholder="请输入成分说明"
          />
        </el-form-item>
        <el-form-item label="使用方法" prop="usage">
          <el-input
            v-model="formData.usage"
            type="textarea"
            :rows="3"
            placeholder="请输入使用方法"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="formData.sortOrder" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">上架</el-radio>
            <el-radio :label="0">下架</el-radio>
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
import {
  getCategoryList,
  getProductList,
  createProduct,
  updateProduct,
  deleteProduct
} from '@/api/product'

// 搜索表单
const searchForm = reactive({
  categoryId: null,
  keyword: '',
  status: null,
  page: 1,
  pageSize: 10
})

// 数据
const loading = ref(false)
const categoryList = ref([])
const productList = ref([])
const total = ref(0)

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新增产品')
const formRef = ref(null)
const formData = reactive({
  id: null,
  categoryId: null,
  name: '',
  nameEn: '',
  subtitle: '',
  description: '',
  image: '',
  price: 0,
  stock: 0,
  weight: '',
  ingredients: '',
  usage: '',
  sortOrder: 0,
  status: 1
})

// 表单验证规则
const formRules = {
  categoryId: [{ required: true, message: '请选择产品分类', trigger: 'change' }],
  name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }]
}

// 获取分类列表
const loadCategoryList = async () => {
  try {
    const res = await getCategoryList({ status: 1 })
    categoryList.value = res.data.list || []
  } catch (error) {
    console.error('获取分类列表失败:', error)
  }
}

// 获取产品列表
const loadProductList = async () => {
  loading.value = true
  try {
    const res = await getProductList(searchForm)
    productList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    ElMessage.error('获取产品列表失败')
  } finally {
    loading.value = false
  }
}

// 获取分类名称
const getCategoryName = (categoryId) => {
  const category = categoryList.value.find(item => item.id === categoryId)
  return category ? category.name : '-'
}

// 搜索
const handleSearch = () => {
  searchForm.page = 1
  loadProductList()
}

// 重置
const handleReset = () => {
  searchForm.categoryId = null
  searchForm.keyword = ''
  searchForm.status = null
  searchForm.page = 1
  loadProductList()
}

// 新增
const handleCreate = () => {
  dialogTitle.value = '新增产品'
  Object.assign(formData, {
    id: null,
    categoryId: null,
    name: '',
    nameEn: '',
    subtitle: '',
    description: '',
    image: '',
    price: 0,
    stock: 0,
    weight: '',
    ingredients: '',
    usage: '',
    sortOrder: 0,
    status: 1
  })
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row) => {
  dialogTitle.value = '编辑产品'
  Object.assign(formData, row)
  dialogVisible.value = true
}

// 提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    try {
      if (formData.id) {
        await updateProduct(formData)
        ElMessage.success('更新成功')
      } else {
        await createProduct(formData)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadProductList()
    } catch (error) {
      ElMessage.error(error.message || '操作失败')
    }
  })
}

// 删除
const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除产品"${row.name}"吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteProduct(row.id)
      ElMessage.success('删除成功')
      loadProductList()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

// 初始化
onMounted(() => {
  loadCategoryList()
  loadProductList()
})
</script>

<style lang="scss" scoped>
.product-list-container {
  padding: 20px;
}

.search-card {
  margin-bottom: 20px;
}

.search-form {
  .el-form-item {
    margin-bottom: 0;
  }
}

.table-card {
  .pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>
