// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductCategoryDao is the data access object for table product_category.
type ProductCategoryDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ProductCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// ProductCategoryColumns defines and stores column names for table product_category.
type ProductCategoryColumns struct {
	Id          string // 分类ID
	Name        string // 分类名称
	NameEn      string // 分类英文名称
	Slug        string // 分类标识
	Description string // 分类描述
	SortOrder   string // 排序
	Status      string // 状态：1启用 0禁用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Image       string // 分类图片
}

// productCategoryColumns holds the columns for table product_category.
var productCategoryColumns = ProductCategoryColumns{
	Id:          "id",
	Name:        "name",
	NameEn:      "name_en",
	Slug:        "slug",
	Description: "description",
	SortOrder:   "sort_order",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Image:       "image",
}

// NewProductCategoryDao creates and returns a new DAO object for table data access.
func NewProductCategoryDao() *ProductCategoryDao {
	return &ProductCategoryDao{
		group:   "default",
		table:   "product_category",
		columns: productCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductCategoryDao) Columns() ProductCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
