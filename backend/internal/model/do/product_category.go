// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductCategory is the golang structure of table product_category for DAO operations like Where/Data.
type ProductCategory struct {
	g.Meta      `orm:"table:product_category, do:true"`
	Id          interface{} // 分类ID
	Name        interface{} // 分类名称
	NameEn      interface{} // 分类英文名称
	Slug        interface{} // 分类标识
	Description interface{} // 分类描述
	SortOrder   interface{} // 排序
	Status      interface{} // 状态：1启用 0禁用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Image       interface{} // 分类图片
}
