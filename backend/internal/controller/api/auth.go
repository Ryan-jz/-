// Package api 提供公开 API 接口
// 不需要认证即可访问的接口
package api

import (
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AuthController 认证控制器
// 处理用户登录、注册等认证相关操作
var Auth = cAuth{}

type cAuth struct{}

// Login 用户登录接口
func (c *cAuth) Login(r *ghttp.Request) {
	var req struct {
		Username string `json:"username" v:"required#用户名不能为空"`
		Password string `json:"password" v:"required#密码不能为空"`
	}
	
	if err := r.Parse(&req); err != nil {
		g.Log().Error(r.Context(), "登录请求解析失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	g.Log().Infof(r.Context(), "用户登录请求: username=%s", req.Username)

	token, userId, userName, err := service.Auth().Login(r.Context(), req.Username, req.Password)
	if err != nil {
		g.Log().Error(r.Context(), "登录失败:", err)
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	g.Log().Infof(r.Context(), "用户登录成功: userId=%d, userName=%s", userId, userName)
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data": g.Map{
			"token":    token,
			"userId":   userId,
			"userName": userName,
		},
	})
}

// GetInfo 获取当前登录用户信息
func (c *cAuth) GetInfo(r *ghttp.Request) {
	// 从上下文获取用户信息
	customCtx := service.Context().Get(r.Context())
	if customCtx == nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": "获取用户信息失败"})
		return
	}

	claims := customCtx.Data["ContextKey"].(*service.Claims)

	userInfo, err := service.Auth().GetUserInfo(r.Context(), claims.UserId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    userInfo,
	})
}
