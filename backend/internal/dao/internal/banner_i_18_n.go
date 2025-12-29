// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerI18NDao is the data access object for table banner_i18n.
type BannerI18NDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns BannerI18NColumns // columns contains all the column names of Table for convenient usage.
}

// BannerI18NColumns defines and stores column names for table banner_i18n.
type BannerI18NColumns struct {
	Id          string // ID
	BannerId    string // Banner ID
	Locale      string // 语言代码: zh-CN, en-US, de-DE
	Title       string // 标题
	Description string // 描述
	ButtonText  string // 按钮文字
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// bannerI18NColumns holds the columns for table banner_i18n.
var bannerI18NColumns = BannerI18NColumns{
	Id:          "id",
	BannerId:    "banner_id",
	Locale:      "locale",
	Title:       "title",
	Description: "description",
	ButtonText:  "button_text",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewBannerI18NDao creates and returns a new DAO object for table data access.
func NewBannerI18NDao() *BannerI18NDao {
	return &BannerI18NDao{
		group:   "default",
		table:   "banner_i18n",
		columns: bannerI18NColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BannerI18NDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BannerI18NDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BannerI18NDao) Columns() BannerI18NColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BannerI18NDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BannerI18NDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BannerI18NDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
