@echo off
chcp 65001 >nul
echo ========================================
echo 白金盐项目 - 生产环境打包脚本
echo ========================================
echo.

REM 检查 pnpm 是否安装
pnpm --version >nul 2>&1
if %errorlevel% neq 0 (
    echo [错误] 未检测到 pnpm！
    echo 请先安装 pnpm: npm install -g pnpm
    pause
    exit /b 1
)

echo [信息] 开始打包...
echo.

REM 打包后台管理系统
echo ========================================
echo [1/3] 打包后台管理系统 (frontend-admin)
echo ========================================
cd frontend-admin
if %errorlevel% neq 0 (
    echo [错误] frontend-admin 目录不存在！
    cd ..
    pause
    exit /b 1
)

echo [1.1] 安装依赖...
call pnpm install
if %errorlevel% neq 0 (
    echo [错误] 依赖安装失败！
    cd ..
    pause
    exit /b 1
)

echo [1.2] 开始构建...
call pnpm build
if %errorlevel% neq 0 (
    echo [错误] 构建失败！
    cd ..
    pause
    exit /b 1
)

echo [✓] 后台管理系统打包完成！
echo 输出目录: frontend-admin\dist
cd ..
echo.

REM 打包前台网站
echo ========================================
echo [2/3] 打包前台网站 (frontend-web)
echo ========================================
cd frontend-web
if %errorlevel% neq 0 (
    echo [错误] frontend-web 目录不存在！
    cd ..
    pause
    exit /b 1
)

echo [2.1] 安装依赖...
call pnpm install
if %errorlevel% neq 0 (
    echo [错误] 依赖安装失败！
    cd ..
    pause
    exit /b 1
)

echo [2.2] 开始构建...
call pnpm build
if %errorlevel% neq 0 (
    echo [错误] 构建失败！
    cd ..
    pause
    exit /b 1
)

echo [✓] 前台网站打包完成！
echo 输出目录: frontend-web\dist
cd ..
echo.

REM 创建部署包
echo ========================================
echo [3/3] 创建部署包
echo ========================================

REM 创建部署目录
if exist deploy (
    echo [信息] 清理旧的部署目录...
    rmdir /s /q deploy
)
mkdir deploy
mkdir deploy\admin
mkdir deploy\web
mkdir deploy\backend

echo [3.1] 复制前端文件...
xcopy /s /e /y frontend-admin\dist\*.* deploy\admin\
xcopy /s /e /y frontend-web\dist\*.* deploy\web\

echo [3.2] 复制后端文件...
xcopy /s /e /y backend\dist\*.* deploy\backend\

echo [3.3] 创建部署说明...
echo 白金盐项目 - 部署包 > deploy\README.txt
echo. >> deploy\README.txt
echo 生成时间: %date% %time% >> deploy\README.txt
echo. >> deploy\README.txt
echo 目录说明: >> deploy\README.txt
echo - admin\    后台管理系统（前端） >> deploy\README.txt
echo - web\      前台网站（前端） >> deploy\README.txt
echo - backend\  后端 API 服务 >> deploy\README.txt
echo. >> deploy\README.txt
echo 部署步骤: >> deploy\README.txt
echo 1. 将 backend 目录部署到服务器，参考 backend\Windows_Server_2008_R2_部署说明.txt >> deploy\README.txt
echo 2. 将 admin 和 web 目录部署到 Web 服务器（Nginx/IIS） >> deploy\README.txt
echo 3. 详细说明请查看：前端部署说明.md >> deploy\README.txt
echo. >> deploy\README.txt
echo 访问地址: >> deploy\README.txt
echo - 前台网站: http://8.133.175.112 >> deploy\README.txt
echo - 后台管理: http://8.133.175.112:8080 >> deploy\README.txt
echo - API 接口: http://8.133.175.112:8000 >> deploy\README.txt

REM 复制部署说明
copy 前端部署说明.md deploy\
copy backend\Windows_Server_2008_R2_部署说明.txt deploy\backend\

echo.
echo ========================================
echo ✅ 打包完成！
echo ========================================
echo.
echo 部署包位置: deploy\
echo.
echo 目录结构:
echo   deploy\
echo   ├── admin\           后台管理系统
echo   ├── web\             前台网站
echo   ├── backend\         后端服务
echo   ├── README.txt       部署说明
echo   └── 前端部署说明.md   详细文档
echo.
echo 文件大小:
dir deploy /s | find "个文件"
echo.
echo 下一步:
echo 1. 将 deploy 目录打包（zip/tar.gz）
echo 2. 上传到服务器 8.133.175.112
echo 3. 按照说明文档部署
echo.
pause
