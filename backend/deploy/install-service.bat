@echo off
REM GF Admin 服务安装脚本
REM 使用 NSSM 将应用安装为 Windows 服务

echo ========================================
echo GF Admin 服务安装脚本
echo ========================================
echo.

REM 设置变量
set SERVICE_NAME=BaiJinYanAPI
set APP_PATH=%~dp0bai-jin-yan-api.exe
set APP_DIR=%~dp0

echo 服务名称: %SERVICE_NAME%
echo 程序路径: %APP_PATH%
echo 工作目录: %APP_DIR%
echo.

REM 检查是否以管理员身份运行
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo [错误] 请以管理员身份运行此脚本！
    pause
    exit /b 1
)

REM 检查 NSSM 是否存在
where nssm >nul 2>&1
if %errorLevel% neq 0 (
    echo [错误] 未找到 NSSM，请先安装 NSSM
    echo 下载地址: https://nssm.cc/download
    pause
    exit /b 1
)

REM 检查服务是否已存在
sc query %SERVICE_NAME% >nul 2>&1
if %errorLevel% equ 0 (
    echo [警告] 服务已存在，正在删除旧服务...
    nssm stop %SERVICE_NAME%
    nssm remove %SERVICE_NAME% confirm
    timeout /t 2 >nul
)

REM 安装服务
echo [1/5] 正在安装服务...
nssm install %SERVICE_NAME% "%APP_PATH%"

REM 设置工作目录
echo [2/5] 设置工作目录...
nssm set %SERVICE_NAME% AppDirectory "%APP_DIR%"

REM 设置启动类型为自动
echo [3/5] 设置启动类型...
nssm set %SERVICE_NAME% Start SERVICE_AUTO_START

REM 设置日志输出
echo [4/5] 配置日志输出...
nssm set %SERVICE_NAME% AppStdout "%APP_DIR%temp\logs\service-stdout.log"
nssm set %SERVICE_NAME% AppStderr "%APP_DIR%temp\logs\service-stderr.log"

REM 设置服务描述
nssm set %SERVICE_NAME% Description "白金盐 API 服务"
nssm set %SERVICE_NAME% DisplayName "BaiJinYan API Service"

REM 设置服务失败后自动重启
nssm set %SERVICE_NAME% AppExit Default Restart
nssm set %SERVICE_NAME% AppRestartDelay 5000

REM 启动服务
echo [5/5] 启动服务...
nssm start %SERVICE_NAME%

echo.
echo ========================================
echo 服务安装完成！
echo ========================================
echo.
echo 服务名称: %SERVICE_NAME%
echo 服务状态: 
sc query %SERVICE_NAME% | findstr "STATE"
echo.
echo 常用命令:
echo   启动服务: nssm start %SERVICE_NAME%
echo   停止服务: nssm stop %SERVICE_NAME%
echo   重启服务: nssm restart %SERVICE_NAME%
echo   查看状态: nssm status %SERVICE_NAME%
echo   删除服务: nssm remove %SERVICE_NAME% confirm
echo.
pause
