<template>
  <div class="i18n-editor">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 中文 -->
      <el-tab-pane label="中文 (zh-CN)" name="zh-CN">
        <el-form :model="formData['zh-CN']" label-width="100px">
          <el-form-item label="分类名称" required>
            <el-input v-model="formData['zh-CN'].name" placeholder="请输入分类名称" />
          </el-form-item>
          
          <el-form-item label="分类描述">
            <el-input 
              v-model="formData['zh-CN'].description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入分类描述"
            />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 英文 -->
      <el-tab-pane label="English (en-US)" name="en-US">
        <el-form :model="formData['en-US']" label-width="100px">
          <el-form-item label="Category Name" required>
            <el-input v-model="formData['en-US'].name" placeholder="Enter category name" />
          </el-form-item>
          
          <el-form-item label="Description">
            <el-input 
              v-model="formData['en-US'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Enter category description"
            />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 德文 -->
      <el-tab-pane label="Deutsch (de-DE)" name="de-DE">
        <el-form :model="formData['de-DE']" label-width="100px">
          <el-form-item label="Kategoriename" required>
            <el-input v-model="formData['de-DE'].name" placeholder="Kategorienamen eingeben" />
          </el-form-item>
          
          <el-form-item label="Beschreibung">
            <el-input 
              v-model="formData['de-DE'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Kategoriebeschreibung eingeben"
            />
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
    description: ''
  },
  'en-US': {
    name: '',
    description: ''
  },
  'de-DE': {
    name: '',
    description: ''
  }
})

// 监听props变化
watch(() => props.modelValue, (newVal) => {
  if (newVal) {
    Object.keys(newVal).forEach(locale => {
      if (newVal[locale]) {
        formData.value[locale] = { ...formData.value[locale], ...newVal[locale] }
      }
    })
  }
}, { immediate: true, deep: true })

// 监听表单数据变化，同步到父组件
watch(formData, (newVal) => {
  emit('update:modelValue', newVal)
}, { deep: true })
</script>

<style scoped>
.i18n-editor {
  margin-top: 20px;
}
</style>
