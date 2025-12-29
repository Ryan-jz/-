// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductDao is the data access object for table product.
type ProductDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ProductColumns // columns contains all the column names of Table for convenient usage.
}

// ProductColumns defines and stores column names for table product.
type ProductColumns struct {
	Id          string // 产品ID
	CategoryId  string // 分类ID
	Name        string // 产品名称
	NameEn      string // 产品英文名称
	Subtitle    string // 副标题
	Description string // 产品描述
	Image       string // 主图片
	Images      string // 产品图片集(JSON数组)
	Price       string // 价格
	Stock       string // 库存
	Weight      string // 重量/规格
	Ingredients string // 成分说明
	Nutrition   string // 营养信息(JSON)
	Usage       string // 使用方法
	Features    string // 产品特点(JSON数组)
	SortOrder   string // 排序
	Status      string // 状态：1上架 0下架
	ViewCount   string // 浏览次数
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// productColumns holds the columns for table product.
var productColumns = ProductColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Name:        "name",
	NameEn:      "name_en",
	Subtitle:    "subtitle",
	Description: "description",
	Image:       "image",
	Images:      "images",
	Price:       "price",
	Stock:       "stock",
	Weight:      "weight",
	Ingredients: "ingredients",
	Nutrition:   "nutrition",
	Usage:       "usage",
	Features:    "features",
	SortOrder:   "sort_order",
	Status:      "status",
	ViewCount:   "view_count",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewProductDao creates and returns a new DAO object for table data access.
func NewProductDao() *ProductDao {
	return &ProductDao{
		group:   "default",
		table:   "product",
		columns: productColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductDao) Columns() ProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
