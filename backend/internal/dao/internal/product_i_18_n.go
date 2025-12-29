// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductI18NDao is the data access object for table product_i18n.
type ProductI18NDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProductI18NColumns // columns contains all the column names of Table for convenient usage.
}

// ProductI18NColumns defines and stores column names for table product_i18n.
type ProductI18NColumns struct {
	Id          string // ID
	ProductId   string // 产品ID
	Locale      string // 语言代码: zh-CN, en-US, de-DE
	Name        string // 产品名称
	Subtitle    string // 副标题
	Description string // 产品描述
	Ingredients string // 成分说明
	Usage       string // 使用方法
	Features    string // 产品特点(JSON数组)
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// productI18NColumns holds the columns for table product_i18n.
var productI18NColumns = ProductI18NColumns{
	Id:          "id",
	ProductId:   "product_id",
	Locale:      "locale",
	Name:        "name",
	Subtitle:    "subtitle",
	Description: "description",
	Ingredients: "ingredients",
	Usage:       "usage",
	Features:    "features",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewProductI18NDao creates and returns a new DAO object for table data access.
func NewProductI18NDao() *ProductI18NDao {
	return &ProductI18NDao{
		group:   "default",
		table:   "product_i18n",
		columns: productI18NColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductI18NDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductI18NDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductI18NDao) Columns() ProductI18NColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductI18NDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductI18NDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductI18NDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
