package service

import (
	"context"
	"encoding/json"
	"gf-admin/api/v1/product"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type IProduct interface {
	GetCategoryList(ctx context.Context, status *int) ([]any, error)
	CreateCategory(ctx context.Context, req any) (uint, error)
	UpdateCategory(ctx context.Context, req any) error
	DeleteCategory(ctx context.Context, id uint) error
	GetCategoryWithProducts(ctx context.Context, status *int) ([]any, error)
	GetList(ctx context.Context, categoryId *uint, keyword string, status *int, page, pageSize int) ([]any, int, error)
	GetDetail(ctx context.Context, id uint) (any, error)
	Create(ctx context.Context, req any) (uint, error)
	Update(ctx context.Context, req any) error
	Delete(ctx context.Context, id uint) error
}

type productImpl struct{}

func init() {
	RegisterProduct(&productImpl{})
}

func (s *productImpl) GetCategoryList(ctx context.Context, status *int) ([]any, error) {
	model := dao.ProductCategory.Ctx(ctx)

	if status != nil {
		model = model.Where("status", *status)
	}

	var list []*entity.ProductCategory
	err := model.Order("sort_order ASC, id ASC").Scan(&list)

	result := make([]any, len(list))
	for i, v := range list {
		result[i] = v
	}
	return result, err
}

func (s *productImpl) CreateCategory(ctx context.Context, req any) (uint, error) {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	result, err := dao.ProductCategory.Ctx(ctx).Data(reqMap).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

func (s *productImpl) UpdateCategory(ctx context.Context, req any) error {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	id := reqMap["id"]
	delete(reqMap, "id")

	count, err := dao.ProductCategory.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("分类不存在")
	}

	_, err = dao.ProductCategory.Ctx(ctx).Data(reqMap).Where("id", id).Update()
	return err
}

func (s *productImpl) DeleteCategory(ctx context.Context, id uint) error {
	count, err := dao.ProductCategory.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("分类不存在")
	}

	productCount, err := dao.Product.Ctx(ctx).Where("category_id", id).Count()
	if err != nil {
		return err
	}
	if productCount > 0 {
		return gerror.New("该分类下还有产品，无法删除")
	}

	_, err = dao.ProductCategory.Ctx(ctx).Where("id", id).Delete()
	return err
}

func (s *productImpl) GetCategoryWithProducts(ctx context.Context, status *int) ([]any, error) {
	m := dao.ProductCategory.Ctx(ctx)
	if status != nil {
		m = m.Where("status", *status)
	}

	var categories []product.CategoryWithProducts
	err := m.Order("sort_order asc, id asc").Scan(&categories)
	if err != nil {
		return nil, err
	}

	for i := range categories {
		pm := dao.Product.Ctx(ctx).Where("category_id", categories[i].Id).Where("status", 1)
		var products []product.ProductItem
		pm.Order("sort_order asc, id desc").Scan(&products)
		categories[i].Products = products
	}

	result := make([]any, len(categories))
	for i, v := range categories {
		result[i] = v
	}
	return result, nil
}

func (s *productImpl) GetList(ctx context.Context, categoryId *uint, keyword string, status *int, page, pageSize int) ([]any, int, error) {
	model := dao.Product.Ctx(ctx)

	if categoryId != nil {
		model = model.Where("category_id", *categoryId)
	}
	if keyword != "" {
		model = model.WhereLike("name", "%"+keyword+"%")
	}
	if status != nil {
		model = model.Where("status", *status)
	}

	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	var list []*entity.Product
	err = model.Order("sort_order ASC, id DESC").Limit(pageSize).Offset(offset).Scan(&list)

	result := make([]any, len(list))
	for i, v := range list {
		result[i] = v
	}
	return result, total, err
}

func (s *productImpl) GetDetail(ctx context.Context, id uint) (any, error) {
	var detail *entity.Product
	err := dao.Product.Ctx(ctx).Where("id", id).Scan(&detail)
	if err != nil {
		return nil, err
	}

	dao.Product.Ctx(ctx).Where("id", id).Increment("view_count", 1)
	return detail, nil
}

func (s *productImpl) Create(ctx context.Context, req any) (uint, error) {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	if reqMap["images"] != nil {
		imagesJson, _ := json.Marshal(reqMap["images"])
		reqMap["images"] = string(imagesJson)
	}
	if reqMap["features"] != nil {
		featuresJson, _ := json.Marshal(reqMap["features"])
		reqMap["features"] = string(featuresJson)
	}

	result, err := dao.Product.Ctx(ctx).Data(reqMap).Insert()
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return uint(id), nil
}

func (s *productImpl) Update(ctx context.Context, req any) error {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	id := reqMap["id"]
	delete(reqMap, "id")

	count, err := dao.Product.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("产品不存在")
	}

	if reqMap["images"] != nil {
		imagesJson, _ := json.Marshal(reqMap["images"])
		reqMap["images"] = string(imagesJson)
	}
	if reqMap["features"] != nil {
		featuresJson, _ := json.Marshal(reqMap["features"])
		reqMap["features"] = string(featuresJson)
	}

	_, err = dao.Product.Ctx(ctx).Data(reqMap).Where("id", id).Update()
	return err
}

func (s *productImpl) Delete(ctx context.Context, id uint) error {
	count, err := dao.Product.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("产品不存在")
	}

	_, err = dao.Product.Ctx(ctx).Where("id", id).Delete()
	return err
}
