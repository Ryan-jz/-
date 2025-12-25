// Package router 路由注册
package router

import (
	"gf-admin/internal/controller/api"

	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(s *ghttp.Server) {
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 认证相关路由
		group.POST("/login", api.Auth.Login)
		group.GET("/getInfo", api.Auth.GetInfo)

		// 产品相关路由
		group.GET("/product/category/list", api.Product.CategoryList)
		group.GET("/product/list", api.Product.List)
		group.GET("/product/detail/:id", api.Product.Detail)
		group.POST("/product/create", api.Product.Create)
		group.PUT("/product/update", api.Product.Update)
		group.DELETE("/product/delete/:id", api.Product.Delete)
	})
}
