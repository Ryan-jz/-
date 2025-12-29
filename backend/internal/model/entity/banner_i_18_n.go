// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BannerI18N is the golang structure for table banner_i18n.
type BannerI18N struct {
	Id          uint        `json:"id"          orm:"id"          description:"ID"`                        // ID
	BannerId    uint        `json:"bannerId"    orm:"banner_id"   description:"Banner ID"`                 // Banner ID
	Locale      string      `json:"locale"      orm:"locale"      description:"语言代码: zh-CN, en-US, de-DE"` // 语言代码: zh-CN, en-US, de-DE
	Title       string      `json:"title"       orm:"title"       description:"标题"`                        // 标题
	Description string      `json:"description" orm:"description" description:"描述"`                        // 描述
	ButtonText  string      `json:"buttonText"  orm:"button_text" description:"按钮文字"`                      // 按钮文字
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`                      // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`                      // 更新时间
}
