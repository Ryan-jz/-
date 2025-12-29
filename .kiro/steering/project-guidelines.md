---
inclusion: always
---

# 项目开发规范和指南

## 项目概述

这是一个全栈项目，包含：
- **Backend**: GoFrame v2 框架的 Go 后端 API
- **Frontend-Admin**: Vue 3 + Element Plus 管理后台
- **Frontend-Web**: Vue 3 + UnoCSS 前端网站

## 技术栈

### 后端 (backend/)
- **框架**: GoFrame v2.6.4
- **Go 版本**: 1.20.14
- **数据库**: MySQL
- **ORM**: GoFrame ORM + DAO 自动生成
- **架构**: MVC 分层架构
  - `api/`: API 定义和请求/响应结构
  - `internal/controller/`: 控制器层
  - `internal/service/`: 业务逻辑层
  - `internal/dao/`: 数据访问层（自动生成）
  - `internal/model/`: 数据模型（自动生成）
  - `internal/middleware/`: 中间件
  - `internal/router/`: 路由配置

### 前端管理后台 (frontend-admin/)
- **框架**: Vue 3 + Vite
- **UI 库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router
- **HTTP**: Axios

### 前端网站 (frontend-web/)
- **框架**: Vue 3 + Vite
- **CSS**: UnoCSS
- **国际化**: Vue I18n
- **HTTP**: Axios

## 开发规范

### 后端开发规范

#### 1. 使用 DAO 自动生成
```bash
# 数据库结构变更后，重新生成 DAO
cd backend
gf gen dao
```

#### 2. 查询规范
- **简单查询**: 使用生成的 DAO
  ```go
  import "gf-admin/internal/dao"
  
  dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Status, 1).Scan(&list)
  ```

- **复杂查询（带 JOIN）**: 使用原生 SQL
  ```go
  sql := `
      SELECT b.*, bi.title, bi.description
      FROM banner b
      LEFT JOIN banner_i18n bi ON b.banner_id = bi.banner_id AND bi.locale = ?
      WHERE b.status = ?
  `
  g.DB().Ctx(ctx).Raw(sql, locale, status).Scan(&list)
  ```

#### 3. 国际化支持
- 所有 API 都通过 i18n 中间件处理
- 从 `Accept-Language` 请求头获取语言
- 支持语言: zh-CN, en-US, de-DE
- 使用 `middleware.GetLocale(ctx)` 获取当前语言

#### 4. 事务处理
```go
err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
    // 使用 tx 执行数据库操作
    _, err := dao.Banner.Ctx(ctx).TX(tx).Data(data).Insert()
    return err
})
```

#### 5. 编译和部署
```bash
# 编译
cd backend
.\build.bat

# 运行
.\gf-admin.exe
```

### 前端开发规范

#### 1. API 请求
- 统一使用 `src/utils/request.js` 封装的 axios 实例
- 自动添加 `Accept-Language` 请求头
- 自动处理错误和 token

#### 2. 国际化编辑器
- 使用 `src/components/I18nEditor/` 下的组件
- 支持三种语言的编辑
- 组件：
  - `BannerI18nEditor.vue`
  - `CategoryI18nEditor.vue`
  - `ProductI18nEditor.vue`
  - `RecipeI18nEditor.vue`

#### 3. 语言切换（仅 frontend-web）
- 使用 `LanguageSwitcher.vue` 组件
- 语言存储在 localStorage
- 自动更新所有 API 请求的 Accept-Language

### 数据库规范

#### 1. 国际化表结构
每个需要国际化的表都有对应的 `_i18n` 表：
- `banner` → `banner_i18n`
- `product` → `product_i18n`
- `product_category` → `product_category_i18n`
- `recipe` → `recipe_i18n`

#### 2. i18n 表字段
- `id`: 主键
- `{table}_id`: 关联主表 ID
- `locale`: 语言代码 (zh-CN, en-US, de-DE)
- 其他可翻译字段
- `created_at`, `updated_at`: 时间戳

#### 3. 查询国际化数据
```sql
SELECT 
    main.*,
    COALESCE(i18n.field, '') as field
FROM main_table main
LEFT JOIN main_table_i18n i18n 
    ON main.id = i18n.main_id 
    AND i18n.locale = ?
```

## 常见任务

### 添加新的 API 接口
1. 在 `backend/api/v1/` 定义请求/响应结构
2. 在 `backend/internal/service/` 实现业务逻辑
3. 在 `backend/internal/controller/api/` 添加控制器方法
4. 在 `backend/internal/router/router.go` 注册路由
5. 重新编译: `cd backend && .\build.bat`

### 添加新的数据表
1. 在数据库中创建表
2. 运行 `gf gen dao` 生成 DAO 和 Model
3. 实现 Service 层业务逻辑
4. 添加 Controller 和路由

### 添加国际化支持
1. 创建 `{table}_i18n` 表
2. 运行 `gf gen dao`
3. 在 Service 层使用 LEFT JOIN 查询
4. 在前端添加对应的 I18nEditor 组件

## 注意事项

### 后端
- ⚠️ **不要手动修改** `internal/dao/internal/` 和 `internal/model/` 下的文件
- ⚠️ 使用 Go 1.20.14，不要升级到更高版本
- ⚠️ sql查询尽量都要使用gf的ORM，不要使用原生sql
- ✅ 所有 API 都要通过 i18n 中间件
- ✅ 使用事务保证数据一致性

### 前端
- ⚠️ API 基础路径根据环境自动切换
- ⚠️ 生产环境使用公网 IP: `http://8.133.175.112:8000`
- ✅ 使用 Element Plus 组件库（管理后台）
- ✅ 使用 UnoCSS 原子化 CSS（前端网站）

## 调试技巧

### 后端调试
1. 查看日志输出（已添加 i18n 中间件日志）
2. 使用 Postman 测试 API
3. 检查 SQL 执行日志

### 前端调试
1. 浏览器开发者工具 → Network 查看请求
2. 检查 `Accept-Language` 请求头
3. 查看 Console 日志

## 相关文档

- [GoFrame 官方文档](https://goframe.org)
- [Element Plus 文档](https://element-plus.org)
- [Vue 3 文档](https://vuejs.org)
- [UnoCSS 文档](https://unocss.dev)
