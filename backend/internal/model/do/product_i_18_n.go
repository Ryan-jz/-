// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductI18N is the golang structure of table product_i18n for DAO operations like Where/Data.
type ProductI18N struct {
	g.Meta      `orm:"table:product_i18n, do:true"`
	Id          interface{} // ID
	ProductId   interface{} // 产品ID
	Locale      interface{} // 语言代码: zh-CN, en-US, de-DE
	Name        interface{} // 产品名称
	Subtitle    interface{} // 副标题
	Description interface{} // 产品描述
	Ingredients interface{} // 成分说明
	Usage       interface{} // 使用方法
	Features    interface{} // 产品特点(JSON数组)
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
