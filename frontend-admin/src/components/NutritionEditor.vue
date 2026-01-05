<template>
  <div class="nutrition-editor">
    <el-table :data="nutritionList" border style="width: 100%">
      <el-table-column label="营养成分" width="200">
        <template #default="{ row }">
          <el-input v-model="row.name" placeholder="例如：能量" />
        </template>
      </el-table-column>
      <el-table-column label="含量" width="200">
        <template #default="{ row }">
          <el-input v-model="row.value" placeholder="例如：0kJ" />
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template #default="{ $index }">
          <el-button type="danger" size="small" @click="removeRow($index)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button type="primary" size="small" style="margin-top: 10px" @click="addRow">添加行</el-button>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const nutritionList = ref([])

const parseValue = (val) => {
  if (!val) return []
  try {
    return JSON.parse(val)
  } catch {
    return []
  }
}

const initData = () => {
  const data = parseValue(props.modelValue)
  nutritionList.value = data.length > 0 ? data : [{ name: '', value: '' }]
}

const addRow = () => {
  nutritionList.value.push({ name: '', value: '' })
  emitChange()
}

const removeRow = (index) => {
  nutritionList.value.splice(index, 1)
  if (nutritionList.value.length === 0) {
    nutritionList.value.push({ name: '', value: '' })
  }
  emitChange()
}

const emitChange = () => {
  const filtered = nutritionList.value.filter(item => item.name || item.value)
  emit('update:modelValue', JSON.stringify(filtered))
}

watch(() => props.modelValue, (val) => {
  if (val !== JSON.stringify(nutritionList.value)) {
    initData()
  }
}, { immediate: true })

watch(nutritionList, () => {
  emitChange()
}, { deep: true })
</script>

<style scoped>
.nutrition-editor {
  width: 100%;
}
</style>
