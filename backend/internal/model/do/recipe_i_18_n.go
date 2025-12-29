// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeI18N is the golang structure of table recipe_i18n for DAO operations like Where/Data.
type RecipeI18N struct {
	g.Meta      `orm:"table:recipe_i18n, do:true"`
	Id          interface{} // ID
	RecipeId    interface{} // 食谱ID
	Locale      interface{} // 语言代码: zh-CN, en-US, de-DE
	Name        interface{} // 食谱名称
	Subtitle    interface{} // 副标题
	Description interface{} // 简介
	Ingredients interface{} // 食材列表(JSON数组)
	Content     interface{} // 制作步骤(富文本HTML)
	Tags        interface{} // 标签（逗号分隔）
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
