# GF Admin 全栈管理系统

基于 GoFrame v2 + Vue 3 + MySQL 8 开发的现代化全栈管理系统

## 项目简介

这是一个完整的企业级后台管理系统解决方案，包含后端 API、管理后台和前台展示三个部分。采用前后端分离架构，遵循工程化开发规范，代码注释完整。

## 技术栈

### 后端
- **框架**: GoFrame v2.6.4
- **数据库**: MySQL 8.0
- **缓存**: Redis
- **认证**: JWT Token
- **ORM**: GoFrame ORM

### 前端管理后台
- **框架**: Vue 3.4
- **构建工具**: Vite 5.0
- **UI 组件**: Element Plus 2.5
- **状态管理**: Pinia 2.1
- **路由**: Vue Router 4.2
- **HTTP**: Axios 1.6

### 前台网站
- **框架**: Vue 3.4
- **构建工具**: Vite 5.0
- **路由**: Vue Router 4.2
- **状态管理**: Pinia 2.1

## 项目结构

```
.
├── backend/                 # 后端项目
│   ├── api/                # API 接口定义
│   ├── hack/               # 工具配置
│   ├── internal/           # 内部代码
│   │   ├── cmd/           # 命令行入口
│   │   ├── consts/        # 常量定义
│   │   ├── controller/    # 控制器层
│   │   ├── model/         # 数据模型
│   │   └── service/       # 业务逻辑层
│   ├── manifest/          # 配置文件
│   │   ├── config/       # 应用配置
│   │   └── sql/          # SQL 脚本
│   ├── utility/           # 工具函数
│   ├── go.mod            # Go 模块文件
│   └── main.go           # 程序入口
│
├── frontend-admin/         # 管理后台前端
│   ├── src/
│   │   ├── api/          # API 接口
│   │   ├── layout/       # 布局组件
│   │   ├── router/       # 路由配置
│   │   ├── stores/       # 状态管理
│   │   ├── utils/        # 工具函数
│   │   └── views/        # 页面组件
│   ├── index.html        # HTML 模板
│   ├── vite.config.js    # Vite 配置
│   └── package.json      # 项目配置
│
└── frontend-web/          # 前台网站
    ├── src/
    │   ├── router/       # 路由配置
    │   ├── stores/       # 状态管理
    │   └── views/        # 页面组件
    ├── index.html        # HTML 模板
    ├── vite.config.js    # Vite 配置
    └── package.json      # 项目配置
```

## 功能特性

### 已实现功能
- ✅ 用户登录/登出
- ✅ JWT Token 认证
- ✅ 用户管理（增删改查）
- ✅ 权限控制
- ✅ 响应式布局
- ✅ 分页查询
- ✅ 数据验证

### 待实现功能
- ⏳ 角色管理
- ⏳ 菜单管理
- ⏳ 部门管理
- ⏳ 字典管理
- ⏳ 日志管理
- ⏳ 文件上传
- ⏳ 数据导出

## 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+

### 1. 数据库初始化

```bash
# 创建数据库并导入初始数据
mysql -u root -p < backend/manifest/sql/init.sql
```

默认管理员账号：
- 用户名: admin
- 密码: admin123

### 2. 启动后端服务

```bash
cd backend

# 安装依赖
go mod tidy

# 修改配置文件
# 编辑 manifest/config/config.yaml，配置数据库连接信息

# 启动服务
go run main.go
```

后端服务将在 `http://localhost:8000` 启动

### 3. 启动管理后台

```bash
cd frontend-admin

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

管理后台将在 `http://localhost:3000` 启动

### 4. 启动前台网站

```bash
cd frontend-web

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前台网站将在 `http://localhost:3001` 启动

## 配置说明

### 后端配置

编辑 `backend/manifest/config/config.yaml`：

```yaml
# 数据库配置
database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/bai_jin_yan"
    
# Redis 配置
redis:
  default:
    address: "127.0.0.1:6379"
    
# JWT 配置
jwt:
  signingKey: "your-secret-key"
  expireTime: 86400
```

### 前端配置

前端项目通过 Vite 代理配置访问后端 API，无需额外配置。

## 开发规范

### 后端规范

1. **分层架构**
   - Controller: 处理 HTTP 请求
   - Service: 业务逻辑处理
   - DAO: 数据访问层

2. **命名规范**
   - 包名: 小写
   - 文件名: 小写下划线
   - 结构体: 大驼峰
   - 函数: 大驼峰（公开）/ 小驼峰（私有）

3. **注释规范**
   - 所有公开的函数、结构体必须添加注释
   - 注释格式: `// FunctionName 函数说明`

### 前端规范

1. **目录命名**
   - 文件夹: 小写短横线（kebab-case）
   - 组件文件: 大驼峰（PascalCase）
   - 工具文件: 小驼峰（camelCase）

2. **代码规范**
   - 使用 Vue 3 Composition API
   - 使用 `<script setup>` 语法
   - 组件必须添加注释说明

## API 文档

### 认证接口

#### 登录
```
POST /login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

#### 获取用户信息
```
GET /getInfo
Authorization: Bearer {token}
```

### 用户管理接口

所有管理接口需要在请求头中携带 JWT Token。

#### 获取用户列表
```
GET /admin/user/list?page=1&pageSize=10
Authorization: Bearer {token}
```

#### 创建用户
```
POST /admin/user/create
Authorization: Bearer {token}
Content-Type: application/json

{
  "userName": "test",
  "nickName": "测试用户",
  "password": "123456",
  "email": "test@example.com",
  "phonenumber": "13800000000",
  "sex": "0",
  "status": 0
}
```

## 部署说明

### 后端部署

```bash
cd backend

# 编译
go build -o gf-admin main.go

# 运行
./gf-admin
```

### 前端部署

```bash
# 管理后台
cd frontend-admin
npm run build

# 前台网站
cd frontend-web
npm run build
```

构建产物在 `dist` 目录，可部署到 Nginx 等 Web 服务器。

## 常见问题

### 1. 数据库连接失败

检查 `config.yaml` 中的数据库配置是否正确，确保 MySQL 服务已启动。

### 2. 前端跨域问题

开发环境已配置 Vite 代理，生产环境需要在后端配置 CORS 或使用 Nginx 反向代理。

### 3. JWT Token 过期

Token 默认有效期为 24 小时，过期后需要重新登录。

## 贡献指南

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 联系方式

- 邮箱: admin@example.com
- 项目地址: https://github.com/your-repo/gf-admin
