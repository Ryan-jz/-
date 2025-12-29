# GoFrame DAO 自动生成指南

## 为什么使用 DAO 生成？

当前项目手写 SQL 查询遇到的问题：
1. ❌ ORM 链式查询语法容易出错（LeftJoin 参数问题）
2. ❌ 类型不安全，容易拼写错误
3. ❌ 维护成本高，数据库变更需要手动修改代码

使用 GoFrame DAO 生成的优势：
1. ✅ 自动生成类型安全的 Model 和 DAO
2. ✅ 标准化的查询 API，避免 SQL 语法错误
3. ✅ 数据库变更后重新生成即可
4. ✅ 支持完整的 CRUD 操作和复杂查询

## 安装 GoFrame CLI

### Windows
```bash
# 下载并安装
go install github.com/gogf/gf/cmd/gf/v2@latest

# 验证安装
gf version
```

### 配置环境变量
确保 `%GOPATH%\bin` 在系统 PATH 中。

## 当前配置

配置文件：`backend/hack/config.yaml`

```yaml
gfcli:
  gen:
    dao:
    - link: "mysql:jz:baijinyan!001@tcp(rm-cn-smw4l1pxp0001f8o.rwlb.rds.aliyuncs.com:3306)/bai_jin_yan"
      tables: ""           # 空表示生成所有表
      removePrefix: ""     # 移除表名前缀
      descriptionTag: true # 生成字段描述
      noModelComment: false
      path: "./internal"   # 生成路径
      group: "default"     # 数据库分组
      jsonCase: "CamelLower" # JSON 字段命名：驼峰小写
```

## 生成 DAO

```bash
cd backend
gf gen dao
```

生成的文件结构：
```
backend/internal/
├── dao/              # DAO 访问对象（自动生成，不要手动修改）
│   ├── banner.go
│   ├── banner_i18n.go
│   ├── product.go
│   ├── product_i18n.go
│   └── ...
├── model/            # 数据模型
│   ├── do/          # Data Object - 用于数据操作
│   │   ├── banner.go
│   │   └── ...
│   └── entity/      # Entity - 数据库表实体
│       ├── banner.go
│       └── ...
└── service/         # 业务逻辑（手动编写）
```

## 使用示例

### 1. 简单查询

**之前（手写 SQL）：**
```go
var list []BannerEntity
err := g.DB().Model("banner").Where("status", 1).Scan(&list)
```

**使用 DAO：**
```go
import "gf-admin/internal/dao"
import "gf-admin/internal/model/entity"

var list []*entity.Banner
err := dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Status, 1).Scan(&list)
```

### 2. 带国际化的查询

**之前（容易出错的 LeftJoin）：**
```go
model := g.DB().Model("banner b").
    LeftJoin("banner_i18n bi", "b.banner_id = bi.banner_id").
    Where("bi.locale = ?", locale)
```

**使用 DAO（推荐使用原生 SQL）：**
```go
sql := `
    SELECT 
        b.*,
        bi.title,
        bi.description
    FROM banner b
    LEFT JOIN banner_i18n bi ON b.banner_id = bi.banner_id AND bi.locale = ?
    WHERE b.status = ?
`
var list []*entity.Banner
err := g.DB().Ctx(ctx).Raw(sql, locale, status).Scan(&list)
```

### 3. 创建记录

**使用 DAO：**
```go
import "gf-admin/internal/model/do"

_, err := dao.Banner.Ctx(ctx).Data(do.Banner{
    MediaType: 1,
    MediaUrl:  "https://example.com/image.jpg",
    Status:    1,
}).Insert()
```

### 4. 更新记录

**使用 DAO：**
```go
_, err := dao.Banner.Ctx(ctx).
    Where(dao.Banner.Columns().BannerId, bannerId).
    Data(do.Banner{
        Status: 0,
    }).
    Update()
```

### 5. 事务操作

**使用 DAO：**
```go
err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
    // 插入主表
    result, err := dao.Banner.Ctx(ctx).TX(tx).Data(bannerData).Insert()
    if err != nil {
        return err
    }
    
    bannerId, _ := result.LastInsertId()
    
    // 插入国际化数据
    _, err = dao.BannerI18n.Ctx(ctx).TX(tx).Data(i18nData).Insert()
    return err
})
```

## 重构建议

### 当前临时方案（已实现）
使用原生 SQL（`g.DB().Raw()`）避免 ORM 语法问题：
- ✅ 立即可用，无需安装工具
- ✅ SQL 语法清晰，易于调试
- ⚠️ 需要手动维护 SQL 字符串

### 长期方案（推荐）
1. 安装 GoFrame CLI 工具
2. 运行 `gf gen dao` 生成 DAO
3. 重构 service 层使用生成的 DAO
4. 数据库变更时重新生成

## 国际化查询最佳实践

对于带国际化的复杂查询，推荐使用原生 SQL：

```go
func (s *bannerImpl) GetList(ctx context.Context, locale string, status int) ([]*entity.Banner, error) {
    sql := `
        SELECT 
            b.banner_id,
            b.media_type,
            b.media_url,
            b.status,
            COALESCE(bi.title, '') as title,
            COALESCE(bi.description, '') as description
        FROM banner b
        LEFT JOIN banner_i18n bi ON b.banner_id = bi.banner_id AND bi.locale = ?
        WHERE b.status = ?
        ORDER BY b.sort_order ASC
    `
    
    var list []*entity.Banner
    err := g.DB().Ctx(ctx).Raw(sql, locale, status).Scan(&list)
    return list, err
}
```

优点：
- SQL 语法标准，不依赖 ORM 特性
- 易于调试和优化
- 性能可控
- 支持复杂的 JOIN 和子查询

## 下一步

1. ✅ **当前已完成**：使用原生 SQL 重写了 Banner 和 Product 的查询
2. ⏳ **编译测试**：`cd backend && .\build.bat`
3. ⏳ **安装 CLI**：`go install github.com/gogf/gf/cmd/gf/v2@latest`
4. ⏳ **生成 DAO**：`gf gen dao`
5. ⏳ **逐步重构**：将 service 层迁移到使用 DAO

## 参考文档

- [GoFrame DAO 生成](https://goframe.org/pages/viewpage.action?pageId=1114119)
- [GoFrame ORM 查询](https://goframe.org/pages/viewpage.action?pageId=1114245)
- [GoFrame 事务处理](https://goframe.org/pages/viewpage.action?pageId=1114252)
