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

	type RawItem struct {
		Id             uint   `json:"id"`
		Name           string `json:"name"`
		NameEn         string `json:"nameEn"`
		Subtitle       string `json:"subtitle"`
		Description    string `json:"description"`
		Image          string `json:"image"`
		ImagesStr      string `orm:"images"`
		IngredientsStr string `orm:"ingredients"`
		Content        string `json:"content"`
		CookingTime    int    `json:"cookingTime"`
		Difficulty     int    `json:"difficulty"`
		Servings       int    `json:"servings"`
		ProductIds     string `json:"productIds"`
		Tags           string `json:"tags"`
		ViewCount      int    `json:"viewCount"`
		LikeCount      int    `json:"likeCount"`
		SortOrder      int    `json:"sortOrder"`
		Status         int    `json:"status"`
	}

	var rawList []RawItem
	err = m.Order("sort_order asc, id desc").Page(page, pageSize).Scan(&rawList)
	if err != nil {
		return nil, 0, err
	}

	list := make([]recipe.ListItem, 0, len(rawList))
	for _, item := range rawList {
		listItem := recipe.ListItem{
			Id:          item.Id,
			Name:        item.Name,
			NameEn:      item.NameEn,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       item.Image,
			Content:     item.Content,
			CookingTime: item.CookingTime,
			Difficulty:  item.Difficulty,
			Servings:    item.Servings,
			ProductIds:  item.ProductIds,
			Tags:        item.Tags,
			ViewCount:   item.ViewCount,
			LikeCount:   item.LikeCount,
			SortOrder:   item.SortOrder,
			Status:      item.Status,
		}

		if item.ImagesStr != "" && item.ImagesStr != "null" {
			json.Unmarshal([]byte(item.ImagesStr), &listItem.Images)
		}
		if item.IngredientsStr != "" && item.IngredientsStr != "null" {
			json.Unmarshal([]byte(item.IngredientsStr), &listItem.Ingredients)
		}

		list = append(list, listItem)
	}

	return list, total, nil
}

func (s *recipeImpl) GetDetail(ctx context.Context, id uint) (*recipe.DetailRes, error) {
	type RawDetail struct {
		Id             uint   `json:"id"`
		Name           string `json:"name"`
		NameEn         string `json:"nameEn"`
		Subtitle       string `json:"subtitle"`
		Description    string `json:"description"`
		Image          string `json:"image"`
		ImagesStr      string `orm:"images"`
		IngredientsStr string `orm:"ingredients"`
		Content        string `json:"content"`
		CookingTime    int    `json:"cookingTime"`
		Difficulty     int    `json:"difficulty"`
		Servings       int    `json:"servings"`
		ProductIds     string `json:"productIds"`
		Tags           string `json:"tags"`
		ViewCount      int    `json:"viewCount"`
		LikeCount      int    `json:"likeCount"`
		SortOrder      int    `json:"sortOrder"`
		Status         int    `json:"status"`
	}

	var rawDetail RawDetail
	err := dao.Recipe.Ctx(ctx).Where("id", id).Scan(&rawDetail)
	if err != nil {
		return nil, err
	}
	if rawDetail.Id == 0 {
		return nil, gerror.New("食谱不存在")
	}

	detail := &recipe.DetailRes{
		Id:          rawDetail.Id,
		Name:        rawDetail.Name,
		NameEn:      rawDetail.NameEn,
		Subtitle:    rawDetail.Subtitle,
		Description: rawDetail.Description,
		Image:       rawDetail.Image,
		Content:     rawDetail.Content,
		CookingTime: rawDetail.CookingTime,
		Difficulty:  rawDetail.Difficulty,
		Servings:    rawDetail.Servings,
		ProductIds:  rawDetail.ProductIds,
		Tags:        rawDetail.Tags,
		ViewCount:   rawDetail.ViewCount,
		LikeCount:   rawDetail.LikeCount,
		SortOrder:   rawDetail.SortOrder,
		Status:      rawDetail.Status,
	}

	if rawDetail.ImagesStr != "" && rawDetail.ImagesStr != "null" {
		json.Unmarshal([]byte(rawDetail.ImagesStr), &detail.Images)
	}
	if rawDetail.IngredientsStr != "" && rawDetail.IngredientsStr != "null" {
		json.Unmarshal([]byte(rawDetail.IngredientsStr), &detail.Ingredients)
	}

	dao.Recipe.Ctx(ctx).Where("id", id).Increment("view_count", 1)
	return detail, nil
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
