// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeTagRelation is the golang structure of table recipe_tag_relation for DAO operations like Where/Data.
type RecipeTagRelation struct {
	g.Meta    `orm:"table:recipe_tag_relation, do:true"`
	Id        interface{} //
	RecipeId  interface{} // 食谱ID
	TagId     interface{} // 标签ID
	CreatedAt *gtime.Time //
}
