# 产品和食谱国际化实施计划

## 当前状态

✅ **已完成**：
- Banner 轮播图国际化（完整实现）
- 数据库国际化表已创建（product_i18n, product_category_i18n, recipe_i18n）
- 产品国际化编辑器组件（ProductI18nEditor.vue）
- 食谱国际化编辑器组件（RecipeI18nEditor.vue）

⚠️ **待完成**：
- 产品管理页面集成国际化编辑
- 产品分类管理页面集成国际化编辑
- 食谱管理页面集成国际化编辑
- 后端 API 支持国际化数据

## 实施步骤

### 第一阶段：产品分类国际化

#### 1.1 前端改造
文件：`frontend-admin/src/views/product/category.vue`

需要修改：
- 添加国际化编辑器组件
- 移除单语言字段（name、nameEn）
- 添加 i18n 数据结构
- 更新表单提交逻辑

#### 1.2 后端改造
文件：
- `backend/internal/controller/api/product.go`
- `backend/internal/service/product.go`

需要修改：
- CategoryCreate() - 接收并保存国际化数据
- CategoryUpdate() - 更新国际化数据
- GetCategoryList() - 根据 locale 返回对应语言

### 第二阶段：产品国际化

#### 2.1 前端改造
文件：`frontend-admin/src/views/product/list.vue`

需要修改：
- 集成 ProductI18nEditor 组件
- 移除单语言字段（name、nameEn、subtitle、description等）
- 添加 i18n 数据结构
- 更新表单提交逻辑
- 添加获取国际化数据的 API 调用

#### 2.2 后端改造
文件：
- `backend/internal/controller/api/product.go`
- `backend/internal/service/product.go`

需要修改：
- Create() - 接收并保存国际化数据（使用事务）
- Update() - 更新国际化数据（使用事务）
- GetList() - 根据 locale 返回对应语言
- GetDetail() - 返回所有语言版本
- 新增 GetI18n() - 获取产品国际化数据

#### 2.3 新增 API
- `GET /api/v1/product/:id/i18n` - 获取产品国际化数据

### 第三阶段：食谱国际化

#### 3.1 前端改造
文件：`frontend-admin/src/views/recipe/list.vue`

需要修改：
- 集成 RecipeI18nEditor 组件
- 移除单语言字段（name、nameEn、subtitle、description等）
- 添加 i18n 数据结构
- 更新表单提交逻辑
- 添加获取国际化数据的 API 调用

#### 3.2 后端改造
文件：
- `backend/internal/controller/api/recipe.go`（需要查看）
- `backend/internal/service/recipe.go`（需要查看）

需要修改：
- Create() - 接收并保存国际化数据（使用事务）
- Update() - 更新国际化数据（使用事务）
- GetList() - 根据 locale 返回对应语言
- GetDetail() - 返回所有语言版本
- 新增 GetI18n() - 获取食谱国际化数据

#### 3.3 新增 API
- `GET /api/v1/recipe/:id/i18n` - 获取食谱国际化数据

## 数据结构设计

### 产品分类 i18n 数据结构
```javascript
{
  'zh-CN': {
    name: '调味盐系列',
    description: 'BBQ调味盐、香草调味盐等特色调味盐产品'
  },
  'en-US': {
    name: 'Seasoning Salt Series',
    description: 'BBQ seasoning salt, herb seasoning salt and other specialty products'
  },
  'de-DE': {
    name: 'Gewürzsalz-Serie',
    description: 'BBQ-Gewürzsalz, Kräutergewürzsalz und andere Spezialprodukte'
  }
}
```

### 产品 i18n 数据结构
```javascript
{
  'zh-CN': {
    name: '阿尔卑斯粗盐',
    subtitle: '经典粗盐，适合烹饪调味',
    description: '产品描述...',
    ingredients: '成分说明...',
    usage: '使用方法...',
    features: ['特点1', '特点2']
  },
  'en-US': { ... },
  'de-DE': { ... }
}
```

### 食谱 i18n 数据结构
```javascript
{
  'zh-CN': {
    name: '喜马拉雅盐烤鸡',
    subtitle: '外酥里嫩，咸香四溢',
    description: '简介...',
    ingredients: '食材列表...',
    tags: '烤鸡,主菜,家常菜'
  },
  'en-US': { ... },
  'de-DE': { ... }
}
```

## 实施优先级

### 高优先级（建议先做）
1. ✅ Banner 国际化（已完成）
2. 🔄 产品分类国际化（简单，字段少）
3. 🔄 产品国际化（核心功能）

### 中优先级
4. 🔄 食谱国际化（功能相对独立）

### 低优先级
5. 前端网站国际化切换（需要前端路由和语言切换功能）

## 技术要点

### 1. 事务处理
产品和食谱的创建/更新需要使用事务，确保主表和国际化表数据一致性：

```go
err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
    // 1. 插入/更新主表
    // 2. 插入/更新国际化表
    return nil
})
```

### 2. 数据迁移
现有数据需要迁移到国际化表（已在 i18n_migration.sql 中完成）

### 3. API 兼容性
- 列表接口：添加 `locale` 参数，默认 `zh-CN`
- 详情接口：返回所有语言版本
- 创建/更新接口：接收 `i18n` 对象

### 4. 前端组件复用
- BannerI18nEditor.vue（已完成）
- ProductI18nEditor.vue（已存在）
- RecipeI18nEditor.vue（已创建）

## 测试计划

### 单元测试
- [ ] 产品分类国际化 CRUD
- [ ] 产品国际化 CRUD
- [ ] 食谱国际化 CRUD

### 集成测试
- [ ] 创建产品并验证国际化数据
- [ ] 更新产品并验证国际化数据
- [ ] 切换语言查看列表
- [ ] 数据迁移验证

### 用户测试
- [ ] 后台管理界面操作流程
- [ ] 多语言内容编辑
- [ ] 前端网站多语言显示

## 预计工作量

- 产品分类国际化：2-3 小时
- 产品国际化：4-5 小时
- 食谱国际化：4-5 小时
- 测试和调试：2-3 小时

**总计：12-16 小时**

## 下一步行动

### 立即可做
1. 查看并确认产品分类管理页面（category.vue）
2. 开始产品分类国际化改造
3. 参考 Banner 的实现方式

### 需要确认
1. 是否需要保留原有的 name_en 字段（建议废弃）
2. 前端网站是否需要语言切换功能
3. 是否需要为管理员提供批量翻译功能

## 参考文档

- Banner 国际化实现：`BANNER_I18N_GUIDE.md`
- 数据库迁移脚本：`backend/manifest/sql/i18n_migration.sql`
- Banner 国际化脚本：`backend/manifest/sql/banner_i18n_migration.sql`
