# 国际化实施指南

本项目已实现完整的中英德三语国际化支持。

## 📁 文件结构

```
backend/
├── manifest/
│   ├── config/config.yaml          # 添加了i18n配置
│   └── sql/
│       └── i18n_migration.sql      # 国际化数据库迁移脚本
├── internal/
│   ├── middleware/
│   │   └── i18n.go                 # 国际化中间件
│   └── service/
│       ├── product_i18n.go         # 产品国际化服务
│       └── recipe_i18n.go          # 食谱国际化服务

frontend-web/
├── src/
│   ├── locales/                    # 语言文件
│   │   ├── zh-CN.js               # 中文
│   │   ├── en-US.js               # 英文
│   │   ├── de-DE.js               # 德文
│   │   └── index.js               # i18n配置
│   ├── components/
│   │   └── LanguageSwitcher.vue   # 语言切换组件
│   ├── utils/
│   │   └── request.js             # 已添加语言头
│   └── main.js                    # 已注册i18n

frontend-admin/
└── src/
    └── components/
        └── I18nEditor/            # 国际化编辑组件
            ├── ProductI18nEditor.vue
            └── RecipeI18nEditor.vue
```

## 🗄️ 数据库部署

### 1. 执行迁移脚本

```bash
# 进入backend目录
cd backend

# 执行SQL脚本
mysql -u root -p gf_admin < manifest/sql/i18n_migration.sql
```

这个脚本会：
- 创建 `product_category_i18n`、`product_i18n`、`recipe_i18n` 三个国际化表
- 自动迁移现有数据到国际化表（中文和英文）
- 插入德语翻译示例数据

### 2. 验证数据

```sql
-- 查看产品国际化数据
SELECT * FROM product_i18n WHERE product_id = 1;

-- 查看食谱国际化数据
SELECT * FROM recipe_i18n WHERE recipe_id = 1;
```

## 🔧 后端使用

### 1. 在路由中添加中间件

```go
// backend/internal/router/router.go
import "gf-admin/internal/middleware"

// 在需要国际化的路由组中添加中间件
group.Middleware(middleware.I18n)
```

### 2. 在Controller中使用国际化Service

```go
// 获取产品列表（自动返回对应语言）
productI18nService := &service.ProductI18nService{}
products, total, err := productI18nService.GetProductListI18n(ctx, categoryId, keyword, status, page, pageSize)

// 获取产品详情（自动返回对应语言）
detail, err := productI18nService.GetProductDetailI18n(ctx, productId)

// 保存产品国际化数据（管理后台）
i18nData := map[string]g.Map{
    "zh-CN": {
        "name": "产品名称",
        "description": "产品描述",
        // ...
    },
    "en-US": {
        "name": "Product Name",
        "description": "Product Description",
        // ...
    },
    "de-DE": {
        "name": "Produktname",
        "description": "Produktbeschreibung",
        // ...
    },
}
err := productI18nService.BatchSaveProductI18n(ctx, productId, i18nData)
```

### 3. 语言检测优先级

中间件会按以下优先级检测语言：
1. URL参数：`?lang=zh-CN`
2. HTTP Header：`Accept-Language: zh-CN`
3. Cookie：`locale=zh-CN`
4. 默认语言：`zh-CN`（配置文件中设置）

## 🎨 前端Web使用

### 1. 在组件中使用翻译

```vue
<template>
  <div>
    <!-- 使用 $t 函数翻译静态文本 -->
    <h1>{{ $t('common.home') }}</h1>
    <button>{{ $t('product.addToCart') }}</button>
    
    <!-- 动态内容直接使用API返回的数据 -->
    <div class="product">
      <h2>{{ product.name }}</h2>
      <p>{{ product.description }}</p>
    </div>
  </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

// 也可以在JS中使用
console.log(t('common.loading'))
</script>
```

### 2. 添加语言切换器

```vue
<template>
  <div class="header">
    <nav>
      <!-- 其他导航 -->
    </nav>
    
    <!-- 添加语言切换器 -->
    <LanguageSwitcher />
  </div>
</template>

<script setup>
import LanguageSwitcher from '@/components/LanguageSwitcher.vue'
</script>
```

### 3. 在API调用中自动携带语言

已在 `request.js` 中配置，所有API请求会自动携带 `Accept-Language` 头。

## 🛠️ 管理后台使用

### 1. 在产品编辑页面使用

```vue
<template>
  <el-form>
    <!-- 基础信息 -->
    <el-form-item label="分类">
      <el-select v-model="form.category_id" />
    </el-form-item>
    
    <el-form-item label="价格">
      <el-input-number v-model="form.price" />
    </el-form-item>
    
    <!-- 国际化内容编辑 -->
    <ProductI18nEditor v-model="form.i18n" />
    
    <el-button @click="handleSubmit">保存</el-button>
  </el-form>
</template>

<script setup>
import { ref } from 'vue'
import ProductI18nEditor from '@/components/I18nEditor/ProductI18nEditor.vue'

const form = ref({
  category_id: null,
  price: 0,
  i18n: {
    'zh-CN': {
      name: '',
      description: '',
      // ...
    },
    'en-US': {
      name: '',
      description: '',
      // ...
    },
    'de-DE': {
      name: '',
      description: '',
      // ...
    }
  }
})

const handleSubmit = async () => {
  // 提交到后端
  await api.saveProduct(form.value)
}
</script>
```

### 2. 在食谱编辑页面使用

```vue
<template>
  <el-form>
    <!-- 基础信息 -->
    <el-form-item label="烹饪时间">
      <el-input-number v-model="form.cooking_time" />
    </el-form-item>
    
    <!-- 国际化内容编辑 -->
    <RecipeI18nEditor v-model="form.i18n" />
    
    <el-button @click="handleSubmit">保存</el-button>
  </el-form>
</template>

<script setup>
import RecipeI18nEditor from '@/components/I18nEditor/RecipeI18nEditor.vue'
</script>
```

## 📝 添加新的翻译

### 前端添加新翻译

编辑对应的语言文件：

```javascript
// frontend-web/src/locales/zh-CN.js
export default {
  // 添加新的翻译键
  newFeature: {
    title: '新功能标题',
    description: '新功能描述'
  }
}
```

### 后端添加新的国际化表

如果需要为其他实体添加国际化支持：

```sql
-- 创建新的国际化表
CREATE TABLE `banner_i18n` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `banner_id` int(11) unsigned NOT NULL,
  `locale` varchar(10) NOT NULL,
  `title` varchar(200),
  `description` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_banner_locale` (`banner_id`, `locale`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

然后创建对应的Service：

```go
// backend/internal/service/banner_i18n.go
type BannerI18nService struct{}

func (s *BannerI18nService) GetBannerListI18n(ctx context.Context) ([]g.Map, error) {
    locale := middleware.GetLocale(ctx)
    // 实现逻辑...
}
```

## 🧪 测试

### 测试后端API

```bash
# 测试中文
curl -H "Accept-Language: zh-CN" http://localhost:8000/api/v1/products

# 测试英文
curl -H "Accept-Language: en-US" http://localhost:8000/api/v1/products

# 测试德文
curl -H "Accept-Language: de-DE" http://localhost:8000/api/v1/products

# 使用URL参数
curl http://localhost:8000/api/v1/products?lang=de-DE
```

### 测试前端

1. 启动前端项目：
```bash
cd frontend-web
pnpm dev
```

2. 在浏览器中访问，使用语言切换器切换语言
3. 检查API请求头是否包含正确的 `Accept-Language`
4. 验证页面内容是否正确切换

## 🎯 最佳实践

### 1. 翻译内容管理

- **静态UI文本**：使用前端语言文件（`locales/*.js`）
- **动态内容**：使用数据库国际化表，通过API返回

### 2. 回退机制

- 如果某个语言的翻译不存在，系统会自动回退到中文
- 在SQL查询中使用 `COALESCE` 确保总是有内容返回

### 3. 性能优化

- 使用LEFT JOIN而不是多次查询
- 在国际化表上建立索引：`(entity_id, locale)`
- 考虑使用Redis缓存常用翻译

### 4. 内容编辑流程

1. 管理员在后台编辑产品/食谱
2. 使用Tab切换不同语言
3. 填写每个语言的内容
4. 保存时批量更新所有语言的数据

## 🔍 故障排查

### 问题1：API返回的内容没有翻译

**检查：**
- 中间件是否正确添加到路由
- 数据库中是否有对应语言的翻译数据
- 请求头中的 `Accept-Language` 是否正确

### 问题2：前端切换语言后内容没变化

**检查：**
- `request.js` 中是否添加了语言头
- 是否调用了 `setLocale()` 函数
- 浏览器localStorage中的locale值是否正确

### 问题3：管理后台保存失败

**检查：**
- 国际化表是否已创建
- 数据格式是否正确（特别是JSON字段）
- 后端Service是否正确调用

## 📚 相关资源

- [Vue I18n 文档](https://vue-i18n.intlify.dev/)
- [GoFrame 国际化](https://goframe.org/pages/viewpage.action?pageId=1114367)
- [MySQL 多语言设计模式](https://stackoverflow.com/questions/316780/schema-for-a-multilanguage-database)

## ✅ 部署检查清单

- [ ] 执行数据库迁移脚本
- [ ] 验证国际化表数据
- [ ] 更新后端配置文件
- [ ] 在路由中添加i18n中间件
- [ ] 前端安装vue-i18n依赖
- [ ] 测试三种语言的API响应
- [ ] 测试前端语言切换功能
- [ ] 测试管理后台编辑功能
- [ ] 更新生产环境配置

---

**实施完成！** 🎉

现在你的项目已经支持完整的中英德三语国际化。用户可以在前端自由切换语言，管理员可以在后台编辑多语言内容。
