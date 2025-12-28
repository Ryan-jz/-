# 部署说明

## 文件清单
- bai-jin-yan-api.exe - 主程�?
- manifest/sql/*.sql - 数据库脚�?
- uninstall-service.bat - 服务卸载脚本

## 部署步骤
1. 修改 manifest/config/config.yaml 中的数据库配�?
3. 以管理员身份运行 install-service.bat

## 测试运行
在安装服务前，建议先手动运行测试�?
bai-jin-yan-api.exe

详细说明请查�?SERVICE_DEPLOYMENT.md
