# GF Admin 前端管理后台

基于 Vue 3 + Vite + Element Plus 开发的管理系统前端

## 技术栈

- **框架**: Vue 3.4
- **构建工具**: Vite 5.0
- **UI 组件库**: Element Plus 2.5
- **状态管理**: Pinia 2.1
- **路由**: Vue Router 4.2
- **HTTP 客户端**: Axios 1.6
- **样式**: SCSS

## 项目结构

```
frontend-admin/
├── public/              # 静态资源
├── src/
│   ├── api/            # API 接口
│   ├── assets/         # 资源文件
│   ├── components/     # 公共组件
│   ├── layout/         # 布局组件
│   ├── router/         # 路由配置
│   ├── stores/         # 状态管理
│   ├── styles/         # 全局样式
│   ├── utils/          # 工具函数
│   ├── views/          # 页面组件
│   ├── App.vue         # 根组件
│   └── main.js         # 入口文件
├── index.html          # HTML 模板
├── vite.config.js      # Vite 配置
└── package.json        # 项目配置

```

## 快速开始

### 1. 安装依赖

```bash
cd frontend-admin
npm install
```

### 2. 启动开发服务器

```bash
npm run dev
```

访问 `http://localhost:3000`

### 3. 构建生产版本

```bash
npm run build
```

### 4. 预览生产构建

```bash
npm run preview
```

## 功能特性

### 已实现功能

- ✅ 用户登录/登出
- ✅ JWT Token 认证
- ✅ 路由权限控制
- ✅ 用户管理（增删改查）
- ✅ 响应式布局
- ✅ 侧边栏折叠
- ✅ 面包屑导航

### 待实现功能

- ⏳ 角色管理
- ⏳ 菜单管理
- ⏳ 部门管理
- ⏳ 字典管理
- ⏳ 日志管理
- ⏳ 文件上传
- ⏳ 数据导出

## 开发规范

### 目录命名

- 文件夹：小写短横线分隔（kebab-case）
- 组件文件：大驼峰（PascalCase）
- 工具文件：小驼峰（camelCase）

### 代码规范

- 使用 Vue 3 Composition API
- 使用 `<script setup>` 语法
- 组件必须添加注释说明
- API 接口必须添加 JSDoc 注释

### Git 提交规范

```
feat: 新功能
fix: 修复 bug
docs: 文档更新
style: 代码格式调整
refactor: 代码重构
test: 测试相关
chore: 构建/工具链相关
```

## 配置说明

### 环境变量

创建 `.env.development` 和 `.env.production` 文件：

```env
# API 基础路径
VITE_API_BASE_URL=http://localhost:8000

# 应用标题
VITE_APP_TITLE=GF Admin
```

### 代理配置

在 `vite.config.js` 中配置开发服务器代理：

```js
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8000',
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, '')
    }
  }
}
```

## 默认账号

- 用户名: admin
- 密码: admin123

## 浏览器支持

- Chrome >= 87
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 许可证

MIT License
