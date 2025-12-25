// Package cmd 提供命令行接口
// 负责解析命令行参数并启动相应的服务
package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-admin/internal/router"
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

			// 注册路由
			router.RegisterRoutes(s)

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
