<template>
  <div class="recipe-list-container">
    <!-- 搜索栏 -->
    <el-card class="search-card" shadow="never">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入食谱名称" clearable style="width: 200px" />
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
          <el-button type="success" @click="handleCreate">新增食谱</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 食谱列表 -->
    <el-card class="table-card" shadow="never">
      <el-table
        v-loading="loading"
        :data="recipeList"
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column label="食谱图片" width="100">
          <template #default="{ row }">
            <el-image
              v-if="row.image"
              :src="getImageUrl(row.image)"
              fit="cover"
              style="width: 60px; height: 60px; border-radius: 4px"
              :preview-src-list="[getImageUrl(row.image)]"
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="食谱名称" min-width="200" />
        <el-table-column prop="nameEn" label="英文名称" min-width="180" />
        <el-table-column label="难度" width="100">
          <template #default="{ row }">
            <el-tag :type="getDifficultyType(row.difficulty)">
              {{ getDifficultyText(row.difficulty) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="cookingTime" label="烹饪时间" width="100">
          <template #default="{ row }">
            {{ row.cookingTime }}分钟
          </template>
        </el-table-column>
        <el-table-column prop="servings" label="份量" width="80">
          <template #default="{ row }">
            {{ row.servings }}人份
          </template>
        </el-table-column>
        <el-table-column prop="viewCount" label="浏览量" width="100" />
        <el-table-column prop="likeCount" label="点赞数" width="100" />
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
      width="900px"
      :close-on-click-modal="false"
      destroy-on-close
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <el-divider content-position="left">基本信息</el-divider>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="烹饪时间" prop="cookingTime">
              <el-input-number v-model="formData.cookingTime" :min="0" :step="5" />
              <span style="margin-left: 10px">分钟</span>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="难度等级" prop="difficulty">
              <el-select v-model="formData.difficulty" style="width: 100%">
                <el-option label="简单" :value="1" />
                <el-option label="中等" :value="2" />
                <el-option label="困难" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="份量" prop="servings">
              <el-input-number v-model="formData.servings" :min="1" />
              <span style="margin-left: 10px">人份</span>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="食材列表" prop="ingredients">
          <div class="ingredients-list">
            <div v-for="(item, index) in formData.ingredients" :key="index" class="ingredient-item">
              <el-input v-model="item.name" placeholder="食材名称" style="width: 200px" />
              <el-input v-model="item.amount" placeholder="用量" style="width: 150px; margin-left: 10px" />
              <el-button type="danger" size="small" @click="removeIngredient(index)" style="margin-left: 10px">删除</el-button>
            </div>
            <el-button type="primary" size="small" @click="addIngredient">添加食材</el-button>
          </div>
        </el-form-item>

        <el-form-item label="关联产品" prop="productIds">
          <el-select v-model="formData.productIds" multiple placeholder="请选择关联产品" style="width: 100%">
            <el-option
              v-for="item in productList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="标签" prop="tags">
          <el-input v-model="formData.tags" placeholder="多个标签用逗号分隔，例如：烤鸡,主菜,家常菜" />
        </el-form-item>

        <el-form-item label="制作步骤" prop="content">
          <RichTextEditor v-model="formData.content" />
        </el-form-item>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="排序" prop="sortOrder">
              <el-input-number v-model="formData.sortOrder" :min="0" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="formData.status">
                <el-radio :label="1">上架</el-radio>
                <el-radio :label="0">下架</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
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
import ImageUpload from '@/components/ImageUpload.vue'
import RichTextEditor from '@/components/RichTextEditor.vue'
import RecipeI18nEditor from '@/components/I18nEditor/RecipeI18nEditor.vue'
import {
  getRecipeList,
  createRecipe,
  updateRecipe,
  deleteRecipe
} from '@/api/recipe'
import { getProductList } from '@/api/product'

// 搜索表单
const searchForm = reactive({
  keyword: '',
  status: null,
  page: 1,
  pageSize: 10
})

// 数据
const loading = ref(false)
const recipeList = ref([])
const productList = ref([])
const total = ref(0)

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新增食谱')
const formRef = ref(null)
const formData = reactive({
  id: null,
  image: '',
  images: [],
  ingredients: [{ name: '', amount: '' }],
  content: '',
  cookingTime: 30,
  difficulty: 1,
  servings: 2,
  productIds: [],
  sortOrder: 0,
  status: 1
})

// 表单验证规则
const formRules = {
  image: [{ required: true, message: '请上传食谱主图', trigger: 'change' }]
}

// 获取完整图片URL
const getImageUrl = (url) => {
  if (!url) return ''
  if (url.startsWith('http')) return url
  return `/api${url}`
}

// 获取难度文本
const getDifficultyText = (difficulty) => {
  const map = { 1: '简单', 2: '中等', 3: '困难' }
  return map[difficulty] || '-'
}

// 获取难度标签类型
const getDifficultyType = (difficulty) => {
  const map = { 1: 'success', 2: 'warning', 3: 'danger' }
  return map[difficulty] || ''
}

// 添加食材
const addIngredient = () => {
  formData.ingredients.push({ name: '', amount: '' })
}

// 删除食材
const removeIngredient = (index) => {
  formData.ingredients.splice(index, 1)
}

// 获取产品列表
const loadProductList = async () => {
  try {
    const res = await getProductList({ status: 1, page: 1, pageSize: 100 })
    productList.value = res.data.list || []
  } catch (error) {
    console.error('获取产品列表失败:', error)
  }
}

// 获取食谱列表
const loadRecipeList = async () => {
  loading.value = true
  try {
    const res = await getRecipeList(searchForm)
    recipeList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    ElMessage.error('获取食谱列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  searchForm.page = 1
  loadRecipeList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  searchForm.status = null
  searchForm.page = 1
  loadRecipeList()
}

// 新增
const handleCreate = () => {
  dialogTitle.value = '新增食谱'
  Object.assign(formData, {
    id: null,
    image: '',
    images: [],
    ingredients: [{ name: '', amount: '' }],
    content: '',
    cookingTime: 30,
    difficulty: 1,
    servings: 2,
    productIds: [],
    sortOrder: 0,
    status: 1
  })
  dialogVisible.value = true
}

// 编辑
const handleEdit = async (row) => {
  dialogTitle.value = '编辑食谱'
  
  // 复制基本数据
  Object.assign(formData, {
    id: row.id,
    image: row.image,
    productIds: row.productIds ? row.productIds.split(',').map(Number) : [],
    ingredients: typeof row.ingredients === 'string' ? JSON.parse(row.ingredients || '[]') : row.ingredients || [{ name: '', amount: '' }],
    images: typeof row.images === 'string' ? JSON.parse(row.images || '[]') : row.images || [],
    content: row.content,
    cookingTime: row.cookingTime,
    difficulty: row.difficulty,
    servings: row.servings,
    sortOrder: row.sortOrder,
    status: row.status
  })
  
  dialogVisible.value = true
}

// 提交
const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    try {
      const data = {
        ...formData,
        productIds: formData.productIds.join(','),
        ingredients: formData.ingredients.filter(item => item.name && item.amount)
      }
      
      if (formData.id) {
        await updateRecipe(data)
        ElMessage.success('更新成功')
      } else {
        await createRecipe(data)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadRecipeList()
    } catch (error) {
      ElMessage.error(error.message || '操作失败')
    }
  })
}

// 删除
const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除食谱"${row.name}"吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteRecipe(row.id)
      ElMessage.success('删除成功')
      loadRecipeList()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

// 初始化
onMounted(() => {
  loadProductList()
  loadRecipeList()
})
</script>

<style lang="scss" scoped>
.recipe-list-container {
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

.ingredients-list {
  width: 100%;
  
  .ingredient-item {
    margin-bottom: 10px;
    display: flex;
    align-items: center;
  }
}
</style>
