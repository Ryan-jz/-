# GF Admin 后端项目

基于 GoFrame v2 框架开发的管理系统后端 API

## 项目结构

```
backend/
├── api/                    # API 接口定义
│   └── v1/                # v1 版本接口
│       ├── admin/         # 管理后台接口定义
│       └── auth/          # 认证接口定义
├── hack/                  # 工具配置
│   └── config.yaml       # GF CLI 工具配置
├── internal/              # 内部代码
│   ├── cmd/              # 命令行入口
│   ├── consts/           # 常量定义
│   ├── controller/       # 控制器层
│   │   ├── admin/       # 管理后台控制器
│   │   └── api/         # 公开 API 控制器
│   ├── model/            # 数据模型
│   │   ├── do/          # 数据操作对象
│   │   └── entity/      # 数据库实体
│   ├── packed/           # 打包资源
│   └── service/          # 业务逻辑层
├── manifest/              # 配置文件
│   ├── config/           # 应用配置
│   └── sql/              # SQL 脚本
├── utility/               # 工具函数
│   └── response/         # 响应处理
├── go.mod                # Go 模块文件
└── main.go               # 程序入口

```

## 技术栈

- **框架**: GoFrame v2.6.4
- **数据库**: MySQL 8.0
- **缓存**: Redis
- **认证**: JWT

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod tidy
```

### 2. 配置数据库

修改 `manifest/config/config.yaml` 中的数据库配置：

```yaml
database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/gf_admin"
```

### 3. 初始化数据库

执行 SQL 脚本：

```bash
mysql -u root -p < manifest/sql/init.sql
```

默认管理员账号：
- 用户名: admin
- 密码: admin123

### 4. 启动服务

```bash
go run main.go
```

服务将在 `http://localhost:8000` 启动

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

### 管理后台接口

所有管理后台接口都需要在请求头中携带 JWT Token：
```
Authorization: Bearer {token}
```

#### 用户管理
- `GET /admin/user/list` - 获取用户列表
- `GET /admin/user/detail` - 获取用户详情
- `POST /admin/user/create` - 创建用户
- `PUT /admin/user/update` - 更新用户
- `DELETE /admin/user/delete` - 删除用户

## 开发指南

### 代码生成

使用 GF CLI 工具生成 DAO 代码：

```bash
gf gen dao
```

### 项目规范

1. **分层架构**
   - Controller: 处理 HTTP 请求
   - Service: 业务逻辑处理
   - DAO: 数据访问层

2. **命名规范**
   - 文件名: 小写下划线
   - 包名: 小写
   - 结构体: 大驼峰
   - 函数: 大驼峰（公开）/ 小驼峰（私有）

3. **注释规范**
   - 所有公开的函数、结构体都需要添加注释
   - 注释格式: `// FunctionName 函数说明`

## 配置说明

### 服务器配置
- 监听地址: `:8000`
- 日志路径: `temp/logs`

### JWT 配置
- 签名密钥: 在 `config.yaml` 中配置
- 过期时间: 24 小时

### 数据库配置
- 字符集: utf8mb4
- 连接池: 最大 100 个连接

## 部署

### 编译

```bash
go build -o gf-admin main.go
```

### 运行

```bash
./gf-admin
```

## 许可证

MIT License
