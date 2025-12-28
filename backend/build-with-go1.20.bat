@echo off
chcp 65001 >nul
echo ========================================
echo Building with Go 1.20 for Windows Server 2008 R2
echo ========================================
echo.

REM Check if Docker is available
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo [错误] 未检测到 Docker！
    echo.
    echo 请选择以下方案之一：
    echo 1. 安装 Docker Desktop
    echo 2. 安装 Go 1.20（参考：安装Go1.20说明.txt）
    echo.
    pause
    exit /b 1
)

echo [信息] 检测到 Docker，使用 Go 1.20 镜像编译...
echo.

REM Set variables
set OUTPUT_DIR=dist
set APP_NAME=bai-jin-yan-api.exe

REM Clean old build directory
if exist %OUTPUT_DIR% (
    echo [1/5] 清理旧的编译目录...
    rmdir /s /q %OUTPUT_DIR%
)

REM Create build directory structure
echo [2/5] 创建目录结构...
mkdir %OUTPUT_DIR%
mkdir %OUTPUT_DIR%\manifest
mkdir %OUTPUT_DIR%\manifest\config
mkdir %OUTPUT_DIR%\manifest\sql
mkdir %OUTPUT_DIR%\public
mkdir %OUTPUT_DIR%\public\uploads
mkdir %OUTPUT_DIR%\public\uploads\images
mkdir %OUTPUT_DIR%\public\uploads\videos
mkdir %OUTPUT_DIR%\temp
mkdir %OUTPUT_DIR%\temp\logs

REM Build with Docker using Go 1.20
echo [3/5] 使用 Go 1.20 编译程序（兼容 Windows Server 2008 R2）...
docker run --rm -v "%cd%":/app -w /app golang:1.20-alpine sh -c "apk add --no-cache git && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-s -w' -o %OUTPUT_DIR%/%APP_NAME% ."

if %errorlevel% neq 0 (
    echo [错误] 编译失败！
    pause
    exit /b 1
)

REM Copy configuration files
echo [4/5] 复制配置文件...
xcopy /y manifest\config\*.* %OUTPUT_DIR%\manifest\config\
xcopy /y manifest\sql\*.* %OUTPUT_DIR%\manifest\sql\

REM Copy public directory static files (if any)
if exist public\*.* (
    xcopy /y /s public\*.* %OUTPUT_DIR%\public\
)

REM Create README
echo [5/5] 创建说明文件...
echo bai-jin-yan-api Backend Program > %OUTPUT_DIR%\README.txt
echo. >> %OUTPUT_DIR%\README.txt
echo ⚠️ 本程序使用 Go 1.20 编译，兼容 Windows Server 2008 R2 >> %OUTPUT_DIR%\README.txt
echo. >> %OUTPUT_DIR%\README.txt
echo Usage: >> %OUTPUT_DIR%\README.txt
echo 1. Make sure MySQL database is installed >> %OUTPUT_DIR%\README.txt
echo 2. Import SQL files from manifest/sql directory >> %OUTPUT_DIR%\README.txt
echo 3. Modify database configuration in manifest/config/config.yaml >> %OUTPUT_DIR%\README.txt
echo 4. Double click %APP_NAME% to run >> %OUTPUT_DIR%\README.txt
echo 5. Default access address: http://localhost:8000 >> %OUTPUT_DIR%\README.txt
echo. >> %OUTPUT_DIR%\README.txt
echo Directory Structure: >> %OUTPUT_DIR%\README.txt
echo - manifest/config/  Configuration files >> %OUTPUT_DIR%\README.txt
echo - manifest/sql/     Database initialization scripts >> %OUTPUT_DIR%\README.txt
echo - public/uploads/   File upload directory >> %OUTPUT_DIR%\README.txt
echo - temp/logs/        Log files directory >> %OUTPUT_DIR%\README.txt
echo. >> %OUTPUT_DIR%\README.txt
echo Notes: >> %OUTPUT_DIR%\README.txt
echo - Configure database connection before first run >> %OUTPUT_DIR%\README.txt
echo - Make sure port 8000 is not occupied >> %OUTPUT_DIR%\README.txt
echo - Uploaded files will be saved in public/uploads directory >> %OUTPUT_DIR%\README.txt

REM Create startup script
echo @echo off > %OUTPUT_DIR%\start.bat
echo echo Starting bai-jin-yan-api Backend Service... >> %OUTPUT_DIR%\start.bat
echo %APP_NAME% >> %OUTPUT_DIR%\start.bat
echo pause >> %OUTPUT_DIR%\start.bat

REM Verify Go version
echo.
echo ========================================
echo 验证编译版本...
echo ========================================
docker run --rm -v "%cd%":/app -w /app golang:1.20-alpine go version -m /app/%OUTPUT_DIR%/%APP_NAME% | findstr "go1.20"

echo.
echo ========================================
echo 编译完成！
echo 输出目录: %OUTPUT_DIR%
echo 可执行文件: %OUTPUT_DIR%\%APP_NAME%
echo ✅ 已使用 Go 1.20 编译，兼容 Windows Server 2008 R2
echo ========================================
echo.
pause
