// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Recipe is the golang structure for table recipe.
type Recipe struct {
	Id          uint        `json:"id"          orm:"id"           description:"食谱ID"`              // 食谱ID
	Name        string      `json:"name"        orm:"name"         description:"食谱名称"`              // 食谱名称
	NameEn      string      `json:"nameEn"      orm:"name_en"      description:"英文名称"`              // 英文名称
	Subtitle    string      `json:"subtitle"    orm:"subtitle"     description:"副标题"`               // 副标题
	Description string      `json:"description" orm:"description"  description:"简介"`                // 简介
	Image       string      `json:"image"       orm:"image"        description:"主图片"`               // 主图片
	Images      string      `json:"images"      orm:"images"       description:"详情图片（JSON数组）"`      // 详情图片（JSON数组）
	Ingredients string      `json:"ingredients" orm:"ingredients"  description:"食材列表（JSON数组）"`      // 食材列表（JSON数组）
	Content     string      `json:"content"     orm:"content"      description:"制作步骤（富文本HTML）"`     // 制作步骤（富文本HTML）
	CookingTime int         `json:"cookingTime" orm:"cooking_time" description:"烹饪时间（分钟）"`          // 烹饪时间（分钟）
	Difficulty  int         `json:"difficulty"  orm:"difficulty"   description:"难度等级（1简单 2中等 3困难）"` // 难度等级（1简单 2中等 3困难）
	Servings    int         `json:"servings"    orm:"servings"     description:"份量（几人份）"`           // 份量（几人份）
	ProductIds  string      `json:"productIds"  orm:"product_ids"  description:"关联产品ID（逗号分隔）"`      // 关联产品ID（逗号分隔）
	Tags        string      `json:"tags"        orm:"tags"         description:"标签（逗号分隔）"`          // 标签（逗号分隔）
	ViewCount   int         `json:"viewCount"   orm:"view_count"   description:"浏览次数"`              // 浏览次数
	LikeCount   int         `json:"likeCount"   orm:"like_count"   description:"点赞次数"`              // 点赞次数
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"   description:"排序"`                // 排序
	Status      int         `json:"status"      orm:"status"       description:"状态（0下架 1上架）"`       // 状态（0下架 1上架）
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`              // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:"更新时间"`              // 更新时间
}
