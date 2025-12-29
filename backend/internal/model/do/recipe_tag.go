// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RecipeTag is the golang structure of table recipe_tag for DAO operations like Where/Data.
type RecipeTag struct {
	g.Meta    `orm:"table:recipe_tag, do:true"`
	Id        interface{} //
	Name      interface{} // 标签名称
	Slug      interface{} // 标签标识
	SortOrder interface{} // 排序
	Status    interface{} // 状态 1启用 0禁用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
