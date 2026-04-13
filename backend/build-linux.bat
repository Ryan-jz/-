@echo off
chcp 65001 >nul
setlocal

echo ========================================
echo Building bai-jin-yan-api for Linux
echo ========================================

set OUTPUT_DIR=dist-linux
set APP_NAME=bai-jin-yan-api-linux-amd64

if exist %OUTPUT_DIR% (
    echo Cleaning old build directory...
    rmdir /s /q %OUTPUT_DIR%
)

echo Creating build directory...
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

echo Building Linux binary...
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o %OUTPUT_DIR%\%APP_NAME% .

if %errorlevel% neq 0 (
    echo Build failed!
    pause
    exit /b 1
)

echo Copying configuration files...
xcopy /y manifest\config\*.* %OUTPUT_DIR%\manifest\config\
xcopy /y manifest\sql\*.* %OUTPUT_DIR%\manifest\sql\

if exist public\*.* (
    xcopy /y /s public\*.* %OUTPUT_DIR%\public\
)

echo Creating README...
echo bai-jin-yan-api Linux package > %OUTPUT_DIR%\README.txt
echo. >> %OUTPUT_DIR%\README.txt
echo Usage: >> %OUTPUT_DIR%\README.txt
echo 1. Upload the whole dist-linux directory to your Linux server >> %OUTPUT_DIR%\README.txt
echo 2. Make sure manifest/config/config.yaml has the correct database host >> %OUTPUT_DIR%\README.txt
echo 3. Run: chmod +x bai-jin-yan-api-linux-amd64 >> %OUTPUT_DIR%\README.txt
echo 4. Start: nohup ./bai-jin-yan-api-linux-amd64 ^> app.log 2^>^&1 ^& >> %OUTPUT_DIR%\README.txt
echo 5. Check: tail -n 100 app.log >> %OUTPUT_DIR%\README.txt

echo Creating start script...
echo #!/bin/sh > %OUTPUT_DIR%\start.sh
echo chmod +x ./bai-jin-yan-api-linux-amd64 >> %OUTPUT_DIR%\start.sh
echo nohup ./bai-jin-yan-api-linux-amd64 ^> app.log 2^>^&1 ^& >> %OUTPUT_DIR%\start.sh

echo.
echo ========================================
echo Build completed!
echo Output directory: %OUTPUT_DIR%
echo Binary file: %OUTPUT_DIR%\%APP_NAME%
echo ========================================
echo.
pause
