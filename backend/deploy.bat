@echo off
chcp 65001 >nul
echo ========================================
echo 部署到线上服务器
echo ========================================

REM 停止旧服务
echo 停止旧服务...
taskkill /F /IM bai-jin-yan-api.exe 2>nul

REM 等待进程完全停止
timeout /t 2 /nobreak >nul

REM 备份旧版本
if exist dist_backup (
    rmdir /s /q dist_backup
)
if exist dist (
    echo 备份旧版本...
    move dist dist_backup
)

REM 重新构建
call build.bat

REM 启动新服务
echo 启动新服务...
cd dist
start "" bai-jin-yan-api.exe

echo.
echo ========================================
echo 部署完成！
echo ========================================
pause
