// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table product for DAO operations like Where/Data.
type Product struct {
	g.Meta        `orm:"table:product, do:true"`
	Id            interface{} // 产品ID
	CategoryId    interface{} // 分类ID
	Name          interface{} // 产品名称
	NameEn        interface{} // 产品英文名称
	Subtitle      interface{} // 副标题
	Description   interface{} // 产品描述
	Image         interface{} // 主图片
	Images        interface{} // 产品图片集(JSON数组)
	Price         interface{} // 价格
	Stock         interface{} // 库存
	Weight        interface{} // 重量/规格
	Ingredients   interface{} // 成分说明
	Nutrition     interface{} // 营养信息(JSON)
	Usage         interface{} // 使用方法
	Features      interface{} // 产品特点(JSON数组)
	OrganicCert   interface{} // 有机认证图标
	RecyclingInfo interface{} // 回收信息
	AllergenInfo  interface{} // 过敏原信息
	StorageInfo   interface{} // 储存信息
	Origin        interface{} // 产地
	SortOrder     interface{} // 排序
	Status        interface{} // 状态：1上架 0下架
	ViewCount     interface{} // 浏览次数
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
