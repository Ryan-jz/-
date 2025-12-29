// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeTag is the golang structure for table recipe_tag.
type RecipeTag struct {
	Id        uint        `json:"id"        orm:"id"         description:""`           //
	Name      string      `json:"name"      orm:"name"       description:"标签名称"`       // 标签名称
	Slug      string      `json:"slug"      orm:"slug"       description:"标签标识"`       // 标签标识
	SortOrder int         `json:"sortOrder" orm:"sort_order" description:"排序"`         // 排序
	Status    int         `json:"status"    orm:"status"     description:"状态 1启用 0禁用"` // 状态 1启用 0禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`           //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`           //
}
