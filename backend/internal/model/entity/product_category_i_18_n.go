// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductCategoryI18N is the golang structure for table product_category_i18n.
type ProductCategoryI18N struct {
	Id          uint        `json:"id"          orm:"id"          description:"ID"`                        // ID
	CategoryId  uint        `json:"categoryId"  orm:"category_id" description:"分类ID"`                      // 分类ID
	Locale      string      `json:"locale"      orm:"locale"      description:"语言代码: zh-CN, en-US, de-DE"` // 语言代码: zh-CN, en-US, de-DE
	Name        string      `json:"name"        orm:"name"        description:"分类名称"`                      // 分类名称
	Description string      `json:"description" orm:"description" description:"分类描述"`                      // 分类描述
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`                      // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`                      // 更新时间
}
