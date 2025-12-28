package service

import (
	"context"
	"encoding/json"
	"gf-admin/internal/middleware"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductI18nService 产品国际化服务
type ProductI18nService struct{}

// GetCategoryListI18n 获取国际化产品分类列表
func (s *ProductI18nService) GetCategoryListI18n(ctx context.Context, status *int) ([]g.Map, error) {
	locale := middleware.GetLocale(ctx)

	var (
		categories []g.Map
		m          = g.DB().Model("product_category pc")
	)

	// 左连接国际化表
	m = m.LeftJoin("product_category_i18n pci", "pc.id = pci.category_id AND pci.locale = ?", locale)

	if status != nil {
		m = m.Where("pc.status", *status)
	}

	// 选择字段，优先使用国际化数据
	m = m.Fields(
		"pc.id",
		"pc.slug",
		"pc.sort_order",
		"pc.status",
		"pc.created_at",
		"pc.updated_at",
		"COALESCE(pci.name, pc.name) as name",
		"COALESCE(pci.description, pc.description) as description",
	)

	err := m.Order("pc.sort_order asc, pc.id asc").Scan(&categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetProductListI18n 获取国际化产品列表
func (s *ProductI18nService) GetProductListI18n(ctx context.Context, categoryId *uint, keyword string, status *int, page, pageSize int) ([]g.Map, int, error) {
	locale := middleware.GetLocale(ctx)

	var (
		products []g.Map
		total    int
		m        = g.DB().Model("product p")
	)

	// 左连接国际化表
	m = m.LeftJoin("product_i18n pi", "p.id = pi.product_id AND pi.locale = ?", locale)

	// 分类筛选
	if categoryId != nil {
		m = m.Where("p.category_id", *categoryId)
	}

	// 关键词搜索（搜索国际化名称）
	if keyword != "" {
		m = m.Where(m.Builder().Where("pi.name LIKE ?", "%"+keyword+"%").
			WhereOr("p.name LIKE ?", "%"+keyword+"%"))
	}

	// 状态筛选
	if status != nil {
		m = m.Where("p.status", *status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	// 选择字段
	m = m.Fields(
		"p.id",
		"p.category_id",
		"p.image",
		"p.price",
		"p.stock",
		"p.weight",
		"p.sort_order",
		"p.status",
		"p.view_count",
		"p.created_at",
		"p.updated_at",
		"COALESCE(pi.name, p.name) as name",
		"COALESCE(pi.subtitle, p.subtitle) as subtitle",
		"COALESCE(pi.description, p.description) as description",
	)

	// 分页查询
	err = m.Order("p.sort_order asc, p.id desc").
		Page(page, pageSize).
		Scan(&products)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductDetailI18n 获取国际化产品详情
func (s *ProductI18nService) GetProductDetailI18n(ctx context.Context, id uint) (g.Map, error) {
	locale := middleware.GetLocale(ctx)

	var detail g.Map

	m := g.DB().Model("product p").
		LeftJoin("product_i18n pi", "p.id = pi.product_id AND pi.locale = ?", locale).
		Where("p.id", id)

	// 选择所有需要的字段
	m = m.Fields(
		"p.id",
		"p.category_id",
		"p.image",
		"p.images",
		"p.price",
		"p.stock",
		"p.weight",
		"p.nutrition",
		"p.sort_order",
		"p.status",
		"p.view_count",
		"p.created_at",
		"p.updated_at",
		"COALESCE(pi.name, p.name) as name",
		"COALESCE(pi.subtitle, p.subtitle) as subtitle",
		"COALESCE(pi.description, p.description) as description",
		"COALESCE(pi.ingredients, p.ingredients) as ingredients",
		"COALESCE(pi.usage, p.usage) as usage",
		"COALESCE(pi.features, p.features) as features",
	)

	err := m.Scan(&detail)
	if err != nil {
		return nil, err
	}

	if len(detail) == 0 {
		return nil, gerror.New("产品不存在")
	}

	// 增加浏览次数
	_, _ = g.DB().Model("product").Where("id", id).Increment("view_count", 1)

	// 解析JSON字段
	if detail["images"] != nil {
		var images []string
		json.Unmarshal([]byte(detail["images"].(string)), &images)
		detail["images"] = images
	}

	if detail["features"] != nil && detail["features"].(string) != "" {
		var features []string
		json.Unmarshal([]byte(detail["features"].(string)), &features)
		detail["features"] = features
	}

	return detail, nil
}

// SaveProductI18n 保存产品国际化数据
func (s *ProductI18nService) SaveProductI18n(ctx context.Context, productId uint, locale string, data g.Map) error {
	// 检查是否已存在
	count, err := g.DB().Model("product_i18n").
		Where("product_id", productId).
		Where("locale", locale).
		Count()
	if err != nil {
		return err
	}

	i18nData := g.Map{
		"product_id":  productId,
		"locale":      locale,
		"name":        data["name"],
		"subtitle":    data["subtitle"],
		"description": data["description"],
		"ingredients": data["ingredients"],
		"usage":       data["usage"],
	}

	// 处理features（数组转JSON）
	if data["features"] != nil {
		featuresJson, _ := json.Marshal(data["features"])
		i18nData["features"] = string(featuresJson)
	}

	if count > 0 {
		// 更新
		_, err = g.DB().Model("product_i18n").
			Data(i18nData).
			Where("product_id", productId).
			Where("locale", locale).
			Update()
	} else {
		// 插入
		_, err = g.DB().Model("product_i18n").Data(i18nData).Insert()
	}

	return err
}

// SaveCategoryI18n 保存分类国际化数据
func (s *ProductI18nService) SaveCategoryI18n(ctx context.Context, categoryId uint, locale string, data g.Map) error {
	// 检查是否已存在
	count, err := g.DB().Model("product_category_i18n").
		Where("category_id", categoryId).
		Where("locale", locale).
		Count()
	if err != nil {
		return err
	}

	i18nData := g.Map{
		"category_id": categoryId,
		"locale":      locale,
		"name":        data["name"],
		"description": data["description"],
	}

	if count > 0 {
		// 更新
		_, err = g.DB().Model("product_category_i18n").
			Data(i18nData).
			Where("category_id", categoryId).
			Where("locale", locale).
			Update()
	} else {
		// 插入
		_, err = g.DB().Model("product_category_i18n").Data(i18nData).Insert()
	}

	return err
}

// GetProductI18nData 获取产品的所有语言数据（用于管理后台编辑）
func (s *ProductI18nService) GetProductI18nData(ctx context.Context, productId uint) (map[string]g.Map, error) {
	var i18nList []g.Map

	err := g.DB().Model("product_i18n").
		Where("product_id", productId).
		Scan(&i18nList)
	if err != nil {
		return nil, err
	}

	result := make(map[string]g.Map)
	for _, item := range i18nList {
		locale := item["locale"].(string)

		// 解析features JSON
		if item["features"] != nil && item["features"].(string) != "" {
			var features []string
			json.Unmarshal([]byte(item["features"].(string)), &features)
			item["features"] = features
		}

		result[locale] = item
	}

	return result, nil
}

// GetCategoryI18nData 获取分类的所有语言数据（用于管理后台编辑）
func (s *ProductI18nService) GetCategoryI18nData(ctx context.Context, categoryId uint) (map[string]g.Map, error) {
	var i18nList []g.Map

	err := g.DB().Model("product_category_i18n").
		Where("category_id", categoryId).
		Scan(&i18nList)
	if err != nil {
		return nil, err
	}

	result := make(map[string]g.Map)
	for _, item := range i18nList {
		locale := item["locale"].(string)
		result[locale] = item
	}

	return result, nil
}

// BatchSaveProductI18n 批量保存产品国际化数据（用于管理后台）
func (s *ProductI18nService) BatchSaveProductI18n(ctx context.Context, productId uint, i18nData map[string]g.Map) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for locale, data := range i18nData {
			if err := s.SaveProductI18n(ctx, productId, locale, data); err != nil {
				return err
			}
		}
		return nil
	})
}

// BatchSaveCategoryI18n 批量保存分类国际化数据（用于管理后台）
func (s *ProductI18nService) BatchSaveCategoryI18n(ctx context.Context, categoryId uint, i18nData map[string]g.Map) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for locale, data := range i18nData {
			if err := s.SaveCategoryI18n(ctx, categoryId, locale, data); err != nil {
				return err
			}
		}
		return nil
	})
}
