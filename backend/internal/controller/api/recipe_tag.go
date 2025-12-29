package api

import (
	"gf-admin/api/v1/recipe"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var RecipeTag = cRecipeTag{}

type cRecipeTag struct{}

func (c *cRecipeTag) List(r *ghttp.Request) {
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

	list, total, err := service.RecipeTag().GetList(r.Context(), req.Keyword, req.Status, req.Page, req.PageSize)
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

func (c *cRecipeTag) All(r *ghttp.Request) {
	list, err := service.RecipeTag().GetAll(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    list,
	})
}

func (c *cRecipeTag) Create(r *ghttp.Request) {
	var req recipe.CreateTagReq

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	id, err := service.RecipeTag().Create(r.Context(), &req)
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

func (c *cRecipeTag) Update(r *ghttp.Request) {
	var req recipe.UpdateTagReq

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	err := service.RecipeTag().Update(r.Context(), &req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}

func (c *cRecipeTag) Delete(r *ghttp.Request) {
	id := r.Get("id").Uint()

	err := service.RecipeTag().Delete(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
	})
}
