# 产品管理系统使用指南

## 功能概述

后台管理系统已集成完整的产品管理功能，包括：
- 产品列表管理（增删改查）
- 产品分类管理
- 图片上传功能
- 搜索和筛选

## 菜单导航

登录后台管理系统后，在左侧菜单可以看到：

```
📊 首页
📦 产品管理
  ├─ 📋 产品列表
  └─ 📁 分类管理
⚙️ 系统管理
  ├─ 👤 用户管理
  ├─ 👥 角色管理
  └─ 📋 菜单管理
```

## 使用流程

### 1. 创建产品分类

**路径：** 产品管理 > 分类管理

**步骤：**
1. 点击"新增分类"按钮
2. 填写分类信息：
   - 分类名称（必填）：例如"高山盐"
   - 英文名称：例如"Alpine Salt"
   - 分类标识（必填）：例如"alpine-salt"（用于URL）
   - 分类描述：简要描述该分类
   - 排序：数字越小越靠前
   - 状态：启用/禁用
3. 点击"确定"保存

**功能：**
- 编辑：修改分类信息
- 删除：删除分类（如果分类下有产品则无法删除）

### 2. 管理产品

**路径：** 产品管理 > 产品列表

#### 2.1 搜索和筛选

- **产品分类**：下拉选择分类进行筛选
- **关键词**：输入产品名称搜索
- **状态**：筛选上架/下架的产品
- 点击"搜索"按钮执行查询
- 点击"重置"清空筛选条件

#### 2.2 新增产品

1. 点击"新增产品"按钮
2. 填写产品信息：

**基本信息：**
- 产品分类（必填）：选择所属分类
- 产品名称（必填）：例如"喜马拉雅粉盐"
- 英文名称：例如"Himalayan Pink Salt"
- 副标题：产品的简短描述

**详细信息：**
- 产品描述：详细介绍产品特点
- 主图片：点击上传产品主图（支持jpg、png、gif、webp，最大5MB）
- 价格：产品售价
- 库存：库存数量
- 规格/重量：例如"500g"

**附加信息：**
- 成分说明：产品成分列表
- 使用方法：如何使用该产品
- 排序：数字越小越靠前显示
- 状态：上架/下架

3. 点击"确定"保存

#### 2.3 编辑产品

1. 在产品列表中找到要编辑的产品
2. 点击"编辑"按钮
3. 修改产品信息
4. 点击"确定"保存

#### 2.4 删除产品

1. 在产品列表中找到要删除的产品
2. 点击"删除"按钮
3. 确认删除操作

### 3. 图片上传

**支持的格式：** jpg、jpeg、png、gif、webp
**文件大小限制：** 5MB

**使用方法：**
1. 点击上传区域
2. 选择图片文件
3. 等待上传完成
4. 可以预览或删除已上传的图片

**图片存储：**
- 图片保存在服务器：`backend/public/uploads/images/`
- 按日期分目录：`年/月/日/`
- 访问路径：`/uploads/images/2024/12/25/文件名.jpg`

## API 接口

### 分类管理

```javascript
// 获取分类列表
GET /api/v1/product/category/list?status=1

// 创建分类
POST /api/v1/product/category/create
{
  "name": "高山盐",
  "nameEn": "Alpine Salt",
  "slug": "alpine-salt",
  "description": "来自高山的天然盐",
  "sortOrder": 0,
  "status": 1
}

// 更新分类
PUT /api/v1/product/category/update
{
  "id": 1,
  "name": "高山盐",
  ...
}

// 删除分类
DELETE /api/v1/product/category/delete/1
```

### 产品管理

```javascript
// 获取产品列表
GET /api/v1/product/list?categoryId=1&keyword=盐&status=1&page=1&pageSize=10

// 获取产品详情
GET /api/v1/product/detail/1

// 创建产品
POST /api/v1/product/create
{
  "categoryId": 1,
  "name": "喜马拉雅粉盐",
  "nameEn": "Himalayan Pink Salt",
  "image": "/uploads/images/2024/12/25/xxx.jpg",
  "price": 29.9,
  "stock": 100,
  ...
}

// 更新产品
PUT /api/v1/product/update
{
  "id": 1,
  "categoryId": 1,
  ...
}

// 删除产品
DELETE /api/v1/product/delete/1
```

### 图片上传

```javascript
// 上传图片
POST /api/v1/upload/image
Content-Type: multipart/form-data
file: [图片文件]

// 响应
{
  "code": 0,
  "message": "上传成功",
  "data": {
    "url": "/uploads/images/2024/12/25/20241225150405_abc12345.jpg"
  }
}

// 删除文件
POST /api/v1/upload/delete
{
  "url": "/uploads/images/2024/12/25/20241225150405_abc12345.jpg"
}
```

## 前端组件

### ImageUpload 组件

可复用的图片上传组件，支持：
- 拖拽上传
- 图片预览
- 删除图片
- 文件类型和大小验证

**使用示例：**

```vue
<template>
  <el-form-item label="产品图片">
    <ImageUpload v-model="form.image" placeholder="上传产品图片" />
  </el-form-item>
</template>

<script setup>
import { ref } from 'vue'
import ImageUpload from '@/components/ImageUpload.vue'

const form = ref({
  image: ''
})
</script>
```

## 注意事项

1. **权限控制**：所有产品管理接口都需要登录认证
2. **图片管理**：定期清理不用的图片文件，避免占用过多磁盘空间
3. **分类删除**：删除分类前需要先删除或转移该分类下的所有产品
4. **数据备份**：重要数据建议定期备份
5. **图片优化**：上传前建议压缩图片，提高加载速度

## 启动步骤

### 1. 启动后端服务

```bash
cd backend
go run main.go
```

服务器将运行在：http://localhost:8000

### 2. 启动前端管理系统

```bash
cd frontend-admin
pnpm install  # 首次运行需要安装依赖
pnpm dev
```

管理系统将运行在：http://localhost:3000

### 3. 登录系统

- 访问：http://localhost:3000
- 默认账号：admin
- 默认密码：admin

## 常见问题

### 1. 图片上传失败
- 检查文件大小是否超过5MB
- 检查文件格式是否支持
- 确认 `backend/public/uploads` 目录有写入权限

### 2. 产品列表为空
- 检查数据库是否有数据
- 查看后端日志是否有错误
- 确认API接口是否正常

### 3. 分类无法删除
- 检查该分类下是否还有产品
- 先删除或转移产品后再删除分类

### 4. 图片无法显示
- 检查图片路径是否正确
- 确认静态文件服务是否正常
- 查看浏览器控制台是否有404错误
