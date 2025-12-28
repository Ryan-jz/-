# 国际化快速开始

## 🚀 5分钟快速部署

### 步骤 1: 数据库迁移

```bash
# Windows CMD
cd backend
mysql -u root -p gf_admin < manifest\sql\i18n_migration.sql
```

输入数据库密码后，脚本会自动：
- ✅ 创建3个国际化表
- ✅ 迁移现有数据
- ✅ 插入示例翻译

### 步骤 2: 启动后端

后端配置已更新，直接启动即可：

```bash
cd backend
.\bai-jin-yan-api.exe
```

或使用编译脚本：
```bash
.\build.bat
```

### 步骤 3: 安装前端依赖

```bash
cd frontend-web
pnpm install
```

依赖已在 `package.json` 中，会自动安装 `vue-i18n`。

### 步骤 4: 启动前端

```bash
# 在 frontend-web 目录
pnpm dev
```

访问 http://localhost:5173

## ✅ 验证安装

### 测试后端API

打开浏览器开发者工具，在Console中执行：

```javascript
// 测试中文
fetch('http://localhost:8000/api/v1/products/1', {
  headers: { 'Accept-Language': 'zh-CN' }
}).then(r => r.json()).then(console.log)

// 测试英文
fetch('http://localhost:8000/api/v1/products/1', {
  headers: { 'Accept-Language': 'en-US' }
}).then(r => r.json()).then(console.log)

// 测试德文
fetch('http://localhost:8000/api/v1/products/1', {
  headers: { 'Accept-Language': 'de-DE' }
}).then(r => r.json()).then(console.log)
```

应该看到不同语言的产品名称和描述。

### 测试前端切换

1. 在页面右上角找到语言切换器（下拉框）
2. 切换到 "English"
3. 页面应该刷新并显示英文内容
4. 切换到 "Deutsch"
5. 页面应该显示德文内容

## 📝 常见使用场景

### 场景1: 在现有页面添加语言切换

```vue
<template>
  <div class="header">
    <nav>
      <!-- 现有导航 -->
    </nav>
    
    <!-- 添加这一行 -->
    <LanguageSwitcher />
  </div>
</template>

<script setup>
import LanguageSwitcher from '@/components/LanguageSwitcher.vue'
</script>
```

### 场景2: 翻译页面文本

```vue
<template>
  <div>
    <!-- 之前 -->
    <h1>产品中心</h1>
    <button>加入购物车</button>
    
    <!-- 之后 -->
    <h1>{{ $t('product.title') }}</h1>
    <button>{{ $t('product.addToCart') }}</button>
  </div>
</template>
```

### 场景3: 在管理后台编辑多语言内容

```vue
<template>
  <el-form>
    <!-- 基础信息 -->
    <el-form-item label="价格">
      <el-input-number v-model="form.price" />
    </el-form-item>
    
    <!-- 添加国际化编辑器 -->
    <ProductI18nEditor v-model="form.i18n" />
  </el-form>
</template>

<script setup>
import ProductI18nEditor from '@/components/I18nEditor/ProductI18nEditor.vue'
</script>
```

## 🎯 下一步

- 查看 [完整实施指南](./I18N_IMPLEMENTATION_GUIDE.md) 了解详细信息
- 查看示例文件：
  - `frontend-web/src/views/ProductDetailExample.vue` - 前端使用示例
  - `frontend-admin/src/views/ProductEditExample.vue` - 管理后台示例
- 根据需要添加更多翻译到 `frontend-web/src/locales/*.js`

## ❓ 遇到问题？

### 问题：API返回的还是中文

**解决：** 检查请求头是否包含 `Accept-Language`

```javascript
// 在浏览器Console中检查
console.log(document.querySelector('html').lang)
```

### 问题：切换语言后页面没变化

**解决：** 清除浏览器缓存和localStorage

```javascript
// 在浏览器Console中执行
localStorage.clear()
location.reload()
```

### 问题：管理后台保存失败

**解决：** 确保数据库迁移脚本已执行

```sql
-- 检查表是否存在
SHOW TABLES LIKE '%_i18n';

-- 应该看到：
-- product_category_i18n
-- product_i18n
-- recipe_i18n
```

## 📞 需要帮助？

查看详细文档：[I18N_IMPLEMENTATION_GUIDE.md](./I18N_IMPLEMENTATION_GUIDE.md)
