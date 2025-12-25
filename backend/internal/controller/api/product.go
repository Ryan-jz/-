package api

import (
	"gf-admin/api/v1/product"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Product 产品控制器
var Product = cProduct{}

type cProduct struct{}

// CategoryList 获取产品分类列表
func (c *cProduct) CategoryList(r *ghttp.Request) {
	var req struct {
		Status *int `json:"status"`
	}
	
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

// CategoryCreate 创建产品分类
func (c *cProduct) CategoryCreate(r *ghttp.Request) {
	var req struct {
		Name        string `json:"name" v:"required#分类名称不能为空"`
		NameEn      string `json:"nameEn"`
		Slug        string `json:"slug" v:"required#分类标识不能为空"`
		Description string `json:"description"`
		SortOrder   int    `json:"sortOrder"`
		Status      int    `json:"status" d:"1"`
	}
	
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.Product().CreateCategory(r.Context(), &req)
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

// CategoryUpdate 更新产品分类
func (c *cProduct) CategoryUpdate(r *ghttp.Request) {
	var req struct {
		Id          uint   `json:"id" v:"required#分类ID不能为空"`
		Name        string `json:"name" v:"required#分类名称不能为空"`
		NameEn      string `json:"nameEn"`
		Slug        string `json:"slug" v:"required#分类标识不能为空"`
		Description string `json:"description"`
		SortOrder   int    `json:"sortOrder"`
		Status      int    `json:"status"`
	}
	
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().UpdateCategory(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

// CategoryDelete 删除产品分类
func (c *cProduct) CategoryDelete(r *ghttp.Request) {
	id := r.Get("id").Uint()

	err := service.Product().DeleteCategory(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

// List 获取产品列表
func (c *cProduct) List(r *ghttp.Request) {
	var req struct {
		CategoryId *uint  `json:"categoryId"`
		Keyword    string `json:"keyword"`
		Status     *int   `json:"status"`
		Page       int    `json:"page" d:"1"`
		PageSize   int    `json:"pageSize" d:"10"`
	}
	
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

// Detail 获取产品详情
func (c *cProduct) Detail(r *ghttp.Request) {
	id := r.Get("id").Uint()

	detail, err := service.Product().GetDetail(r.Context(), id)
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

// Create 创建产品
func (c *cProduct) Create(r *ghttp.Request) {
	var req product.CreateReq
	
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.Product().Create(r.Context(), &req)
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

// Update 更新产品
func (c *cProduct) Update(r *ghttp.Request) {
	var req product.UpdateReq
	
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.Product().Update(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{},
	})
}

// Delete 删除产品
func (c *cProduct) Delete(r *ghttp.Request) {
	id := r.Get("id").Uint()

	err := service.Product().Delete(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    g.Map{},
	})
}
