package api

import (
	v1 "gf-admin/api/v1/product"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Product = cProduct{}

type cProduct struct{}

func (c *cProduct) CategoryList(r *ghttp.Request) {
	var req *v1.CategoryListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	list, err := service.Product().GetCategoryList(r.Context(), req.Status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{"list": list},
	})
}

func (c *cProduct) CategoryCreate(r *ghttp.Request) {
	var req *v1.CategoryCreateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.Product().CreateCategory(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{"id": id},
	})
}

func (c *cProduct) CategoryUpdate(r *ghttp.Request) {
	var req *v1.CategoryUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().UpdateCategory(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

func (c *cProduct) CategoryDelete(r *ghttp.Request) {
	var req *v1.CategoryDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().DeleteCategory(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

func (c *cProduct) CategoryWithProducts(r *ghttp.Request) {
	var req *v1.CategoryWithProductsReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	list, err := service.Product().GetCategoryWithProducts(r.Context(), req.Status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{"list": list},
	})
}

func (c *cProduct) List(r *ghttp.Request) {
	var req *v1.ListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	list, total, err := service.Product().GetList(r.Context(), req.CategoryId, req.Keyword, req.Status, req.Page, req.PageSize)
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
			"page":  req.Page,
		},
	})
}

func (c *cProduct) Detail(r *ghttp.Request) {
	var req *v1.DetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	detail, err := service.Product().GetDetail(r.Context(), req.Id)
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

func (c *cProduct) Create(r *ghttp.Request) {
	var req *v1.CreateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.Product().Create(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{"id": id},
	})
}

func (c *cProduct) Update(r *ghttp.Request) {
	var req *v1.UpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().Update(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

func (c *cProduct) Delete(r *ghttp.Request) {
	var req *v1.DeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().Delete(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}
