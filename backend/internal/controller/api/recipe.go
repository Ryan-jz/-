package api

import (
	"gf-admin/api/v1/recipe"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Recipe 食谱控制器
var Recipe = cRecipe{}

type cRecipe struct{}

// List 获取食谱列表
func (c *cRecipe) List(r *ghttp.Request) {
	var req struct {
		Keyword  string `json:"keyword"`
		Status   *int   `json:"status"`
		Page     int    `json:"page" d:"1"`
		PageSize int    `json:"pageSize" d:"10"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	list, total, err := service.Recipe().GetList(r.Context(), req.Keyword, req.Status, req.Page, req.PageSize)
	if err != nil {
		g.Log().Error(r.Context(), "获取食谱列表失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data": g.Map{
			"list":  list,
			"total": total,
			"page":  req.Page,
		},
	})
}

// Detail 获取食谱详情
func (c *cRecipe) Detail(r *ghttp.Request) {
	id := r.Get("id").Uint()

	detail, err := service.Recipe().GetDetail(r.Context(), id)
	if err != nil {
		g.Log().Error(r.Context(), "获取食谱详情失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    detail,
	})
}

// Create 创建食谱
func (c *cRecipe) Create(r *ghttp.Request) {
	var req recipe.CreateReq

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.Recipe().Create(r.Context(), &req)
	if err != nil {
		g.Log().Error(r.Context(), "创建食谱失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	g.Log().Infof(r.Context(), "创建食谱成功: id=%d", id)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{"id": id},
	})
}

// Update 更新食谱
func (c *cRecipe) Update(r *ghttp.Request) {
	var req recipe.UpdateReq

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Recipe().Update(r.Context(), &req)
	if err != nil {
		g.Log().Error(r.Context(), "更新食谱失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	g.Log().Infof(r.Context(), "更新食谱成功: id=%d", req.Id)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

// Delete 删除食谱
func (c *cRecipe) Delete(r *ghttp.Request) {
	id := r.Get("id").Uint()

	err := service.Recipe().Delete(r.Context(), id)
	if err != nil {
		g.Log().Error(r.Context(), "删除食谱失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	g.Log().Infof(r.Context(), "删除食谱成功: id=%d", id)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}
