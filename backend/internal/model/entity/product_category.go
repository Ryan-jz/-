// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductCategory is the golang structure for table product_category.
type ProductCategory struct {
	Id          uint        `json:"id"          orm:"id"          description:"分类ID"`       // 分类ID
	Name        string      `json:"name"        orm:"name"        description:"分类名称"`       // 分类名称
	NameEn      string      `json:"nameEn"      orm:"name_en"     description:"分类英文名称"`     // 分类英文名称
	Slug        string      `json:"slug"        orm:"slug"        description:"分类标识"`       // 分类标识
	Description string      `json:"description" orm:"description" description:"分类描述"`       // 分类描述
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"  description:"排序"`         // 排序
	Status      int         `json:"status"      orm:"status"      description:"状态：1启用 0禁用"` // 状态：1启用 0禁用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`       // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`       // 更新时间
	Image       string      `json:"image"       orm:"image"       description:"分类图片"`       // 分类图片
}
