// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeI18N is the golang structure for table recipe_i18n.
type RecipeI18N struct {
	Id          uint        `json:"id"          orm:"id"          description:"ID"`                        // ID
	RecipeId    uint        `json:"recipeId"    orm:"recipe_id"   description:"食谱ID"`                      // 食谱ID
	Locale      string      `json:"locale"      orm:"locale"      description:"语言代码: zh-CN, en-US, de-DE"` // 语言代码: zh-CN, en-US, de-DE
	Name        string      `json:"name"        orm:"name"        description:"食谱名称"`                      // 食谱名称
	Subtitle    string      `json:"subtitle"    orm:"subtitle"    description:"副标题"`                       // 副标题
	Description string      `json:"description" orm:"description" description:"简介"`                        // 简介
	Ingredients string      `json:"ingredients" orm:"ingredients" description:"食材列表(JSON数组)"`              // 食材列表(JSON数组)
	Content     string      `json:"content"     orm:"content"     description:"制作步骤(富文本HTML)"`             // 制作步骤(富文本HTML)
	Tags        string      `json:"tags"        orm:"tags"        description:"标签（逗号分隔）"`                  // 标签（逗号分隔）
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`                      // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`                      // 更新时间
}
