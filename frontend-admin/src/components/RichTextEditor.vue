<template>
  <div class="rich-text-editor">
    <Toolbar
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      :mode="mode"
      class="editor-toolbar"
    />
    <Editor
      v-model="valueHtml"
      :defaultConfig="editorConfig"
      :mode="mode"
      :style="{ height: `${height}px` }"
      class="editor-content"
      @onCreated="handleCreated"
      @onChange="handleChange"
    />
  </div>
</template>

<script setup>
import { ref, shallowRef, onBeforeUnmount, watch } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import '@wangeditor/editor/dist/css/style.css'
import { ElMessage } from 'element-plus'
import { uploadImage } from '@/api/upload'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  height: {
    type: Number,
    default: 400
  },
  placeholder: {
    type: String,
    default: '请输入内容...'
  }
})

const emit = defineEmits(['update:modelValue'])

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()
const valueHtml = ref(props.modelValue)
const mode = 'default' // 或 'simple'

// 工具栏配置
const toolbarConfig = {
  toolbarKeys: [
    'headerSelect',
    'bold',
    'italic',
    'underline',
    'through',
    '|',
    'color',
    'bgColor',
    'fontSize',
    '|',
    'bulletedList',
    'numberedList',
    'todo',
    '|',
    'justifyLeft',
    'justifyCenter',
    'justifyRight',
    '|',
    'insertLink',
    'uploadImage',
    'insertTable',
    'codeBlock',
    'divider',
    '|',
    'undo',
    'redo',
    '|',
    'fullScreen'
  ]
}

// 编辑器配置
const editorConfig = {
  placeholder: props.placeholder,
  MENU_CONF: {
    // 配置上传图片
    uploadImage: {
      async customUpload(file, insertFn) {
        try {
          const res = await uploadImage(file)
          const url = `/api${res.data.url}`
          insertFn(url, file.name, url)
          ElMessage.success('图片上传成功')
        } catch (error) {
          ElMessage.error('图片上传失败')
        }
      }
    }
  }
}

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

// 编辑器创建完成
const handleCreated = (editor) => {
  editorRef.value = editor
}

// 内容变化
const handleChange = (editor) => {
  emit('update:modelValue', valueHtml.value)
}

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  if (valueHtml.value !== newVal) {
    valueHtml.value = newVal
  }
})
</script>

<style lang="scss" scoped>
.rich-text-editor {
  border: 1px solid #ccc;
  border-radius: 4px;
  overflow: hidden;
  
  .editor-toolbar {
    border-bottom: 1px solid #ccc;
  }
  
  .editor-content {
    overflow-y: auto;
    
    :deep(.w-e-text-container) {
      background-color: #fff;
    }
    
    :deep(.w-e-text-placeholder) {
      top: 10px;
      left: 10px;
    }
  }
}
</style>
