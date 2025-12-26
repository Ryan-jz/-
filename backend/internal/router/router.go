// Package router 路由注册
package router

import (
	"gf-admin/internal/controller/api"

	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(s *ghttp.Server) {
	// 配置静态文件服务 - 用于访问上传的文件
	s.AddStaticPath("/uploads", "public/uploads")

	// 兼容旧版 API 路径
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.GET("/getInfo", api.Auth.GetInfo)
	})

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 认证相关路由
		group.POST("/login", api.Auth.Login)
		group.GET("/getInfo", api.Auth.GetInfo)

		// 文件上传路由
		group.POST("/upload/image", api.Upload.Image)
		group.POST("/upload/video", api.Upload.Video)
		group.POST("/upload/delete", api.Upload.Delete)

		// 产品相关路由
		group.GET("/product/category/list", api.Product.CategoryList)
		group.POST("/product/category/create", api.Product.CategoryCreate)
		group.PUT("/product/category/update", api.Product.CategoryUpdate)
		group.DELETE("/product/category/delete/:id", api.Product.CategoryDelete)
		
		group.GET("/product/list", api.Product.List)
		group.GET("/product/detail/:id", api.Product.Detail)
		group.POST("/product/create", api.Product.Create)
		group.PUT("/product/update", api.Product.Update)
		group.DELETE("/product/delete/:id", api.Product.Delete)

		// 食谱相关路由
		group.GET("/recipe/list", api.Recipe.List)
		group.GET("/recipe/detail/:id", api.Recipe.Detail)
		group.POST("/recipe/create", api.Recipe.Create)
		group.PUT("/recipe/update", api.Recipe.Update)
		group.DELETE("/recipe/delete/:id", api.Recipe.Delete)

		// 轮播图相关路由
		group.GET("/banner/list", api.Banner.GetList)
		group.GET("/banner/detail", api.Banner.GetDetail)
		group.POST("/banner/create", api.Banner.Create)
		group.PUT("/banner/update", api.Banner.Update)
		group.DELETE("/banner/delete", api.Banner.Delete)
	})
}
