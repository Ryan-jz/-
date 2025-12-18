# GF Admin 前台网站

基于 Vue 3 + Vite 开发的前台展示网站

## 技术栈

- **框架**: Vue 3.4
- **构建工具**: Vite 5.0
- **路由**: Vue Router 4.2
- **状态管理**: Pinia 2.1
- **HTTP 客户端**: Axios 1.6
- **样式**: SCSS

## 项目结构

```
frontend-web/
├── public/              # 静态资源
├── src/
│   ├── api/            # API 接口
│   ├── assets/         # 资源文件
│   ├── components/     # 公共组件
│   ├── router/         # 路由配置
│   ├── stores/         # 状态管理
│   ├── styles/         # 全局样式
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
cd frontend-web
npm install
```

### 2. 启动开发服务器

```bash
npm run dev
```

访问 `http://localhost:3001`

### 3. 构建生产版本

```bash
npm run build
```

## 页面说明

- **首页**: 展示系统介绍和核心特性
- **关于我们**: 项目介绍和技术栈说明
- **联系我们**: 联系方式展示

## 许可证

MIT License
