package service

import (
	"context"
	"encoding/json"
	"gf-admin/internal/middleware"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// RecipeI18nService 食谱国际化服务
type RecipeI18nService struct{}

// GetRecipeListI18n 获取国际化食谱列表
func (s *RecipeI18nService) GetRecipeListI18n(ctx context.Context, keyword string, difficulty *int, status *int, page, pageSize int) ([]g.Map, int, error) {
	locale := middleware.GetLocale(ctx)

	var (
		recipes []g.Map
		total   int
		m       = g.DB().Model("recipe r")
	)

	// 左连接国际化表
	m = m.LeftJoin("recipe_i18n ri", "r.id = ri.recipe_id AND ri.locale = ?", locale)

	// 关键词搜索
	if keyword != "" {
		m = m.Where(m.Builder().Where("ri.name LIKE ?", "%"+keyword+"%").
			WhereOr("r.name LIKE ?", "%"+keyword+"%"))
	}

	// 难度筛选
	if difficulty != nil {
		m = m.Where("r.difficulty", *difficulty)
	}

	// 状态筛选
	if status != nil {
		m = m.Where("r.status", *status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	// 选择字段
	m = m.Fields(
		"r.id",
		"r.image",
		"r.cooking_time",
		"r.difficulty",
		"r.servings",
		"r.view_count",
		"r.like_count",
		"r.sort_order",
		"r.status",
		"r.create_time",
		"r.update_time",
		"COALESCE(ri.name, r.name) as name",
		"COALESCE(ri.subtitle, r.subtitle) as subtitle",
		"COALESCE(ri.description, r.description) as description",
		"COALESCE(ri.tags, r.tags) as tags",
	)

	// 分页查询
	err = m.Order("r.sort_order asc, r.id desc").
		Page(page, pageSize).
		Scan(&recipes)
	if err != nil {
		return nil, 0, err
	}

	return recipes, total, nil
}

// GetRecipeDetailI18n 获取国际化食谱详情
func (s *RecipeI18nService) GetRecipeDetailI18n(ctx context.Context, id uint) (g.Map, error) {
	locale := middleware.GetLocale(ctx)

	var detail g.Map

	m := g.DB().Model("recipe r").
		LeftJoin("recipe_i18n ri", "r.id = ri.recipe_id AND ri.locale = ?", locale).
		Where("r.id", id)

	// 选择所有需要的字段
	m = m.Fields(
		"r.id",
		"r.image",
		"r.images",
		"r.cooking_time",
		"r.difficulty",
		"r.servings",
		"r.product_ids",
		"r.view_count",
		"r.like_count",
		"r.sort_order",
		"r.status",
		"r.create_time",
		"r.update_time",
		"COALESCE(ri.name, r.name) as name",
		"COALESCE(ri.subtitle, r.subtitle) as subtitle",
		"COALESCE(ri.description, r.description) as description",
		"COALESCE(ri.ingredients, r.ingredients) as ingredients",
		"COALESCE(ri.content, r.content) as content",
		"COALESCE(ri.tags, r.tags) as tags",
	)

	err := m.Scan(&detail)
	if err != nil {
		return nil, err
	}

	if len(detail) == 0 {
		return nil, gerror.New("食谱不存在")
	}

	// 增加浏览次数
	_, _ = g.DB().Model("recipe").Where("id", id).Increment("view_count", 1)

	// 解析JSON字段
	if detail["images"] != nil && detail["images"].(string) != "" {
		var images []string
		json.Unmarshal([]byte(detail["images"].(string)), &images)
		detail["images"] = images
	}

	if detail["ingredients"] != nil && detail["ingredients"].(string) != "" {
		var ingredients []map[string]interface{}
		json.Unmarshal([]byte(detail["ingredients"].(string)), &ingredients)
		detail["ingredients"] = ingredients
	}

	return detail, nil
}

// SaveRecipeI18n 保存食谱国际化数据
func (s *RecipeI18nService) SaveRecipeI18n(ctx context.Context, recipeId uint, locale string, data g.Map) error {
	// 检查是否已存在
	count, err := g.DB().Model("recipe_i18n").
		Where("recipe_id", recipeId).
		Where("locale", locale).
		Count()
	if err != nil {
		return err
	}

	i18nData := g.Map{
		"recipe_id":   recipeId,
		"locale":      locale,
		"name":        data["name"],
		"subtitle":    data["subtitle"],
		"description": data["description"],
		"content":     data["content"],
		"tags":        data["tags"],
	}

	// 处理ingredients（数组转JSON）
	if data["ingredients"] != nil {
		ingredientsJson, _ := json.Marshal(data["ingredients"])
		i18nData["ingredients"] = string(ingredientsJson)
	}

	if count > 0 {
		// 更新
		_, err = g.DB().Model("recipe_i18n").
			Data(i18nData).
			Where("recipe_id", recipeId).
			Where("locale", locale).
			Update()
	} else {
		// 插入
		_, err = g.DB().Model("recipe_i18n").Data(i18nData).Insert()
	}

	return err
}

// GetRecipeI18nData 获取食谱的所有语言数据（用于管理后台编辑）
func (s *RecipeI18nService) GetRecipeI18nData(ctx context.Context, recipeId uint) (map[string]g.Map, error) {
	var i18nList []g.Map

	err := g.DB().Model("recipe_i18n").
		Where("recipe_id", recipeId).
		Scan(&i18nList)
	if err != nil {
		return nil, err
	}

	result := make(map[string]g.Map)
	for _, item := range i18nList {
		locale := item["locale"].(string)

		// 解析ingredients JSON
		if item["ingredients"] != nil && item["ingredients"].(string) != "" {
			var ingredients []map[string]interface{}
			json.Unmarshal([]byte(item["ingredients"].(string)), &ingredients)
			item["ingredients"] = ingredients
		}

		result[locale] = item
	}

	return result, nil
}

// BatchSaveRecipeI18n 批量保存食谱国际化数据（用于管理后台）
func (s *RecipeI18nService) BatchSaveRecipeI18n(ctx context.Context, recipeId uint, i18nData map[string]g.Map) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for locale, data := range i18nData {
			if err := s.SaveRecipeI18n(ctx, recipeId, locale, data); err != nil {
				return err
			}
		}
		return nil
	})
}
