package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BannerEntity 轮播图实体
type BannerEntity struct {
	BannerId    int64       `json:"bannerId"    orm:"banner_id"`
	Title       string      `json:"title"       orm:"title"`
	Description string      `json:"description" orm:"description"`
	MediaType   int         `json:"mediaType"   orm:"media_type"`
	MediaUrl    string      `json:"mediaUrl"    orm:"media_url"`
	ButtonText  string      `json:"buttonText"  orm:"button_text"`
	ButtonLink  string      `json:"buttonLink"  orm:"button_link"`
	SortOrder   int         `json:"sortOrder"   orm:"sort_order"`
	Status      int         `json:"status"      orm:"status"`
	Position    string      `json:"position"    orm:"position"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"`
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"`
}

// IBanner 轮播图服务接口
type IBanner interface {
	GetList(ctx context.Context, position string, status *int, page, pageSize int) (list []*BannerEntity, total int, err error)
	Create(ctx context.Context, data map[string]interface{}) (bannerId int64, err error)
	Update(ctx context.Context, bannerId int64, data map[string]interface{}) error
	Delete(ctx context.Context, bannerId int64) error
	GetDetail(ctx context.Context, bannerId int64) (*BannerEntity, error)
}

type bannerImpl struct{}

var localBanner IBanner

func init() {
	RegisterBanner(&bannerImpl{})
}

// Banner 获取轮播图服务实例
func Banner() IBanner {
	if localBanner == nil {
		panic("implement not found for interface IBanner, forgot register?")
	}
	return localBanner
}

// RegisterBanner 注册轮播图服务实现
func RegisterBanner(i IBanner) {
	localBanner = i
}

// GetList 获取轮播图列表
func (s *bannerImpl) GetList(ctx context.Context, position string, status *int, page, pageSize int) (list []*BannerEntity, total int, err error) {
	model := g.DB().Model("banner")

	// 条件筛选
	if position != "" {
		model = model.Where("position", position)
	}
	if status != nil {
		model = model.Where("status", *status)
	}

	// 获取总数
	totalCount, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = model.
		Order("sort_order ASC, banner_id DESC").
		Page(page, pageSize).
		Scan(&list)

	return list, totalCount, err
}

// Create 创建轮播图
func (s *bannerImpl) Create(ctx context.Context, data map[string]interface{}) (bannerId int64, err error) {
	result, err := g.DB().Model("banner").Data(data).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update 更新轮播图
func (s *bannerImpl) Update(ctx context.Context, bannerId int64, data map[string]interface{}) error {
	// 检查轮播图是否存在
	count, err := g.DB().Model("banner").Where("banner_id", bannerId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("轮播图不存在")
	}

	// 更新
	_, err = g.DB().Model("banner").
		Where("banner_id", bannerId).
		Data(data).
		Update()

	return err
}

// Delete 删除轮播图
func (s *bannerImpl) Delete(ctx context.Context, bannerId int64) error {
	// 检查轮播图是否存在
	count, err := g.DB().Model("banner").Where("banner_id", bannerId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("轮播图不存在")
	}

	// 删除
	_, err = g.DB().Model("banner").Where("banner_id", bannerId).Delete()
	return err
}

// GetDetail 获取轮播图详情
func (s *bannerImpl) GetDetail(ctx context.Context, bannerId int64) (*BannerEntity, error) {
	var result *BannerEntity
	err := g.DB().Model("banner").
		Where("banner_id", bannerId).
		Scan(&result)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, gerror.New("轮播图不存在")
	}

	return result, nil
}
