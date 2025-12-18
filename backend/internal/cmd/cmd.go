// Package cmd 提供命令行接口
// 负责解析命令行参数并启动相应的服务
package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-admin/internal/controller/admin"
	"gf-admin/internal/controller/api"
	"gf-admin/internal/service"
)

// Main 主命令对象
var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "启动 HTTP 服务器",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化服务
			service.Init(ctx)

			// 创建 HTTP 服务器
			s := g.Server()

			// 配置跨域中间件
			s.Use(MiddlewareCORS)

			// 注册路由组
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 不需要认证的路由
				group.Bind(
					api.Auth,  // 认证相关接口
				)

				// 需要认证的管理后台路由
				group.Group("/admin", func(adminGroup *ghttp.RouterGroup) {
					// JWT 认证中间件
					adminGroup.Middleware(service.Middleware().Auth)
					
					adminGroup.Bind(
						admin.User,     // 用户管理
						admin.Role,     // 角色管理
						admin.Menu,     // 菜单管理
						admin.Dept,     // 部门管理
						admin.Dict,     // 字典管理
						admin.Config,   // 配置管理
						admin.Log,      // 日志管理
					)
				})
			})

			// 启动服务器
			s.Run()
			return nil
		},
	}
)

// MiddlewareCORS 跨域中间件
// 处理跨域请求，允许前端应用访问 API
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
