<template>
  <div class="i18n-editor">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 中文 -->
      <el-tab-pane label="中文 (zh-CN)" name="zh-CN">
        <el-form :model="formData['zh-CN']" label-width="100px">
          <el-form-item label="标题" required>
            <el-input v-model="formData['zh-CN'].title" placeholder="请输入标题" />
          </el-form-item>
          
          <el-form-item label="描述">
            <el-input 
              v-model="formData['zh-CN'].description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入描述"
            />
          </el-form-item>
          
          <el-form-item label="按钮文字">
            <el-input v-model="formData['zh-CN'].buttonText" placeholder="请输入按钮文字" />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 英文 -->
      <el-tab-pane label="English (en-US)" name="en-US">
        <el-form :model="formData['en-US']" label-width="100px">
          <el-form-item label="Title" required>
            <el-input v-model="formData['en-US'].title" placeholder="Enter title" />
          </el-form-item>
          
          <el-form-item label="Description">
            <el-input 
              v-model="formData['en-US'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Enter description"
            />
          </el-form-item>
          
          <el-form-item label="Button Text">
            <el-input v-model="formData['en-US'].buttonText" placeholder="Enter button text" />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 德文 -->
      <el-tab-pane label="Deutsch (de-DE)" name="de-DE">
        <el-form :model="formData['de-DE']" label-width="100px">
          <el-form-item label="Titel" required>
            <el-input v-model="formData['de-DE'].title" placeholder="Titel eingeben" />
          </el-form-item>
          
          <el-form-item label="Beschreibung">
            <el-input 
              v-model="formData['de-DE'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Beschreibung eingeben"
            />
          </el-form-item>
          
          <el-form-item label="Button-Text">
            <el-input v-model="formData['de-DE'].buttonText" placeholder="Button-Text eingeben" />
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
    title: '',
    description: '',
    buttonText: ''
  },
  'en-US': {
    title: '',
    description: '',
    buttonText: ''
  },
  'de-DE': {
    title: '',
    description: '',
    buttonText: ''
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
