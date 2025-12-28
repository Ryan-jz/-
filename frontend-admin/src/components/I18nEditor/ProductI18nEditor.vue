<template>
  <div class="i18n-editor">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 中文 -->
      <el-tab-pane label="中文 (zh-CN)" name="zh-CN">
        <el-form :model="formData['zh-CN']" label-width="120px">
          <el-form-item label="产品名称" required>
            <el-input v-model="formData['zh-CN'].name" placeholder="请输入产品名称" />
          </el-form-item>
          
          <el-form-item label="副标题">
            <el-input v-model="formData['zh-CN'].subtitle" placeholder="请输入副标题" />
          </el-form-item>
          
          <el-form-item label="产品描述">
            <el-input 
              v-model="formData['zh-CN'].description" 
              type="textarea" 
              :rows="4"
              placeholder="请输入产品描述"
            />
          </el-form-item>
          
          <el-form-item label="成分说明">
            <el-input 
              v-model="formData['zh-CN'].ingredients" 
              type="textarea" 
              :rows="3"
              placeholder="请输入成分说明"
            />
          </el-form-item>
          
          <el-form-item label="使用方法">
            <el-input 
              v-model="formData['zh-CN'].usage" 
              type="textarea" 
              :rows="3"
              placeholder="请输入使用方法"
            />
          </el-form-item>
          
          <el-form-item label="产品特点">
            <el-input 
              v-model="featuresText['zh-CN']" 
              type="textarea" 
              :rows="3"
              placeholder="每行一个特点"
            />
            <div class="tip">提示：每行输入一个特点</div>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 英文 -->
      <el-tab-pane label="English (en-US)" name="en-US">
        <el-form :model="formData['en-US']" label-width="120px">
          <el-form-item label="Product Name" required>
            <el-input v-model="formData['en-US'].name" placeholder="Enter product name" />
          </el-form-item>
          
          <el-form-item label="Subtitle">
            <el-input v-model="formData['en-US'].subtitle" placeholder="Enter subtitle" />
          </el-form-item>
          
          <el-form-item label="Description">
            <el-input 
              v-model="formData['en-US'].description" 
              type="textarea" 
              :rows="4"
              placeholder="Enter product description"
            />
          </el-form-item>
          
          <el-form-item label="Ingredients">
            <el-input 
              v-model="formData['en-US'].ingredients" 
              type="textarea" 
              :rows="3"
              placeholder="Enter ingredients"
            />
          </el-form-item>
          
          <el-form-item label="Usage">
            <el-input 
              v-model="formData['en-US'].usage" 
              type="textarea" 
              :rows="3"
              placeholder="Enter usage instructions"
            />
          </el-form-item>
          
          <el-form-item label="Features">
            <el-input 
              v-model="featuresText['en-US']" 
              type="textarea" 
              :rows="3"
              placeholder="One feature per line"
            />
            <div class="tip">Tip: Enter one feature per line</div>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 德文 -->
      <el-tab-pane label="Deutsch (de-DE)" name="de-DE">
        <el-form :model="formData['de-DE']" label-width="120px">
          <el-form-item label="Produktname" required>
            <el-input v-model="formData['de-DE'].name" placeholder="Produktname eingeben" />
          </el-form-item>
          
          <el-form-item label="Untertitel">
            <el-input v-model="formData['de-DE'].subtitle" placeholder="Untertitel eingeben" />
          </el-form-item>
          
          <el-form-item label="Beschreibung">
            <el-input 
              v-model="formData['de-DE'].description" 
              type="textarea" 
              :rows="4"
              placeholder="Produktbeschreibung eingeben"
            />
          </el-form-item>
          
          <el-form-item label="Zutaten">
            <el-input 
              v-model="formData['de-DE'].ingredients" 
              type="textarea" 
              :rows="3"
              placeholder="Zutaten eingeben"
            />
          </el-form-item>
          
          <el-form-item label="Verwendung">
            <el-input 
              v-model="formData['de-DE'].usage" 
              type="textarea" 
              :rows="3"
              placeholder="Verwendungshinweise eingeben"
            />
          </el-form-item>
          
          <el-form-item label="Eigenschaften">
            <el-input 
              v-model="featuresText['de-DE']" 
              type="textarea" 
              :rows="3"
              placeholder="Eine Eigenschaft pro Zeile"
            />
            <div class="tip">Tipp: Geben Sie eine Eigenschaft pro Zeile ein</div>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      'zh-CN': {},
      'en-US': {},
      'de-DE': {}
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const activeTab = ref('zh-CN')

// 表单数据
const formData = ref({
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
})

// 特点文本（用于编辑）
const featuresText = ref({
  'zh-CN': '',
  'en-US': '',
  'de-DE': ''
})

// 监听props变化
watch(() => props.modelValue, (newVal) => {
  if (newVal) {
    Object.keys(newVal).forEach(locale => {
      if (newVal[locale]) {
        formData.value[locale] = { ...formData.value[locale], ...newVal[locale] }
        
        // 转换features数组为文本
        if (newVal[locale].features && Array.isArray(newVal[locale].features)) {
          featuresText.value[locale] = newVal[locale].features.join('\n')
        }
      }
    })
  }
}, { immediate: true, deep: true })

// 监听表单数据变化，同步到父组件
watch(formData, (newVal) => {
  const result = {}
  Object.keys(newVal).forEach(locale => {
    result[locale] = {
      ...newVal[locale],
      features: featuresText.value[locale] 
        ? featuresText.value[locale].split('\n').filter(f => f.trim())
        : []
    }
  })
  emit('update:modelValue', result)
}, { deep: true })

// 监听特点文本变化
watch(featuresText, () => {
  const result = {}
  Object.keys(formData.value).forEach(locale => {
    result[locale] = {
      ...formData.value[locale],
      features: featuresText.value[locale] 
        ? featuresText.value[locale].split('\n').filter(f => f.trim())
        : []
    }
  })
  emit('update:modelValue', result)
}, { deep: true })
</script>

<style scoped>
.i18n-editor {
  margin-top: 20px;
}

.tip {
  font-size: 12px;
  color: #999;
  margin-top: 5px;
}
</style>
