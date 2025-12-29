// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeTagRelation is the golang structure for table recipe_tag_relation.
type RecipeTagRelation struct {
	Id        uint        `json:"id"        orm:"id"         description:""`     //
	RecipeId  uint        `json:"recipeId"  orm:"recipe_id"  description:"食谱ID"` // 食谱ID
	TagId     uint        `json:"tagId"     orm:"tag_id"     description:"标签ID"` // 标签ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`     //
}
