# 数据库初始化说明

## 执行顺序

请按以下顺序执行 SQL 文件：

1. `init.sql` - 创建数据库
2. `sys_user.sql` - 创建用户表
3. `product.sql` - 创建产品相关表
4. `recipe.sql` - 创建食谱表
5. `banner.sql` - 创建轮播图表

## 轮播图表说明

### 字段说明

- `banner_id`: 轮播图ID（主键）
- `title`: 标题
- `description`: 描述
- `media_type`: 媒体类型（1=图片，2=视频）
- `media_url`: 媒体文件URL
- `button_text`: 按钮文字
- `button_link`: 按钮链接
- `sort_order`: 排序（数字越小越靠前）
- `status`: 状态（0=停用，1=启用）
- `position`: 位置（home=首页，product=产品页）

### 示例数据

已包含3条示例轮播图数据，位置为首页。
