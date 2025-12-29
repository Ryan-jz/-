// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerDao is the data access object for table banner.
type BannerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns BannerColumns // columns contains all the column names of Table for convenient usage.
}

// BannerColumns defines and stores column names for table banner.
type BannerColumns struct {
	BannerId    string // 轮播图ID
	Title       string // 标题
	Description string // 描述
	MediaType   string // 媒体类型（1图片 2视频）
	MediaUrl    string // 媒体URL
	ButtonText  string // 按钮文字
	ButtonLink  string // 按钮链接
	SortOrder   string // 排序
	Status      string // 状态（0停用 1启用）
	Position    string // 位置（home首页 product产品页）
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// bannerColumns holds the columns for table banner.
var bannerColumns = BannerColumns{
	BannerId:    "banner_id",
	Title:       "title",
	Description: "description",
	MediaType:   "media_type",
	MediaUrl:    "media_url",
	ButtonText:  "button_text",
	ButtonLink:  "button_link",
	SortOrder:   "sort_order",
	Status:      "status",
	Position:    "position",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewBannerDao creates and returns a new DAO object for table data access.
func NewBannerDao() *BannerDao {
	return &BannerDao{
		group:   "default",
		table:   "banner",
		columns: bannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BannerDao) Columns() BannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
