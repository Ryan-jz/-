# 部署检查清单

## ✅ 部署前检查

### 1. 环境准备
- [ ] MySQL 8.0 已安装并运行
- [ ] Redis 已安装并运行（可选）
- [ ] NSSM 已下载并添加到 PATH
- [ ] 以管理员身份运行命令行

### 2. 数据库准备
- [ ] 创建数据库：`gf_admin`
- [ ] 导入初始数据：`mysql -u root -p < manifest/sql/init.sql`
- [ ] 验证数据导入成功
- [ ] 测试数据库连接

### 3. 配置文件
- [ ] 修改 `manifest/config/config.yaml` 中的数据库连接信息
- [ ] 修改 JWT 签名密钥
- [ ] 检查端口配置（默认 8000）
- [ ] 配置 Redis 连接（如果使用）

### 4. 目录结构
- [ ] 确保存在 `temp/logs` 目录
- [ ] 确保存在 `resource/public` 目录
- [ ] 检查目录权限

### 5. 程序编译
- [ ] 执行编译命令：`$env:CGO_ENABLED="0"; go build -o bai-jin-yan-api.exe main.go`
- [ ] 验证 exe 文件生成成功
- [ ] 手动运行测试：`.\bai-jin-yan-api.exe`

## ✅ 服务安装

### 方法 A: 使用自动脚本（推荐）
```cmd
# 以管理员身份运行
install-service.bat
```

### 方法 B: 手动安装
```cmd
# 1. 安装服务
nssm install BaiJinYanAPI "完整路径\bai-jin-yan-api.exe"

# 2. 设置工作目录
nssm set BaiJinYanAPI AppDirectory "完整路径\backend"

# 3. 设置日志
nssm set BaiJinYanAPI AppStdout "完整路径\backend\temp\logs\service-stdout.log"
nssm set BaiJinYanAPI AppStderr "完整路径\backend\temp\logs\service-stderr.log"

# 4. 设置自动重启
nssm set BaiJinYanAPI AppExit Default Restart
nssm set BaiJinYanAPI AppRestartDelay 5000

# 5. 启动服务
nssm start BaiJinYanAPI
```

## ✅ 部署后验证

### 1. 检查服务状态
```cmd
# 查看服务状态
nssm status BaiJinYanAPI

# 或使用 Windows 服务管理
sc query BaiJinYanAPI
```

### 2. 检查日志
```cmd
# 查看标准输出日志
type temp\logs\service-stdout.log

# 查看错误日志
type temp\logs\service-stderr.log

# 查看应用日志
type temp\logs\server\*.log
```

### 3. 测试 API 接口
```cmd
# 测试登录接口
curl -X POST http://localhost:8000/login ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"admin\",\"password\":\"admin123\"}"
```

或在浏览器访问：
```
http://localhost:8000
```

### 4. 检查端口监听
```cmd
netstat -ano | findstr :8000
```

## ❌ 常见问题排查

### 问题 1: 服务显示 "PAUSED" 或立即停止

**排查步骤：**

1. 查看错误日志
```cmd
type temp\logs\service-stderr.log
```

2. 手动运行程序
```cmd
cd backend目录
.\bai-jin-yan-api.exe
```

3. 常见原因：
   - 数据库连接失败
   - 配置文件错误
   - 端口被占用
   - 缺少必要目录

### 问题 2: 数据库连接失败

**检查项：**
- [ ] MySQL 服务是否运行
- [ ] 数据库名称是否正确
- [ ] 用户名密码是否正确
- [ ] 数据库是否已创建
- [ ] 防火墙是否阻止连接

**测试连接：**
```cmd
mysql -h 127.0.0.1 -u root -p gf_admin
```

### 问题 3: 端口被占用

**查找占用进程：**
```cmd
netstat -ano | findstr :8000
```

**解决方法：**
- 关闭占用端口的程序
- 或修改配置文件使用其他端口

### 问题 4: 权限问题

**解决方法：**
- 以管理员身份运行
- 检查目录权限
- 配置服务账户权限

## 📋 服务管理命令速查

| 操作 | 命令 |
|------|------|
| 启动服务 | `nssm start BaiJinYanAPI` |
| 停止服务 | `nssm stop BaiJinYanAPI` |
| 重启服务 | `nssm restart BaiJinYanAPI` |
| 查看状态 | `nssm status BaiJinYanAPI` |
| 编辑配置 | `nssm edit BaiJinYanAPI` |
| 删除服务 | `nssm remove BaiJinYanAPI confirm` |

## 🔧 配置文件位置

| 文件 | 路径 |
|------|------|
| 主配置 | `manifest/config/config.yaml` |
| 数据库脚本 | `manifest/sql/init.sql` |
| 服务日志 | `temp/logs/service-*.log` |
| 应用日志 | `temp/logs/server/*.log` |

## 📞 获取帮助

如果遇到问题：
1. 查看 `SERVICE_DEPLOYMENT.md` 详细文档
2. 检查日志文件
3. 手动运行程序测试
4. 查看 Windows 事件查看器

## ✨ 部署成功标志

- [ ] 服务状态显示 "RUNNING"
- [ ] 可以访问 API 接口
- [ ] 日志文件正常输出
- [ ] 端口正常监听
- [ ] 登录功能正常

恭喜！部署成功！🎉
