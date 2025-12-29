package service

import (
	"context"
	"encoding/json"
	"gf-admin/api/v1/recipe"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type IRecipe interface {
	GetList(ctx context.Context, keyword string, status *int, page, pageSize int) ([]recipe.ListItem, int, error)
	GetDetail(ctx context.Context, id uint) (*recipe.DetailRes, error)
	Create(ctx context.Context, req *recipe.CreateReq) (uint, error)
	Update(ctx context.Context, req *recipe.UpdateReq) error
	Delete(ctx context.Context, id uint) error
}

type recipeImpl struct{}

var localRecipe IRecipe

func init() {
	RegisterRecipe(&recipeImpl{})
}

func Recipe() IRecipe {
	if localRecipe == nil {
		panic("implement not found for interface IRecipe, forgot register?")
	}
	return localRecipe
}

func RegisterRecipe(i IRecipe) {
	localRecipe = i
}

func (s *recipeImpl) GetList(ctx context.Context, keyword string, status *int, page, pageSize int) ([]recipe.ListItem, int, error) {
	var list []recipe.ListItem
	m := dao.Recipe.Ctx(ctx)

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

func (s *recipeImpl) GetDetail(ctx context.Context, id uint) (*recipe.DetailRes, error) {
	var detail recipe.DetailRes
	err := dao.Recipe.Ctx(ctx).Where("id", id).Scan(&detail)
	if err != nil {
		return nil, err
	}
	if detail.Id == 0 {
		return nil, gerror.New("食谱不存在")
	}

	dao.Recipe.Ctx(ctx).Where("id", id).Increment("view_count", 1)
	return &detail, nil
}

func (s *recipeImpl) Create(ctx context.Context, req *recipe.CreateReq) (uint, error) {
	imagesJson, _ := json.Marshal(req.Images)
	ingredientsJson, _ := json.Marshal(req.Ingredients)

	result, err := dao.Recipe.Ctx(ctx).Data(&do.Recipe{
		Name:        req.Name,
		NameEn:      req.NameEn,
		Subtitle:    req.Subtitle,
		Description: req.Description,
		Image:       req.Image,
		Images:      string(imagesJson),
		Ingredients: string(ingredientsJson),
		Content:     req.Content,
		CookingTime: req.CookingTime,
		Difficulty:  req.Difficulty,
		Servings:    req.Servings,
		ProductIds:  req.ProductIds,
		Tags:        req.Tags,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

func (s *recipeImpl) Update(ctx context.Context, req *recipe.UpdateReq) error {
	count, err := dao.Recipe.Ctx(ctx).Where("id", req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("食谱不存在")
	}

	imagesJson, _ := json.Marshal(req.Images)
	ingredientsJson, _ := json.Marshal(req.Ingredients)

	_, err = dao.Recipe.Ctx(ctx).Data(&do.Recipe{
		Name:        req.Name,
		NameEn:      req.NameEn,
		Subtitle:    req.Subtitle,
		Description: req.Description,
		Image:       req.Image,
		Images:      string(imagesJson),
		Ingredients: string(ingredientsJson),
		Content:     req.Content,
		CookingTime: req.CookingTime,
		Difficulty:  req.Difficulty,
		Servings:    req.Servings,
		ProductIds:  req.ProductIds,
		Tags:        req.Tags,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
	}).Where("id", req.Id).Update()
	return err
}

func (s *recipeImpl) Delete(ctx context.Context, id uint) error {
	count, err := dao.Recipe.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("食谱不存在")
	}

	_, err = dao.Recipe.Ctx(ctx).Where("id", id).Delete()
	return err
}
