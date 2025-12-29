package service

import (
	"context"
	"gf-admin/api/v1/recipe"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type IRecipeTag interface {
	GetList(ctx context.Context, keyword string, status *int, page, pageSize int) ([]recipe.TagListItem, int, error)
	GetAll(ctx context.Context) ([]recipe.TagListItem, error)
	Create(ctx context.Context, req *recipe.CreateTagReq) (uint, error)
	Update(ctx context.Context, req *recipe.UpdateTagReq) error
	Delete(ctx context.Context, id uint) error
}

type recipeTagImpl struct{}

var localRecipeTag IRecipeTag

func init() {
	RegisterRecipeTag(&recipeTagImpl{})
}

func RecipeTag() IRecipeTag {
	if localRecipeTag == nil {
		panic("implement not found for interface IRecipeTag, forgot register?")
	}
	return localRecipeTag
}

func RegisterRecipeTag(i IRecipeTag) {
	localRecipeTag = i
}

func (s *recipeTagImpl) GetList(ctx context.Context, keyword string, status *int, page, pageSize int) ([]recipe.TagListItem, int, error) {
	var list []recipe.TagListItem
	m := dao.RecipeTag.Ctx(ctx)

	if keyword != "" {
		m = m.WhereLike("name", "%"+keyword+"%")
	}
	if status != nil {
		m = m.Where("status", *status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	err = m.Order("sort_order asc, id desc").Page(page, pageSize).Scan(&list)
	return list, total, err
}

func (s *recipeTagImpl) GetAll(ctx context.Context) ([]recipe.TagListItem, error) {
	var list []recipe.TagListItem
	err := dao.RecipeTag.Ctx(ctx).Where("status", 1).Order("sort_order asc, id desc").Scan(&list)
	return list, err
}

func (s *recipeTagImpl) Create(ctx context.Context, req *recipe.CreateTagReq) (uint, error) {
	result, err := dao.RecipeTag.Ctx(ctx).Data(&do.RecipeTag{
		Name:      req.Name,
		Slug:      req.Slug,
		SortOrder: req.SortOrder,
		Status:    req.Status,
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

func (s *recipeTagImpl) Update(ctx context.Context, req *recipe.UpdateTagReq) error {
	count, err := dao.RecipeTag.Ctx(ctx).Where("id", req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("标签不存在")
	}

	_, err = dao.RecipeTag.Ctx(ctx).Data(&do.RecipeTag{
		Name:      req.Name,
		Slug:      req.Slug,
		SortOrder: req.SortOrder,
		Status:    req.Status,
	}).Where("id", req.Id).Update()
	return err
}

func (s *recipeTagImpl) Delete(ctx context.Context, id uint) error {
	count, err := dao.RecipeTag.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("标签不存在")
	}

	_, err = dao.RecipeTag.Ctx(ctx).Where("id", id).Delete()
	return err
}
