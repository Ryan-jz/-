// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductI18N is the golang structure for table product_i18n.
type ProductI18N struct {
	Id          uint        `json:"id"          orm:"id"          description:"ID"`                        // ID
	ProductId   uint        `json:"productId"   orm:"product_id"  description:"产品ID"`                      // 产品ID
	Locale      string      `json:"locale"      orm:"locale"      description:"语言代码: zh-CN, en-US, de-DE"` // 语言代码: zh-CN, en-US, de-DE
	Name        string      `json:"name"        orm:"name"        description:"产品名称"`                      // 产品名称
	Subtitle    string      `json:"subtitle"    orm:"subtitle"    description:"副标题"`                       // 副标题
	Description string      `json:"description" orm:"description" description:"产品描述"`                      // 产品描述
	Ingredients string      `json:"ingredients" orm:"ingredients" description:"成分说明"`                      // 成分说明
	Usage       string      `json:"usage"       orm:"usage"       description:"使用方法"`                      // 使用方法
	Features    string      `json:"features"    orm:"features"    description:"产品特点(JSON数组)"`              // 产品特点(JSON数组)
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`                      // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`                      // 更新时间
}
