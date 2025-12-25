package product

import "github.com/gogf/gf/v2/frame/g"

// CategoryListReq 产品分类列表请求
type CategoryListReq struct {
	g.Meta `path:"/product/category/list" method:"get" tags:"产品" summary:"获取产品分类列表"`
	Status *int `json:"status" dc:"状态：1启用 0禁用"`
}

// CategoryListRes 产品分类列表响应
type CategoryListRes struct {
	List []CategoryItem `json:"list"`
}

// CategoryItem 产品分类项
type CategoryItem struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	NameEn      string `json:"nameEn"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
}

// ListReq 产品列表请求
type ListReq struct {
	g.Meta     `path:"/product/list" method:"get" tags:"产品" summary:"获取产品列表"`
	CategoryId *uint  `json:"categoryId" dc:"分类ID"`
	Keyword    string `json:"keyword" dc:"关键词搜索"`
	Status     *int   `json:"status" dc:"状态：1上架 0下架"`
	Page       int    `json:"page" d:"1" dc:"页码"`
	PageSize   int    `json:"pageSize" d:"10" dc:"每页数量"`
}

// ListRes 产品列表响应
type ListRes struct {
	List  []ListItem `json:"list"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
}

// ListItem 产品项
type ListItem struct {
	Id          uint    `json:"id"`
	CategoryId  uint    `json:"categoryId"`
	Name        string  `json:"name"`
	NameEn      string  `json:"nameEn"`
	Subtitle    string  `json:"subtitle"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Weight      string  `json:"weight"`
	SortOrder   int     `json:"sortOrder"`
	Status      int     `json:"status"`
	ViewCount   int     `json:"viewCount"`
}

// DetailReq 产品详情请求
type DetailReq struct {
	g.Meta `path:"/product/detail/:id" method:"get" tags:"产品" summary:"获取产品详情"`
	Id     uint `json:"id" v:"required#产品ID不能为空"`
}

// DetailRes 产品详情响应
type DetailRes struct {
	Id          uint     `json:"id"`
	CategoryId  uint     `json:"categoryId"`
	Name        string   `json:"name"`
	NameEn      string   `json:"nameEn"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Weight      string   `json:"weight"`
	Ingredients string   `json:"ingredients"`
	Nutrition   string   `json:"nutrition"`
	Usage       string   `json:"usage"`
	Features    []string `json:"features"`
	SortOrder   int      `json:"sortOrder"`
	Status      int      `json:"status"`
	ViewCount   int      `json:"viewCount"`
}

// CreateReq 创建产品请求
type CreateReq struct {
	g.Meta      `path:"/product/create" method:"post" tags:"产品" summary:"创建产品"`
	CategoryId  uint     `json:"categoryId" v:"required#分类ID不能为空"`
	Name        string   `json:"name" v:"required#产品名称不能为空"`
	NameEn      string   `json:"nameEn"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Weight      string   `json:"weight"`
	Ingredients string   `json:"ingredients"`
	Nutrition   string   `json:"nutrition"`
	Usage       string   `json:"usage"`
	Features    []string `json:"features"`
	SortOrder   int      `json:"sortOrder"`
	Status      int      `json:"status" d:"1"`
}

// CreateRes 创建产品响应
type CreateRes struct {
	Id uint `json:"id"`
}

// UpdateReq 更新产品请求
type UpdateReq struct {
	g.Meta      `path:"/product/update" method:"put" tags:"产品" summary:"更新产品"`
	Id          uint     `json:"id" v:"required#产品ID不能为空"`
	CategoryId  uint     `json:"categoryId"`
	Name        string   `json:"name"`
	NameEn      string   `json:"nameEn"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Weight      string   `json:"weight"`
	Ingredients string   `json:"ingredients"`
	Nutrition   string   `json:"nutrition"`
	Usage       string   `json:"usage"`
	Features    []string `json:"features"`
	SortOrder   int      `json:"sortOrder"`
	Status      int      `json:"status"`
}

// UpdateRes 更新产品响应
type UpdateRes struct{}

// DeleteReq 删除产品请求
type DeleteReq struct {
	g.Meta `path:"/product/delete/:id" method:"delete" tags:"产品" summary:"删除产品"`
	Id     uint `json:"id" v:"required#产品ID不能为空"`
}

// DeleteRes 删除产品响应
type DeleteRes struct{}
