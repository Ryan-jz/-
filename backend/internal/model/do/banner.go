// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure of table banner for DAO operations like Where/Data.
type Banner struct {
	g.Meta      `orm:"table:banner, do:true"`
	BannerId    interface{} // 轮播图ID
	Title       interface{} // 标题
	Description interface{} // 描述
	MediaType   interface{} // 媒体类型（1图片 2视频）
	MediaUrl    interface{} // 媒体URL
	ButtonText  interface{} // 按钮文字
	ButtonLink  interface{} // 按钮链接
	SortOrder   interface{} // 排序
	Status      interface{} // 状态（0停用 1启用）
	Position    interface{} // 位置（home首页 product产品页）
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
