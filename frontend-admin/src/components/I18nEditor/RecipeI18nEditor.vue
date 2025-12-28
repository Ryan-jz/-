<template>
  <div class="i18n-editor">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 中文 -->
      <el-tab-pane label="中文 (zh-CN)" name="zh-CN">
        <el-form :model="formData['zh-CN']" label-width="120px">
          <el-form-item label="食谱名称" required>
            <el-input v-model="formData['zh-CN'].name" placeholder="请输入食谱名称" />
          </el-form-item>
          
          <el-form-item label="副标题">
            <el-input v-model="formData['zh-CN'].subtitle" placeholder="请输入副标题" />
          </el-form-item>
          
          <el-form-item label="简介">
            <el-input 
              v-model="formData['zh-CN'].description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入简介"
            />
          </el-form-item>
          
          <el-form-item label="食材列表">
            <el-input 
              v-model="ingredientsText['zh-CN']" 
              type="textarea" 
              :rows="5"
              placeholder="格式：食材名称|用量（每行一个）&#10;例如：鸡肉|500g"
            />
            <div class="tip">格式：食材名称|用量（每行一个），例如：鸡肉|500g</div>
          </el-form-item>
          
          <el-form-item label="制作步骤">
            <el-input 
              v-model="formData['zh-CN'].content" 
              type="textarea" 
              :rows="8"
              placeholder="请输入制作步骤（支持HTML）"
            />
          </el-form-item>
          
          <el-form-item label="标签">
            <el-input 
              v-model="formData['zh-CN'].tags" 
              placeholder="多个标签用逗号分隔"
            />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 英文 -->
      <el-tab-pane label="English (en-US)" name="en-US">
        <el-form :model="formData['en-US']" label-width="120px">
          <el-form-item label="Recipe Name" required>
            <el-input v-model="formData['en-US'].name" placeholder="Enter recipe name" />
          </el-form-item>
          
          <el-form-item label="Subtitle">
            <el-input v-model="formData['en-US'].subtitle" placeholder="Enter subtitle" />
          </el-form-item>
          
          <el-form-item label="Description">
            <el-input 
              v-model="formData['en-US'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Enter description"
            />
          </el-form-item>
          
          <el-form-item label="Ingredients">
            <el-input 
              v-model="ingredientsText['en-US']" 
              type="textarea" 
              :rows="5"
              placeholder="Format: ingredient|amount (one per line)&#10;Example: Chicken|500g"
            />
            <div class="tip">Format: ingredient|amount (one per line), e.g., Chicken|500g</div>
          </el-form-item>
          
          <el-form-item label="Steps">
            <el-input 
              v-model="formData['en-US'].content" 
              type="textarea" 
              :rows="8"
              placeholder="Enter cooking steps (HTML supported)"
            />
          </el-form-item>
          
          <el-form-item label="Tags">
            <el-input 
              v-model="formData['en-US'].tags" 
              placeholder="Separate multiple tags with commas"
            />
          </el-form-item>
        </el-form>
      </el-tab-pane>
      
      <!-- 德文 -->
      <el-tab-pane label="Deutsch (de-DE)" name="de-DE">
        <el-form :model="formData['de-DE']" label-width="120px">
          <el-form-item label="Rezeptname" required>
            <el-input v-model="formData['de-DE'].name" placeholder="Rezeptname eingeben" />
          </el-form-item>
          
          <el-form-item label="Untertitel">
            <el-input v-model="formData['de-DE'].subtitle" placeholder="Untertitel eingeben" />
          </el-form-item>
          
          <el-form-item label="Beschreibung">
            <el-input 
              v-model="formData['de-DE'].description" 
              type="textarea" 
              :rows="3"
              placeholder="Beschreibung eingeben"
            />
          </el-form-item>
          
          <el-form-item label="Zutaten">
            <el-input 
              v-model="ingredientsText['de-DE']" 
              type="textarea" 
              :rows="5"
              placeholder="Format: Zutat|Menge (eine pro Zeile)&#10;Beispiel: Hähnchen|500g"
            />
            <div class="tip">Format: Zutat|Menge (eine pro Zeile), z.B. Hähnchen|500g</div>
          </el-form-item>
          
          <el-form-item label="Schritte">
            <el-input 
              v-model="formData['de-DE'].content" 
              type="textarea" 
              :rows="8"
              placeholder="Kochschritte eingeben (HTML unterstützt)"
            />
          </el-form-item>
          
          <el-form-item label="Tags">
            <el-input 
              v-model="formData['de-DE'].tags" 
              placeholder="Mehrere Tags durch Kommas trennen"
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
    subtitle: '',
    description: '',
    ingredients: [],
    content: '',
    tags: ''
  },
  'en-US': {
    name: '',
    subtitle: '',
    description: '',
    ingredients: [],
    content: '',
    tags: ''
  },
  'de-DE': {
    name: '',
    subtitle: '',
    description: '',
    ingredients: [],
    content: '',
    tags: ''
  }
})

// 食材文本（用于编辑）
const ingredientsText = ref({
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
        
        // 转换ingredients数组为文本
        if (newVal[locale].ingredients && Array.isArray(newVal[locale].ingredients)) {
          ingredientsText.value[locale] = newVal[locale].ingredients
            .map(item => `${item.name}|${item.amount}`)
            .join('\n')
        }
      }
    })
  }
}, { immediate: true, deep: true })

// 解析食材文本为数组
const parseIngredients = (text) => {
  if (!text) return []
  return text.split('\n')
    .filter(line => line.trim())
    .map(line => {
      const [name, amount] = line.split('|').map(s => s.trim())
      return { name: name || '', amount: amount || '' }
    })
}

// 监听表单数据变化，同步到父组件
watch(formData, (newVal) => {
  const result = {}
  Object.keys(newVal).forEach(locale => {
    result[locale] = {
      ...newVal[locale],
      ingredients: parseIngredients(ingredientsText.value[locale])
    }
  })
  emit('update:modelValue', result)
}, { deep: true })

// 监听食材文本变化
watch(ingredientsText, () => {
  const result = {}
  Object.keys(formData.value).forEach(locale => {
    result[locale] = {
      ...formData.value[locale],
      ingredients: parseIngredients(ingredientsText.value[locale])
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
