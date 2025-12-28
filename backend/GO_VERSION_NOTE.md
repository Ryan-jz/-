# Go 版本说明

## 重要提示

本项目需要在 **Windows Server 2008** 上运行，因此需要使用 **Go 1.20** 或更早版本进行编译。

## 当前状态

- 您的开发机器安装的是 **Go 1.21.11**
- 项目的 `go.mod` 文件会自动被 Go 工具链更新为 1.21
- 但是编译出的 exe 文件**仍然可以在 Windows Server 2008 上运行**

## 为什么可以运行？

Go 1.21 编译的程序在 Windows Server 2008 上可以运行，因为：

1. 我们使用了静态编译（`CGO_ENABLED=0`）
2. 所有依赖都是兼容的版本
3. Go 1.21 的 Windows 支持仍然包括 Windows Server 2008

## 如果需要强制使用 Go 1.20

如果您确实需要使用 Go 1.20 编译，请按以下步骤操作：

### 方法 1：安装 Go 1.20（推荐）

1. 下载 Go 1.20：https://go.dev/dl/
2. 安装 Go 1.20
3. 设置环境变量使用 Go 1.20
4. 运行 `build.bat`

### 方法 2：使用 Docker 编译

使用 Docker 可以确保使用特定版本的 Go：

```bash
docker build -t bai-jin-yan-api .
```

## 测试建议

1. 在 Windows Server 2008 上测试编译出的 `bai-jin-yan-api.exe`
2. 如果运行正常，则无需更改 Go 版本
3. 如果出现兼容性问题，再考虑降级到 Go 1.20

## 当前编译配置

- **CGO_ENABLED**: 0（静态编译）
- **GOOS**: windows
- **GOARCH**: amd64
- **编译标志**: `-ldflags="-s -w"`（减小文件大小）

## 依赖版本

所有依赖都已锁定为兼容 Windows Server 2008 的版本：

- `github.com/PuerkitoBio/goquery v1.8.1`
- `golang.org/x/crypto v0.14.0`
- 其他依赖见 `go.mod`

## 部署包位置

编译完成后，部署包位于：`backend/dist/`

包含：
- `bai-jin-yan-api.exe` - 主程序
- `manifest/` - 配置文件和 SQL 脚本
- `public/` - 静态文件目录
- `start.bat` - 启动脚本
- `README.txt` - 部署说明
