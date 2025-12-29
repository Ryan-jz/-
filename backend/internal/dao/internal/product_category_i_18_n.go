// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductCategoryI18NDao is the data access object for table product_category_i18n.
type ProductCategoryI18NDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns ProductCategoryI18NColumns // columns contains all the column names of Table for convenient usage.
}

// ProductCategoryI18NColumns defines and stores column names for table product_category_i18n.
type ProductCategoryI18NColumns struct {
	Id          string // ID
	CategoryId  string // 分类ID
	Locale      string // 语言代码: zh-CN, en-US, de-DE
	Name        string // 分类名称
	Description string // 分类描述
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// productCategoryI18NColumns holds the columns for table product_category_i18n.
var productCategoryI18NColumns = ProductCategoryI18NColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Locale:      "locale",
	Name:        "name",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewProductCategoryI18NDao creates and returns a new DAO object for table data access.
func NewProductCategoryI18NDao() *ProductCategoryI18NDao {
	return &ProductCategoryI18NDao{
		group:   "default",
		table:   "product_category_i18n",
		columns: productCategoryI18NColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductCategoryI18NDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductCategoryI18NDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductCategoryI18NDao) Columns() ProductCategoryI18NColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductCategoryI18NDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductCategoryI18NDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductCategoryI18NDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
