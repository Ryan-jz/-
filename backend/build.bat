@echo off
chcp 65001 >nul
echo ========================================
echo Building bai-jin-yan-api Backend
echo ========================================

REM Set variables
set OUTPUT_DIR=dist
set APP_NAME=bai-jin-yan-api.exe

REM Clean old build directory
if exist %OUTPUT_DIR% (
    echo Cleaning old build directory...
    rmdir /s /q %OUTPUT_DIR%
)

REM Create build directory structure
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

REM Build Go program (64-bit, static compilation for Windows Server 2008)
echo Building Go program (64-bit, compatible with Windows Server 2008)...
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o %OUTPUT_DIR%\%APP_NAME% .

if %errorlevel% neq 0 (
    echo Build failed!
    pause
    exit /b 1
)

REM Copy configuration files
echo Copying configuration files...
xcopy /y manifest\config\*.* %OUTPUT_DIR%\manifest\config\
xcopy /y manifest\sql\*.* %OUTPUT_DIR%\manifest\sql\

REM Copy public directory static files (if any)
if exist public\*.* (
    xcopy /y /s public\*.* %OUTPUT_DIR%\public\
)

REM Create README
echo Creating README...
echo bai-jin-yan-api Backend Program > %OUTPUT_DIR%\README.txt
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
echo Creating startup script...
echo @echo off > %OUTPUT_DIR%\start.bat
echo echo Starting bai-jin-yan-api Backend Service... >> %OUTPUT_DIR%\start.bat
echo %APP_NAME% >> %OUTPUT_DIR%\start.bat
echo pause >> %OUTPUT_DIR%\start.bat

echo.
echo ========================================
echo Build completed!
echo Output directory: %OUTPUT_DIR%
echo Executable file: %OUTPUT_DIR%\%APP_NAME%
echo ========================================
echo.
pause
