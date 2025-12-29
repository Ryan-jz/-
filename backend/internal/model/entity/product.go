// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	Id          uint        `json:"id"          orm:"id"          description:"产品ID"`          // 产品ID
	CategoryId  uint        `json:"categoryId"  orm:"category_id" description:"分类ID"`          // 分类ID
	Name        string      `json:"name"        orm:"name"        description:"产品名称"`          // 产品名称
	NameEn      string      `json:"nameEn"      orm:"name_en"     description:"产品英文名称"`        // 产品英文名称
	Subtitle    string      `json:"subtitle"    orm:"subtitle"    description:"副标题"`           // 副标题
	Description string      `json:"description" orm:"description" description:"产品描述"`          // 产品描述
	Image       string      `json:"image"       orm:"image"       description:"主图片"`           // 主图片
	Images      string      `json:"images"      orm:"images"      description:"产品图片集(JSON数组)"` // 产品图片集(JSON数组)
	Price       float64     `json:"price"       orm:"price"       description:"价格"`            // 价格
	Stock       int         `json:"stock"       orm:"stock"       description:"库存"`            // 库存
	Weight      string      `json:"weight"      orm:"weight"      description:"重量/规格"`         // 重量/规格
	Ingredients string      `json:"ingredients" orm:"ingredients" description:"成分说明"`          // 成分说明
	Nutrition   string      `json:"nutrition"   orm:"nutrition"   description:"营养信息(JSON)"`    // 营养信息(JSON)
	Usage       string      `json:"usage"       orm:"usage"       description:"使用方法"`          // 使用方法
	Features    string      `json:"features"    orm:"features"    description:"产品特点(JSON数组)"`  // 产品特点(JSON数组)
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"  description:"排序"`            // 排序
	Status      int         `json:"status"      orm:"status"      description:"状态：1上架 0下架"`    // 状态：1上架 0下架
	ViewCount   int         `json:"viewCount"   orm:"view_count"  description:"浏览次数"`          // 浏览次数
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`          // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`          // 更新时间
}
