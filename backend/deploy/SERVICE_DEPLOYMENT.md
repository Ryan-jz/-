# Windows 服务部署指南

## 前置要求

1. **安装 NSSM**
   - 下载地址: https://nssm.cc/download
   - 下载后解压，将 `nssm.exe` 添加到系统 PATH 环境变量
   - 或者将 `nssm.exe` 复制到项目目录

2. **配置数据库**
   - 确保 MySQL 8.0 已安装并运行
   - 导入数据库: `mysql -u root -p < manifest/sql/init.sql`
   - 修改 `manifest/config/config.yaml` 中的数据库连接信息

3. **配置 Redis**（可选）
   - 如果使用 Redis，确保 Redis 服务已启动
   - 修改配置文件中的 Redis 连接信息

## 快速部署

### 方法 1: 使用自动安装脚本（推荐）

1. **以管理员身份运行 PowerShell 或 CMD**

2. **执行安装脚本**
   ```cmd
   install-service.bat
   ```

3. **验证服务状态**
   ```cmd
   nssm status GFAdminService
   ```

### 方法 2: 手动安装

1. **以管理员身份运行 CMD**

2. **安装服务**
   ```cmd
   nssm install BaiJinYanAPI "C:\path\to\bai-jin-yan-api.exe"
   ```

3. **设置工作目录**
   ```cmd
   nssm set BaiJinYanAPI AppDirectory "C:\path\to\backend"
   ```

4. **设置日志输出**
   ```cmd
   nssm set BaiJinYanAPI AppStdout "C:\path\to\backend\temp\logs\service-stdout.log"
   nssm set BaiJinYanAPI AppStderr "C:\path\to\backend\temp\logs\service-stderr.log"
   ```

5. **设置自动重启**
   ```cmd
   nssm set BaiJinYanAPI AppExit Default Restart
   nssm set BaiJinYanAPI AppRestartDelay 5000
   ```

6. **启动服务**
   ```cmd
   nssm start BaiJinYanAPI
   ```

## 常见问题排查

### 问题 1: 服务启动后立即停止或显示 "paused"

**原因：**
- 程序启动失败
- 配置文件错误
- 数据库连接失败
- 缺少必要的目录或文件

**解决方法：**

1. **查看日志文件**
   ```
   temp/logs/service-stderr.log
   temp/logs/server/error.log
   ```

2. **手动运行程序测试**
   ```cmd
   cd C:\path\to\backend
   bai-jin-yan-api.exe
   ```
   观察是否有错误输出

3. **检查配置文件**
   - 确认 `manifest/config/config.yaml` 配置正确
   - 数据库连接信息是否正确
   - 端口是否被占用

4. **检查目录权限**
   - 确保服务账户有读写权限
   - 创建必要的目录：`temp/logs`、`resource/public`

### 问题 2: 数据库连接失败

**解决方法：**
1. 检查 MySQL 服务是否运行
2. 验证数据库连接信息
3. 确认数据库已创建并导入初始数据
4. 检查防火墙设置

### 问题 3: 端口被占用

**解决方法：**
1. 查看端口占用情况
   ```cmd
   netstat -ano | findstr :8000
   ```

2. 修改配置文件中的端口号
   ```yaml
   server:
     address: ":8001"  # 改为其他端口
   ```

### 问题 4: 服务无法启动

**解决方法：**
1. 检查是否以管理员身份运行
2. 确认 NSSM 已正确安装
3. 查看 Windows 事件查看器中的错误日志
4. 尝试重新安装服务

## 服务管理命令

### 启动服务
```cmd
nssm start BaiJinYanAPI
```

### 停止服务
```cmd
nssm stop BaiJinYanAPI
```

### 重启服务
```cmd
nssm restart BaiJinYanAPI
```

### 查看服务状态
```cmd
nssm status BaiJinYanAPI
```

### 编辑服务配置
```cmd
nssm edit BaiJinYanAPI
```

### 删除服务
```cmd
nssm remove BaiJinYanAPI confirm
```

或使用卸载脚本：
```cmd
uninstall-service.bat
```

## 服务配置说明

### 关键配置项

| 配置项 | 说明 | 示例值 |
|--------|------|--------|
| Application | 程序路径 | C:\app\bai-jin-yan-api.exe |
| Startup directory | 工作目录 | C:\app\backend |
| Arguments | 启动参数 | 无 |
| Service name | 服务名称 | BaiJinYanAPI |
| Display name | 显示名称 | BaiJinYan API Service |
| Description | 服务描述 | 白金燕 API 服务 |
| Startup type | 启动类型 | Automatic |

### 日志配置

| 配置项 | 路径 |
|--------|------|
| Output (stdout) | temp\logs\service-stdout.log |
| Error (stderr) | temp\logs\service-stderr.log |

### 重启策略

- 程序异常退出时自动重启
- 重启延迟：5 秒

## 生产环境建议

1. **使用专用服务账户**
   - 创建专门的 Windows 服务账户
   - 授予最小必要权限

2. **配置日志轮转**
   - 定期清理旧日志文件
   - 使用日志管理工具

3. **监控服务状态**
   - 配置服务监控告警
   - 定期检查日志

4. **备份配置文件**
   - 定期备份配置文件
   - 版本控制管理

5. **安全加固**
   - 修改默认密码和密钥
   - 配置防火墙规则
   - 启用 HTTPS

## 验证部署

1. **检查服务状态**
   ```cmd
   sc query BaiJinYanAPI
   ```

2. **测试 API 接口**
   ```cmd
   curl http://localhost:8000/login
   ```

3. **查看日志**
   ```cmd
   type temp\logs\service-stdout.log
   ```

## 卸载服务

使用卸载脚本：
```cmd
uninstall-service.bat
```

或手动卸载：
```cmd
nssm stop BaiJinYanAPI
nssm remove BaiJinYanAPI confirm
```

## 技术支持

如遇到问题，请提供以下信息：
- 错误日志内容
- 服务状态截图
- 配置文件内容
- Windows 版本信息
