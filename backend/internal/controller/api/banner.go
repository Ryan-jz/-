package api

import (
	v1 "gf-admin/api/v1/banner"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Banner = cBanner{}

type cBanner struct{}

func (c *cBanner) GetList(r *ghttp.Request) {
	var req *v1.GetListReq
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

func (c *cBanner) Create(r *ghttp.Request) {
	var req *v1.CreateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	bannerId, err := service.Banner().Create(r.Context(), req)
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

func (c *cBanner) Update(r *ghttp.Request) {
	var req *v1.UpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Banner().Update(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "更新成功",
	})
}

func (c *cBanner) Delete(r *ghttp.Request) {
	var req *v1.DeleteReq
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

func (c *cBanner) GetDetail(r *ghttp.Request) {
	var req *v1.GetDetailReq
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
