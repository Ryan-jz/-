# Banner 国际化功能使用指南

## 概述

Banner 轮播图现已支持多语言国际化功能，支持中文（zh-CN）、英文（en-US）、德文（de-DE）三种语言。

## 部署步骤

### 1. 数据库迁移

执行以下命令创建国际化表并迁移现有数据：

```bash
# 方式1：使用 MySQL 命令行
mysql -u your_username -p your_database < backend/manifest/sql/banner_i18n_migration.sql

# 方式2：在 MySQL 客户端中执行
source backend/manifest/sql/banner_i18n_migration.sql;
```

### 2. 重新编译后端

```bash
cd backend
build.bat
```

### 3. 重启后端服务

停止当前运行的后端服务，然后重新启动：

```bash
cd backend
gf-admin.exe
```

### 4. 前端无需重新编译

前端代码已更新，刷新浏览器即可看到新功能。

## 功能说明

### 后台管理界面

1. **新增轮播图**
   - 点击"新增轮播图"按钮
   - 在弹出的对话框中，首先看到"多语言内容"部分
   - 切换标签页（中文/English/Deutsch）填写不同语言的内容
   - 填写媒体设置（上传图片/视频）
   - 填写其他设置（位置、排序、状态）
   - 点击"确定"保存

2. **编辑轮播图**
   - 点击列表中的"编辑"按钮
   - 系统会自动加载该轮播图的所有语言版本
   - 可以分别修改各语言的内容
   - 点击"确定"保存更新

3. **多语言字段**
   - **标题**：轮播图的主标题
   - **描述**：轮播图的详细描述
   - **按钮文字**：行动按钮上显示的文字

### 前端网站显示

前端网站会根据用户选择的语言自动显示对应语言版本的轮播图内容。

## API 接口说明

### 1. 获取轮播图列表（带语言参数）

```javascript
GET /api/v1/banner/list?locale=zh-CN

参数：
- position: 位置（home/product）
- status: 状态（0/1）
- page: 页码
- pageSize: 每页数量
- locale: 语言代码（zh-CN/en-US/de-DE）

返回：
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "bannerId": 1,
        "title": "标题（对应语言版本）",
        "description": "描述（对应语言版本）",
        "buttonText": "按钮文字（对应语言版本）",
        "mediaType": 1,
        "mediaUrl": "...",
        ...
      }
    ],
    "total": 10
  }
}
```

### 2. 获取轮播图国际化数据

```javascript
GET /api/v1/banner/:bannerId/i18n

返回：
{
  "code": 0,
  "message": "success",
  "data": {
    "zh-CN": {
      "title": "中文标题",
      "description": "中文描述",
      "buttonText": "中文按钮"
    },
    "en-US": {
      "title": "English Title",
      "description": "English Description",
      "buttonText": "English Button"
    },
    "de-DE": {
      "title": "Deutscher Titel",
      "description": "Deutsche Beschreibung",
      "buttonText": "Deutscher Button"
    }
  }
}
```

### 3. 创建轮播图（带国际化数据）

```javascript
POST /api/v1/banner/create

请求体：
{
  "mediaType": 1,
  "mediaUrl": "...",
  "buttonLink": "...",
  "position": "home",
  "sortOrder": 0,
  "status": 1,
  "i18n": {
    "zh-CN": {
      "title": "中文标题",
      "description": "中文描述",
      "buttonText": "中文按钮"
    },
    "en-US": {
      "title": "English Title",
      "description": "English Description",
      "buttonText": "English Button"
    },
    "de-DE": {
      "title": "Deutscher Titel",
      "description": "Deutsche Beschreibung",
      "buttonText": "Deutscher Button"
    }
  }
}
```

### 4. 更新轮播图（带国际化数据）

```javascript
PUT /api/v1/banner/update

请求体格式同创建接口，需额外包含 bannerId
```

## 数据库表结构

### banner 表（主表）
- 存储轮播图的基本信息（媒体文件、链接、排序等）
- 不再存储 title、description、button_text 字段

### banner_i18n 表（国际化表）
- 存储轮播图的多语言内容
- 每个轮播图对应 3 条记录（zh-CN、en-US、de-DE）
- 通过 banner_id 和 locale 联合唯一索引确保数据一致性

## 注意事项

1. **必填字段**：创建轮播图时，至少需要填写中文版本的标题
2. **数据迁移**：现有轮播图数据会自动迁移，每个轮播图会创建 3 个语言版本（初始内容相同）
3. **语言切换**：前端需要根据用户选择的语言传递 `locale` 参数
4. **默认语言**：如果不传递 `locale` 参数，默认返回中文（zh-CN）版本

## 后续扩展

如需添加更多语言支持：

1. 在 `banner_i18n` 表中插入新语言的数据
2. 在前端 `BannerI18nEditor.vue` 组件中添加新的标签页
3. 更新 `GetI18n` 方法中的 `locales` 数组

## 故障排查

### 问题1：编辑时看不到国际化内容

**解决方案**：
- 检查数据库中 `banner_i18n` 表是否有对应数据
- 检查浏览器控制台是否有 API 错误
- 确认后端路由 `/api/v1/banner/:bannerId/i18n` 是否正常

### 问题2：保存时提示"国际化内容不能为空"

**解决方案**：
- 确保至少填写了中文版本的标题
- 检查前端是否正确传递了 `i18n` 对象

### 问题3：前端显示的还是旧数据

**解决方案**：
- 清除浏览器缓存
- 检查 API 请求是否包含正确的 `locale` 参数
- 确认数据库中的国际化数据是否正确
