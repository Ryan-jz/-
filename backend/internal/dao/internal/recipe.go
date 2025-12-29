// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecipeDao is the data access object for table recipe.
type RecipeDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns RecipeColumns // columns contains all the column names of Table for convenient usage.
}

// RecipeColumns defines and stores column names for table recipe.
type RecipeColumns struct {
	Id          string // 食谱ID
	Name        string // 食谱名称
	NameEn      string // 英文名称
	Subtitle    string // 副标题
	Description string // 简介
	Image       string // 主图片
	Images      string // 详情图片（JSON数组）
	Ingredients string // 食材列表（JSON数组）
	Content     string // 制作步骤（富文本HTML）
	CookingTime string // 烹饪时间（分钟）
	Difficulty  string // 难度等级（1简单 2中等 3困难）
	Servings    string // 份量（几人份）
	ProductIds  string // 关联产品ID（逗号分隔）
	Tags        string // 标签（逗号分隔）
	ViewCount   string // 浏览次数
	LikeCount   string // 点赞次数
	SortOrder   string // 排序
	Status      string // 状态（0下架 1上架）
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// recipeColumns holds the columns for table recipe.
var recipeColumns = RecipeColumns{
	Id:          "id",
	Name:        "name",
	NameEn:      "name_en",
	Subtitle:    "subtitle",
	Description: "description",
	Image:       "image",
	Images:      "images",
	Ingredients: "ingredients",
	Content:     "content",
	CookingTime: "cooking_time",
	Difficulty:  "difficulty",
	Servings:    "servings",
	ProductIds:  "product_ids",
	Tags:        "tags",
	ViewCount:   "view_count",
	LikeCount:   "like_count",
	SortOrder:   "sort_order",
	Status:      "status",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewRecipeDao creates and returns a new DAO object for table data access.
func NewRecipeDao() *RecipeDao {
	return &RecipeDao{
		group:   "default",
		table:   "recipe",
		columns: recipeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecipeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecipeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecipeDao) Columns() RecipeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecipeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecipeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecipeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
