// Package service 提供业务逻辑服务
package service

import (
	"gf-admin/internal/consts"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// IMiddleware 中间件接口
// 定义系统中间件的标准接口
type IMiddleware interface {
	// Auth JWT认证中间件
	Auth(r *ghttp.Request)
	// Ctx 上下文中间件，用于初始化请求上下文
	Ctx(r *ghttp.Request)
}

type middlewareImpl struct{}

// Middleware 获取中间件服务实例
func Middleware() IMiddleware {
	return &middlewareImpl{}
}

// Ctx 上下文中间件
// 为每个请求初始化上下文信息
func (s *middlewareImpl) Ctx(r *ghttp.Request) {
	// 初始化自定义上下文
	customCtx := &ContextModel{
		Data: make(g.Map),
	}
	Context().Init(r, customCtx)
	r.Middleware.Next()
}

// Auth JWT认证中间件
// 验证请求中的 JWT Token，并将用户信息存入上下文
func (s *middlewareImpl) Auth(r *ghttp.Request) {
	// 获取 Authorization 头
	token := r.Header.Get("Authorization")
	if token == "" {
		response.JsonExit(r, 401, "未授权访问")
	}

	// 移除 "Bearer " 前缀
	if gstr.HasPrefix(token, "Bearer ") {
		token = gstr.SubStr(token, 7)
	}

	// 验证 Token
	claims, err := JWT().ParseToken(token)
	if err != nil {
		response.JsonExit(r, 401, "Token验证失败: "+err.Error())
	}

	// 将用户信息存入上下文
	customCtx := Context().Get(r.Context())
	if customCtx != nil {
		customCtx.Data[consts.ContextKey] = claims
	}

	r.Middleware.Next()
}
