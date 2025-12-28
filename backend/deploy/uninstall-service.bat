@echo off
REM GF Admin 服务卸载脚本

echo ========================================
echo GF Admin 服务卸载脚本
echo ========================================
echo.

set SERVICE_NAME=BaiJinYanAPI

REM 检查是否以管理员身份运行
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo [错误] 请以管理员身份运行此脚本！
    pause
    exit /b 1
)

REM 检查服务是否存在
sc query %SERVICE_NAME% >nul 2>&1
if %errorLevel% neq 0 (
    echo [警告] 服务不存在
    pause
    exit /b 0
)

echo 正在停止服务...
nssm stop %SERVICE_NAME%
timeout /t 2 >nul

echo 正在删除服务...
nssm remove %SERVICE_NAME% confirm

echo.
echo 服务已成功卸载！
echo.
pause
