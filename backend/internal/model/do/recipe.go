// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Recipe is the golang structure of table recipe for DAO operations like Where/Data.
type Recipe struct {
	g.Meta      `orm:"table:recipe, do:true"`
	Id          interface{} // 食谱ID
	Name        interface{} // 食谱名称
	NameEn      interface{} // 英文名称
	Subtitle    interface{} // 副标题
	Description interface{} // 简介
	Image       interface{} // 主图片
	Images      interface{} // 详情图片（JSON数组）
	Ingredients interface{} // 食材列表（JSON数组）
	Content     interface{} // 制作步骤（富文本HTML）
	CookingTime interface{} // 烹饪时间（分钟）
	Difficulty  interface{} // 难度等级（1简单 2中等 3困难）
	Servings    interface{} // 份量（几人份）
	ProductIds  interface{} // 关联产品ID（逗号分隔）
	Tags        interface{} // 标签（逗号分隔）
	ViewCount   interface{} // 浏览次数
	LikeCount   interface{} // 点赞次数
	SortOrder   interface{} // 排序
	Status      interface{} // 状态（0下架 1上架）
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
