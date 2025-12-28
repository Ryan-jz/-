<template>
  <div class="product-edit">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>{{ isEdit ? '编辑产品' : '新增产品' }}</span>
          <el-button @click="handleBack">返回</el-button>
        </div>
      </template>

      <el-form :model="form" :rules="rules" ref="formRef" label-width="120px">
        <!-- 基础信息 -->
        <el-divider content-position="left">基础信息</el-divider>
        
        <el-form-item label="产品分类" prop="category_id" required>
          <el-select v-model="form.category_id" placeholder="请选择分类">
            <el-option
              v-for="cat in categories"
              :key="cat.id"
              :label="cat.name"
              :value="cat.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="主图片" prop="image">
          <el-input v-model="form.image" placeholder="请输入图片URL" />
        </el-form-item>

        <el-form-item label="价格" prop="price" required>
          <el-input-number v-model="form.price" :min="0" :precision="2" />
          <span style="margin-left: 10px;">€</span>
        </el-form-item>

        <el-form-item label="库存" prop="stock">
          <el-input-number v-model="form.stock" :min="0" />
        </el-form-item>

        <el-form-item label="规格/重量" prop="weight">
          <el-input v-model="form.weight" placeholder="例如：500g" />
        </el-form-item>

        <el-form-item label="营养信息" prop="nutrition">
          <el-input
            v-model="form.nutrition"
            type="textarea"
            :rows="3"
            placeholder="JSON格式"
          />
        </el-form-item>

        <el-form-item label="排序" prop="sort_order">
          <el-input-number v-model="form.sort_order" :min="0" />
        </el-form-item>

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">上架</el-radio>
            <el-radio :label="0">下架</el-radio>
          </el-radio-group>
        </el-form-item>

        <!-- 国际化内容 -->
        <el-divider content-position="left">多语言内容</el-divider>
        
        <ProductI18nEditor v-model="form.i18n" />

        <!-- 操作按钮 -->
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" :loading="submitting">
            保存
          </el-button>
          <el-button @click="handleBack">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import ProductI18nEditor from '@/components/I18nEditor/ProductI18nEditor.vue'
import request from '@/utils/request'

const route = useRoute()
const router = useRouter()
const formRef = ref(null)

const isEdit = computed(() => !!route.params.id)

// 分类列表
const categories = ref([])

// 表单数据
const form = reactive({
  category_id: null,
  image: '',
  price: 0,
  stock: 0,
  weight: '',
  nutrition: '',
  sort_order: 0,
  status: 1,
  // 国际化数据
  i18n: {
    'zh-CN': {
      name: '',
      subtitle: '',
      description: '',
      ingredients: '',
      usage: '',
      features: []
    },
    'en-US': {
      name: '',
      subtitle: '',
      description: '',
      ingredients: '',
      usage: '',
      features: []
    },
    'de-DE': {
      name: '',
      subtitle: '',
      description: '',
      ingredients: '',
      usage: '',
      features: []
    }
  }
})

// 表单验证规则
const rules = {
  category_id: [
    { required: true, message: '请选择产品分类', trigger: 'change' }
  ],
  price: [
    { required: true, message: '请输入价格', trigger: 'blur' }
  ]
}

const submitting = ref(false)

// 获取分类列表
const fetchCategories = async () => {
  try {
    const res = await request.get('/product/categories')
    categories.value = res.data || []
  } catch (error) {
    console.error('获取分类失败:', error)
  }
}

// 获取产品详情（编辑时）
const fetchProductDetail = async () => {
  try {
    const res = await request.get(`/admin/products/${route.params.id}`)
    const product = res.data
    
    // 填充基础信息
    Object.assign(form, {
      category_id: product.category_id,
      image: product.image,
      price: product.price,
      stock: product.stock,
      weight: product.weight,
      nutrition: product.nutrition,
      sort_order: product.sort_order,
      status: product.status
    })
    
    // 获取国际化数据
    const i18nRes = await request.get(`/admin/products/${route.params.id}/i18n`)
    if (i18nRes.data) {
      form.i18n = i18nRes.data
    }
  } catch (error) {
    console.error('获取产品详情失败:', error)
    ElMessage.error('获取产品详情失败')
  }
}

// 提交表单
const handleSubmit = async () => {
  try {
    // 验证表单
    await formRef.value.validate()
    
    submitting.value = true
    
    // 准备提交数据
    const submitData = {
      ...form,
      // 将i18n数据转换为后端需要的格式
      i18n: form.i18n
    }
    
    if (isEdit.value) {
      // 更新
      await request.put(`/admin/products/${route.params.id}`, submitData)
      ElMessage.success('更新成功')
    } else {
      // 创建
      await request.post('/admin/products', submitData)
      ElMessage.success('创建成功')
    }
    
    // 返回列表页
    router.push('/products')
  } catch (error) {
    console.error('保存失败:', error)
    ElMessage.error('保存失败：' + (error.message || '未知错误'))
  } finally {
    submitting.value = false
  }
}

// 返回
const handleBack = () => {
  router.back()
}

onMounted(() => {
  fetchCategories()
  if (isEdit.value) {
    fetchProductDetail()
  }
})
</script>

<style scoped>
.product-edit {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

:deep(.el-divider__text) {
  font-size: 16px;
  font-weight: bold;
  color: #409eff;
}
</style>
