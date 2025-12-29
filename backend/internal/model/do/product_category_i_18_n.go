// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductCategoryI18N is the golang structure of table product_category_i18n for DAO operations like Where/Data.
type ProductCategoryI18N struct {
	g.Meta      `orm:"table:product_category_i18n, do:true"`
	Id          interface{} // ID
	CategoryId  interface{} // 分类ID
	Locale      interface{} // 语言代码: zh-CN, en-US, de-DE
	Name        interface{} // 分类名称
	Description interface{} // 分类描述
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
