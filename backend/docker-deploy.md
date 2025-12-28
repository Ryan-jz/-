# Docker 部署指南

## 前置要求

1. 安装 Docker Desktop（Windows）或 Docker Engine（Linux）
2. 安装 Docker Compose

## 快速开始

### 1. 使用 Docker Compose（推荐）

包含后端 + MySQL 数据库：

```bash
# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 停止并删除数据
docker-compose down -v
```

### 2. 仅构建后端镜像

如果使用外部数据库：

```bash
# 构建镜像
docker build -t gf-admin:latest .

# 运行容器
docker run -d \
  --name gf-admin-backend \
  -p 8000:8000 \
  -v $(pwd)/manifest/config:/app/manifest/config \
  -v $(pwd)/public/uploads:/app/public/uploads \
  -v $(pwd)/temp/logs:/app/temp/logs \
  gf-admin:latest
```

## 配置说明

### 修改数据库配置

编辑 `manifest/config/config.yaml`：

```yaml
database:
  default:
    link: "mysql:用户名:密码@tcp(mysql:3306)/数据库名"
```

如果使用 docker-compose，主机名使用 `mysql`（服务名）
如果使用外部数据库，使用实际的 IP 或域名

### 环境变量

在 `docker-compose.yml` 中可以配置：

- `TZ`: 时区设置
- `MYSQL_ROOT_PASSWORD`: MySQL root 密码
- `MYSQL_DATABASE`: 数据库名
- `MYSQL_USER`: 数据库用户
- `MYSQL_PASSWORD`: 数据库密码

## 访问应用

- 后端 API: http://localhost:8000
- MySQL: localhost:3306

## 数据持久化

以下目录会持久化到宿主机：

- `./manifest/config` - 配置文件
- `./public/uploads` - 上传的文件
- `./temp/logs` - 日志文件
- `mysql-data` - MySQL 数据（Docker volume）

## 常用命令

```bash
# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 查看日志
docker logs gf-admin-backend
docker logs gf-admin-mysql

# 进入容器
docker exec -it gf-admin-backend sh
docker exec -it gf-admin-mysql bash

# 重启容器
docker restart gf-admin-backend

# 删除容器
docker rm -f gf-admin-backend

# 删除镜像
docker rmi gf-admin:latest
```

## 生产环境建议

1. 修改默认密码
2. 使用外部 MySQL 数据库
3. 配置反向代理（Nginx）
4. 启用 HTTPS
5. 定期备份数据库和上传文件
6. 监控容器资源使用

## Windows Server 部署

1. 安装 Docker Desktop for Windows
2. 确保启用 WSL 2
3. 运行 `docker-compose up -d`

如果是 Windows Server 2008，建议升级到 2016+ 以支持 Docker。
