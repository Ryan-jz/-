// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecipeTagDao is the data access object for table recipe_tag.
type RecipeTagDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns RecipeTagColumns // columns contains all the column names of Table for convenient usage.
}

// RecipeTagColumns defines and stores column names for table recipe_tag.
type RecipeTagColumns struct {
	Id        string //
	Name      string // 标签名称
	Slug      string // 标签标识
	SortOrder string // 排序
	Status    string // 状态 1启用 0禁用
	CreatedAt string //
	UpdatedAt string //
}

// recipeTagColumns holds the columns for table recipe_tag.
var recipeTagColumns = RecipeTagColumns{
	Id:        "id",
	Name:      "name",
	Slug:      "slug",
	SortOrder: "sort_order",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRecipeTagDao creates and returns a new DAO object for table data access.
func NewRecipeTagDao() *RecipeTagDao {
	return &RecipeTagDao{
		group:   "default",
		table:   "recipe_tag",
		columns: recipeTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecipeTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecipeTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecipeTagDao) Columns() RecipeTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecipeTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecipeTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecipeTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
