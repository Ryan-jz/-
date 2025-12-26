package api

import (
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Banner = cBanner{}

type cBanner struct{}

// GetList 获取轮播图列表
func (c *cBanner) GetList(r *ghttp.Request) {
	var req struct {
		Position string `json:"position"`
		Status   *int   `json:"status"`
		Page     int    `json:"page" d:"1"`
		PageSize int    `json:"pageSize" d:"10"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	list, total, err := service.Banner().GetList(r.Context(), req.Position, req.Status, req.Page, req.PageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data": g.Map{
			"list":  list,
			"total": total,
		},
	})
}

// Create 创建轮播图
func (c *cBanner) Create(r *ghttp.Request) {
	var req struct {
		Title       string `json:"title" v:"required#标题不能为空"`
		Description string `json:"description"`
		MediaType   int    `json:"mediaType" v:"required|in:1,2#媒体类型不能为空|媒体类型必须是1或2"`
		MediaUrl    string `json:"mediaUrl" v:"required#媒体URL不能为空"`
		ButtonText  string `json:"buttonText"`
		ButtonLink  string `json:"buttonLink"`
		SortOrder   int    `json:"sortOrder" d:"0"`
		Status      int    `json:"status" d:"1"`
		Position    string `json:"position" d:"home"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	data := g.Map{
		"title":       req.Title,
		"description": req.Description,
		"media_type":  req.MediaType,
		"media_url":   req.MediaUrl,
		"button_text": req.ButtonText,
		"button_link": req.ButtonLink,
		"sort_order":  req.SortOrder,
		"status":      req.Status,
		"position":    req.Position,
	}

	bannerId, err := service.Banner().Create(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "创建成功",
		"data":    g.Map{"bannerId": bannerId},
	})
}

// Update 更新轮播图
func (c *cBanner) Update(r *ghttp.Request) {
	var req struct {
		BannerId    int64  `json:"bannerId" v:"required#轮播图ID不能为空"`
		Title       string `json:"title" v:"required#标题不能为空"`
		Description string `json:"description"`
		MediaType   int    `json:"mediaType" v:"required|in:1,2#媒体类型不能为空|媒体类型必须是1或2"`
		MediaUrl    string `json:"mediaUrl" v:"required#媒体URL不能为空"`
		ButtonText  string `json:"buttonText"`
		ButtonLink  string `json:"buttonLink"`
		SortOrder   int    `json:"sortOrder"`
		Status      int    `json:"status"`
		Position    string `json:"position"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	data := g.Map{
		"title":       req.Title,
		"description": req.Description,
		"media_type":  req.MediaType,
		"media_url":   req.MediaUrl,
		"button_text": req.ButtonText,
		"button_link": req.ButtonLink,
		"sort_order":  req.SortOrder,
		"status":      req.Status,
		"position":    req.Position,
	}

	err := service.Banner().Update(r.Context(), req.BannerId, data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "更新成功",
	})
}

// Delete 删除轮播图
func (c *cBanner) Delete(r *ghttp.Request) {
	var req struct {
		BannerId int64 `json:"bannerId" v:"required#轮播图ID不能为空"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Banner().Delete(r.Context(), req.BannerId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "删除成功",
	})
}

// GetDetail 获取轮播图详情
func (c *cBanner) GetDetail(r *ghttp.Request) {
	var req struct {
		BannerId int64 `json:"bannerId" v:"required#轮播图ID不能为空"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	detail, err := service.Banner().GetDetail(r.Context(), req.BannerId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    detail,
	})
}
