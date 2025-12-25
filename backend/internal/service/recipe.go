package service

import (
	"context"
	"encoding/json"
	"gf-admin/api/v1/recipe"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// IRecipe 食谱服务接口
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

// Recipe 获取食谱服务实例
func Recipe() IRecipe {
	if localRecipe == nil {
		panic("implement not found for interface IRecipe, forgot register?")
	}
	return localRecipe
}

// RegisterRecipe 注册食谱服务实现
func RegisterRecipe(i IRecipe) {
	localRecipe = i
}

// GetList 获取食谱列表
func (s *recipeImpl) GetList(ctx context.Context, keyword string, status *int, page, pageSize int) ([]recipe.ListItem, int, error) {
	var (
		list  []recipe.ListItem
		total int
		m     = g.DB().Model("recipe")
	)

	// 关键词搜索
	if keyword != "" {
		m = m.WhereLike("name", "%"+keyword+"%")
	}

	// 状态筛选
	if status != nil {
		m = m.Where("status", *status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = m.Order("sort_order asc, id desc").
		Page(page, pageSize).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetDetail 获取食谱详情
func (s *recipeImpl) GetDetail(ctx context.Context, id uint) (*recipe.DetailRes, error) {
	var detail recipe.DetailRes

	err := g.DB().Model("recipe").Where("id", id).Scan(&detail)
	if err != nil {
		return nil, err
	}

	if detail.Id == 0 {
		return nil, gerror.New("食谱不存在")
	}

	// 增加浏览次数
	_, _ = g.DB().Model("recipe").Where("id", id).Increment("view_count", 1)

	return &detail, nil
}

// Create 创建食谱
func (s *recipeImpl) Create(ctx context.Context, req *recipe.CreateReq) (uint, error) {
	// 转换JSON字段
	imagesJson, _ := json.Marshal(req.Images)
	ingredientsJson, _ := json.Marshal(req.Ingredients)

	data := g.Map{
		"name":         req.Name,
		"name_en":      req.NameEn,
		"subtitle":     req.Subtitle,
		"description":  req.Description,
		"image":        req.Image,
		"images":       string(imagesJson),
		"ingredients":  string(ingredientsJson),
		"content":      req.Content,
		"cooking_time": req.CookingTime,
		"difficulty":   req.Difficulty,
		"servings":     req.Servings,
		"product_ids":  req.ProductIds,
		"tags":         req.Tags,
		"sort_order":   req.SortOrder,
		"status":       req.Status,
	}

	result, err := g.DB().Model("recipe").Data(data).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

// Update 更新食谱
func (s *recipeImpl) Update(ctx context.Context, req *recipe.UpdateReq) error {
	// 检查食谱是否存在
	count, err := g.DB().Model("recipe").Where("id", req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("食谱不存在")
	}

	// 转换JSON字段
	imagesJson, _ := json.Marshal(req.Images)
	ingredientsJson, _ := json.Marshal(req.Ingredients)

	data := g.Map{
		"name":         req.Name,
		"name_en":      req.NameEn,
		"subtitle":     req.Subtitle,
		"description":  req.Description,
		"image":        req.Image,
		"images":       string(imagesJson),
		"ingredients":  string(ingredientsJson),
		"content":      req.Content,
		"cooking_time": req.CookingTime,
		"difficulty":   req.Difficulty,
		"servings":     req.Servings,
		"product_ids":  req.ProductIds,
		"tags":         req.Tags,
		"sort_order":   req.SortOrder,
		"status":       req.Status,
	}

	_, err = g.DB().Model("recipe").Data(data).Where("id", req.Id).Update()
	return err
}

// Delete 删除食谱
func (s *recipeImpl) Delete(ctx context.Context, id uint) error {
	// 检查食谱是否存在
	count, err := g.DB().Model("recipe").Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("食谱不存在")
	}

	_, err = g.DB().Model("recipe").Where("id", id).Delete()
	return err
}
