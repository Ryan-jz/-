// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BannerI18N is the golang structure of table banner_i18n for DAO operations like Where/Data.
type BannerI18N struct {
	g.Meta      `orm:"table:banner_i18n, do:true"`
	Id          interface{} // ID
	BannerId    interface{} // Banner ID
	Locale      interface{} // 语言代码: zh-CN, en-US, de-DE
	Title       interface{} // 标题
	Description interface{} // 描述
	ButtonText  interface{} // 按钮文字
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
