package service

import (
	"context"
	"encoding/json"
	"gf-admin/api/v1/product"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// IProduct 产品服务接口
type IProduct interface {
	GetCategoryList(ctx context.Context, status *int) ([]product.CategoryItem, error)
	CreateCategory(ctx context.Context, req interface{}) (uint, error)
	UpdateCategory(ctx context.Context, req interface{}) error
	DeleteCategory(ctx context.Context, id uint) error
	GetList(ctx context.Context, categoryId *uint, keyword string, status *int, page, pageSize int) ([]product.ListItem, int, error)
	GetDetail(ctx context.Context, id uint) (*product.DetailRes, error)
	Create(ctx context.Context, req *product.CreateReq) (uint, error)
	Update(ctx context.Context, req *product.UpdateReq) error
	Delete(ctx context.Context, id uint) error
}

type productImpl struct{}

func init() {
	RegisterProduct(New())
}

func New() IProduct {
	return &productImpl{}
}

// GetCategoryList 获取产品分类列表
func (s *productImpl) GetCategoryList(ctx context.Context, status *int) ([]product.CategoryItem, error) {
	var (
		list []product.CategoryItem
		m    = g.DB().Model("product_category")
	)

	if status != nil {
		m = m.Where("status", *status)
	}

	err := m.Order("sort_order asc, id asc").Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// CreateCategory 创建产品分类
func (s *productImpl) CreateCategory(ctx context.Context, req interface{}) (uint, error) {
	result, err := g.DB().Model("product_category").Data(req).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

// UpdateCategory 更新产品分类
func (s *productImpl) UpdateCategory(ctx context.Context, req interface{}) error {
	// 将 interface{} 转换为 map
	var reqMap g.Map
	
	// 使用类型断言或 JSON 转换
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	
	if err := json.Unmarshal(jsonBytes, &reqMap); err != nil {
		return err
	}

	id := reqMap["id"]
	delete(reqMap, "id")

	// 检查分类是否存在
	count, err := g.DB().Model("product_category").Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("分类不存在")
	}

	_, err = g.DB().Model("product_category").Data(reqMap).Where("id", id).Update()
	return err
}

// DeleteCategory 删除产品分类
func (s *productImpl) DeleteCategory(ctx context.Context, id uint) error {
	// 检查分类是否存在
	count, err := g.DB().Model("product_category").Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("分类不存在")
	}

	// 检查是否有产品使用该分类
	productCount, err := g.DB().Model("product").Where("category_id", id).Count()
	if err != nil {
		return err
	}
	if productCount > 0 {
		return gerror.New("该分类下还有产品，无法删除")
	}

	_, err = g.DB().Model("product_category").Where("id", id).Delete()
	return err
}

// GetList 获取产品列表
func (s *productImpl) GetList(ctx context.Context, categoryId *uint, keyword string, status *int, page, pageSize int) ([]product.ListItem, int, error) {
	var (
		list  []product.ListItem
		total int
		m     = g.DB().Model("product")
	)

	// 分类筛选
	if categoryId != nil {
		m = m.Where("category_id", *categoryId)
	}

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

// GetDetail 获取产品详情
func (s *productImpl) GetDetail(ctx context.Context, id uint) (*product.DetailRes, error) {
	var detail product.DetailRes

	err := g.DB().Model("product").Where("id", id).Scan(&detail)
	if err != nil {
		return nil, err
	}

	if detail.Id == 0 {
		return nil, gerror.New("产品不存在")
	}

	// 增加浏览次数
	_, _ = g.DB().Model("product").Where("id", id).Increment("view_count", 1)

	return &detail, nil
}

// Create 创建产品
func (s *productImpl) Create(ctx context.Context, req *product.CreateReq) (uint, error) {
	// 转换JSON字段
	imagesJson, _ := json.Marshal(req.Images)
	featuresJson, _ := json.Marshal(req.Features)

	data := g.Map{
		"category_id": req.CategoryId,
		"name":        req.Name,
		"name_en":     req.NameEn,
		"subtitle":    req.Subtitle,
		"description": req.Description,
		"image":       req.Image,
		"images":      string(imagesJson),
		"price":       req.Price,
		"stock":       req.Stock,
		"weight":      req.Weight,
		"ingredients": req.Ingredients,
		"nutrition":   req.Nutrition,
		"usage":       req.Usage,
		"features":    string(featuresJson),
		"sort_order":  req.SortOrder,
		"status":      req.Status,
	}

	result, err := g.DB().Model("product").Data(data).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

// Update 更新产品
func (s *productImpl) Update(ctx context.Context, req *product.UpdateReq) error {
	// 检查产品是否存在
	count, err := g.DB().Model("product").Where("id", req.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("产品不存在")
	}

	// 转换JSON字段
	imagesJson, _ := json.Marshal(req.Images)
	featuresJson, _ := json.Marshal(req.Features)

	data := g.Map{
		"category_id": req.CategoryId,
		"name":        req.Name,
		"name_en":     req.NameEn,
		"subtitle":    req.Subtitle,
		"description": req.Description,
		"image":       req.Image,
		"images":      string(imagesJson),
		"price":       req.Price,
		"stock":       req.Stock,
		"weight":      req.Weight,
		"ingredients": req.Ingredients,
		"nutrition":   req.Nutrition,
		"usage":       req.Usage,
		"features":    string(featuresJson),
		"sort_order":  req.SortOrder,
		"status":      req.Status,
	}

	_, err = g.DB().Model("product").Data(data).Where("id", req.Id).Update()
	return err
}

// Delete 删除产品
func (s *productImpl) Delete(ctx context.Context, id uint) error {
	// 检查产品是否存在
	count, err := g.DB().Model("product").Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("产品不存在")
	}

	_, err = g.DB().Model("product").Where("id", id).Delete()
	return err
}
