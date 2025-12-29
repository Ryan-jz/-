// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecipeTagRelationDao is the data access object for table recipe_tag_relation.
type RecipeTagRelationDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns RecipeTagRelationColumns // columns contains all the column names of Table for convenient usage.
}

// RecipeTagRelationColumns defines and stores column names for table recipe_tag_relation.
type RecipeTagRelationColumns struct {
	Id        string //
	RecipeId  string // 食谱ID
	TagId     string // 标签ID
	CreatedAt string //
}

// recipeTagRelationColumns holds the columns for table recipe_tag_relation.
var recipeTagRelationColumns = RecipeTagRelationColumns{
	Id:        "id",
	RecipeId:  "recipe_id",
	TagId:     "tag_id",
	CreatedAt: "created_at",
}

// NewRecipeTagRelationDao creates and returns a new DAO object for table data access.
func NewRecipeTagRelationDao() *RecipeTagRelationDao {
	return &RecipeTagRelationDao{
		group:   "default",
		table:   "recipe_tag_relation",
		columns: recipeTagRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecipeTagRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecipeTagRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecipeTagRelationDao) Columns() RecipeTagRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecipeTagRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecipeTagRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecipeTagRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
