// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecipeI18NDao is the data access object for table recipe_i18n.
type RecipeI18NDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns RecipeI18NColumns // columns contains all the column names of Table for convenient usage.
}

// RecipeI18NColumns defines and stores column names for table recipe_i18n.
type RecipeI18NColumns struct {
	Id          string // ID
	RecipeId    string // 食谱ID
	Locale      string // 语言代码: zh-CN, en-US, de-DE
	Name        string // 食谱名称
	Subtitle    string // 副标题
	Description string // 简介
	Ingredients string // 食材列表(JSON数组)
	Content     string // 制作步骤(富文本HTML)
	Tags        string // 标签（逗号分隔）
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// recipeI18NColumns holds the columns for table recipe_i18n.
var recipeI18NColumns = RecipeI18NColumns{
	Id:          "id",
	RecipeId:    "recipe_id",
	Locale:      "locale",
	Name:        "name",
	Subtitle:    "subtitle",
	Description: "description",
	Ingredients: "ingredients",
	Content:     "content",
	Tags:        "tags",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewRecipeI18NDao creates and returns a new DAO object for table data access.
func NewRecipeI18NDao() *RecipeI18NDao {
	return &RecipeI18NDao{
		group:   "default",
		table:   "recipe_i18n",
		columns: recipeI18NColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecipeI18NDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecipeI18NDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecipeI18NDao) Columns() RecipeI18NColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecipeI18NDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecipeI18NDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecipeI18NDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
