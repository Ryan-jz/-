<template>
  <div class="purchase-link-editor">
    <div v-for="(links, platform) in purchaseLinks" :key="platform" class="platform-group">
      <div class="platform-header">
        <el-input v-model="platformNames[platform]" placeholder="平台名称" style="width: 200px" @blur="updatePlatformName(platform)" />
        <el-button type="danger" size="small" @click="removePlatform(platform)">删除平台</el-button>
      </div>
      <div v-for="(link, index) in links" :key="index" class="link-item">
        <el-input v-model="link.spec" placeholder="规格" style="width: 150px" />
        <el-input v-model="link.url" placeholder="链接" style="flex: 1" />
        <el-button type="danger" size="small" @click="removeLink(platform, index)">删除</el-button>
      </div>
      <el-button type="primary" size="small" @click="addLink(platform)">添加链接</el-button>
    </div>
    <el-button type="success" @click="addPlatform">添加平台</el-button>
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

const purchaseLinks = ref({})
const platformNames = ref({})

const initData = () => {
  try {
    if (props.modelValue) {
      const data = JSON.parse(props.modelValue)
      purchaseLinks.value = data
      platformNames.value = Object.keys(data).reduce((acc, key) => {
        acc[key] = key
        return acc
      }, {})
    }
  } catch (e) {
    purchaseLinks.value = {}
    platformNames.value = {}
  }
}

const addPlatform = () => {
  const key = `platform_${Date.now()}`
  purchaseLinks.value[key] = []
  platformNames.value[key] = '新平台'
}

const removePlatform = (platform) => {
  delete purchaseLinks.value[platform]
  delete platformNames.value[platform]
}

const updatePlatformName = (oldKey) => {
  const newKey = platformNames.value[oldKey]
  if (newKey && newKey !== oldKey) {
    const newData = {}
    Object.keys(purchaseLinks.value).forEach(key => {
      if (key === oldKey) {
        newData[newKey] = purchaseLinks.value[key]
      } else {
        newData[key] = purchaseLinks.value[key]
      }
    })
    purchaseLinks.value = newData
    
    const newNames = {}
    Object.keys(platformNames.value).forEach(key => {
      if (key === oldKey) {
        newNames[newKey] = newKey
      } else {
        newNames[key] = platformNames.value[key]
      }
    })
    platformNames.value = newNames
  }
  emitChange()
}

const addLink = (platform) => {
  if (!purchaseLinks.value[platform]) {
    purchaseLinks.value[platform] = []
  }
  purchaseLinks.value[platform].push({ spec: '', url: '' })
}

const removeLink = (platform, index) => {
  purchaseLinks.value[platform].splice(index, 1)
}

const emitChange = () => {
  const result = {}
  Object.keys(purchaseLinks.value).forEach(key => {
    const platformName = platformNames.value[key] || key
    result[platformName] = purchaseLinks.value[key].filter(link => link.spec || link.url)
  })
  emit('update:modelValue', JSON.stringify(result))
}

watch(() => props.modelValue, initData, { immediate: true })
watch(purchaseLinks, () => {
  emitChange()
}, { deep: true })
</script>

<style lang="scss" scoped>
.purchase-link-editor {
  .platform-group {
    margin-bottom: 20px;
    padding: 15px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    
    .platform-header {
      display: flex;
      gap: 10px;
      margin-bottom: 10px;
      align-items: center;
    }
    
    .link-item {
      display: flex;
      gap: 10px;
      margin-bottom: 10px;
      align-items: center;
    }
  }
}
</style>
