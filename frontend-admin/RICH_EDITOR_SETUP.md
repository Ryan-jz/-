# 富文本编辑器安装指南

## 使用的插件

**wangEditor** - 开源 Web 富文本编辑器
- 官网：https://www.wangeditor.com/
- GitHub：https://github.com/wangeditor-team/wangEditor
- 文档：https://www.wangeditor.com/v5/

## 安装步骤

### 1. 安装依赖

在 `frontend-admin` 目录下执行：

```bash
pnpm install @wangeditor/editor@^5.1.23 @wangeditor/editor-for-vue@^5.1.12
```

或者使用 npm：

```bash
npm install @wangeditor/editor@^5.1.23 @wangeditor/editor-for-vue@^5.1.12
```

### 2. 组件已创建

富文本编辑器组件已创建在：
- `frontend-admin/src/components/RichTextEditor.vue`

### 3. 使用方法

在食谱管理页面中已集成：

```vue
<template>
  <el-form-item label="制作步骤" prop="content">
    <RichTextEditor v-model="formData.content" :height="400" />
  </el-form-item>
</template>

<script setup>
import RichTextEditor from '@/components/RichTextEditor.vue'
import { ref } from 'vue'

const formData = ref({
  content: ''
})
</script>
```

## 功能特性

### 工具栏功能

1. **文本格式**
   - 标题选择（H1-H6）
   - 粗体、斜体、下划线、删除线
   - 字体颜色、背景色
   - 字号调整

2. **段落格式**
   - 无序列表
   - 有序列表
   - 待办事项
   - 对齐方式（左对齐、居中、右对齐）

3. **插入内容**
   - 插入链接
   - 上传图片（集成后端上传接口）
   - 插入表格
   - 代码块
   - 分割线

4. **编辑操作**
   - 撤销/重做
   - 全屏编辑

### 图片上传

已集成后端图片上传接口：
- 点击工具栏的"上传图片"按钮
- 选择图片文件
- 自动上传到服务器
- 插入到编辑器中

### 自定义配置

可以在组件中修改配置：

```javascript
// 工具栏配置
const toolbarConfig = {
  toolbarKeys: [
    'headerSelect',
    'bold',
    'italic',
    // ... 更多工具
  ]
}

// 编辑器配置
const editorConfig = {
  placeholder: '请输入内容...',
  MENU_CONF: {
    uploadImage: {
      // 自定义上传逻辑
    }
  }
}
```

## 数据存储

### 保存格式

编辑器内容以 HTML 格式保存到数据库：

```html
<h3>步骤1：准备食材</h3>
<p>将整鸡清洗干净，沥干水分。</p>
<h3>步骤2：腌制</h3>
<p>用喜马拉雅粉盐、黑胡椒均匀涂抹在鸡身内外，腌制2小时。</p>
<img src="/uploads/images/2024/12/25/xxx.jpg" alt="图片" />
```

### 前端渲染

在用户端直接使用 `v-html` 渲染：

```vue
<template>
  <div class="recipe-content" v-html="recipe.content"></div>
</template>

<style>
.recipe-content {
  h3 {
    font-size: 20px;
    font-weight: bold;
    margin: 15px 0 10px;
  }
  
  p {
    line-height: 1.8;
    margin: 10px 0;
  }
  
  img {
    max-width: 100%;
    border-radius: 8px;
    margin: 15px 0;
  }
}
</style>
```

## 常见问题

### 1. 编辑器不显示

确保已安装依赖：
```bash
pnpm install
```

### 2. 图片上传失败

- 检查后端上传接口是否正常
- 确认文件大小不超过限制（5MB）
- 查看浏览器控制台错误信息

### 3. 样式问题

确保引入了编辑器样式：
```javascript
import '@wangeditor/editor/dist/css/style.css'
```

### 4. 内容不更新

使用 `shallowRef` 而不是 `ref`：
```javascript
const editorRef = shallowRef()
```

## 高级用法

### 自定义菜单

```javascript
const toolbarConfig = {
  excludeKeys: ['group-video'], // 排除视频菜单
  insertKeys: {
    index: 5,
    keys: ['customMenu'] // 插入自定义菜单
  }
}
```

### 只读模式

```javascript
const editorConfig = {
  readOnly: true
}
```

### 最大长度限制

```javascript
const editorConfig = {
  maxLength: 10000
}
```

## 参考资源

- [wangEditor 官方文档](https://www.wangeditor.com/v5/)
- [Vue3 集成文档](https://www.wangeditor.com/v5/for-frame.html#vue3)
- [API 文档](https://www.wangeditor.com/v5/API.html)
- [菜单配置](https://www.wangeditor.com/v5/menu-config.html)
