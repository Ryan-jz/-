package service

import (
	"context"
	"encoding/json"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type IBanner interface {
	GetList(ctx context.Context, position string, status *int, page, pageSize int) (list []*entity.Banner, total int, err error)
	Create(ctx context.Context, req any) (bannerId int64, err error)
	Update(ctx context.Context, req any) error
	Delete(ctx context.Context, bannerId int64) error
	GetDetail(ctx context.Context, bannerId int64) (*entity.Banner, error)
}

type bannerImpl struct{}

var localBanner IBanner

func init() {
	RegisterBanner(&bannerImpl{})
}

func Banner() IBanner {
	if localBanner == nil {
		panic("implement not found for interface IBanner, forgot register?")
	}
	return localBanner
}

func RegisterBanner(i IBanner) {
	localBanner = i
}

func (s *bannerImpl) GetList(ctx context.Context, position string, status *int, page, pageSize int) (list []*entity.Banner, total int, err error) {
	model := dao.Banner.Ctx(ctx)

	if position != "" {
		model = model.Where(dao.Banner.Columns().Position, position)
	}
	if status != nil {
		model = model.Where(dao.Banner.Columns().Status, *status)
	}

	totalCount, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = model.Order(dao.Banner.Columns().SortOrder + " ASC," + dao.Banner.Columns().BannerId + " DESC").
		Limit(pageSize).
		Offset(offset).
		Scan(&list)

	if err != nil {
		return nil, 0, err
	}

	return list, totalCount, nil
}

func (s *bannerImpl) Create(ctx context.Context, req any) (bannerId int64, err error) {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	result, err := dao.Banner.Ctx(ctx).Data(reqMap).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *bannerImpl) Update(ctx context.Context, req any) error {
	var reqMap g.Map
	jsonBytes, _ := json.Marshal(req)
	json.Unmarshal(jsonBytes, &reqMap)

	bannerId := reqMap["bannerId"]
	delete(reqMap, "bannerId")

	count, err := dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().BannerId, bannerId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("轮播图不存在")
	}

	_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().BannerId, bannerId).Data(reqMap).Update()
	return err
}

func (s *bannerImpl) Delete(ctx context.Context, bannerId int64) error {
	count, err := dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().BannerId, bannerId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("轮播图不存在")
	}

	_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().BannerId, bannerId).Delete()
	return err
}

func (s *bannerImpl) GetDetail(ctx context.Context, bannerId int64) (*entity.Banner, error) {
	var result *entity.Banner
	err := dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().BannerId, bannerId).Scan(&result)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, gerror.New("轮播图不存在")
	}

	return result, nil
}
