// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure for table banner.
type Banner struct {
	BannerId    int64       `json:"bannerId"    orm:"banner_id"   description:"轮播图ID"`                 // 轮播图ID
	Title       string      `json:"title"       orm:"title"       description:"标题"`                    // 标题
	Description string      `json:"description" orm:"description" description:"描述"`                    // 描述
	MediaType   int         `json:"mediaType"   orm:"media_type"  description:"媒体类型（1图片 2视频）"`         // 媒体类型（1图片 2视频）
	MediaUrl    string      `json:"mediaUrl"    orm:"media_url"   description:"媒体URL"`                 // 媒体URL
	ButtonText  string      `json:"buttonText"  orm:"button_text" description:"按钮文字"`                  // 按钮文字
	ButtonLink  string      `json:"buttonLink"  orm:"button_link" description:"按钮链接"`                  // 按钮链接
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"  description:"排序"`                    // 排序
	Status      int         `json:"status"      orm:"status"      description:"状态（0停用 1启用）"`           // 状态（0停用 1启用）
	Position    string      `json:"position"    orm:"position"    description:"位置（home首页 product产品页）"` // 位置（home首页 product产品页）
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"`                  // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新时间"`                  // 更新时间
}
