package product

import "github.com/gogf/gf/v2/frame/g"

type CategoryListReq struct {
	g.Meta `path:"/product/category/list" method:"get" tags:"产品" summary:"获取产品分类列表"`
	Status *int `json:"status" dc:"状态：1启用 0禁用"`
}

type CategoryListRes struct {
	List []CategoryItem `json:"list"`
}

type CategoryItem struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CategoryCreateReq struct {
	g.Meta      `path:"/product/category/create" method:"post" tags:"产品" summary:"创建产品分类"`
	Name        string `json:"name" v:"required#分类名称不能为空"`
	Slug        string `json:"slug" v:"required#分类标识不能为空"`
	Description string `json:"description"`
	Image       string `json:"image"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status" d:"1"`
}

type CategoryCreateRes struct {
	Id uint `json:"id"`
}

type CategoryUpdateReq struct {
	g.Meta      `path:"/product/category/update" method:"put" tags:"产品" summary:"更新产品分类"`
	Id          uint   `json:"id" v:"required#分类ID不能为空"`
	Name        string `json:"name" v:"required#分类名称不能为空"`
	Slug        string `json:"slug" v:"required#分类标识不能为空"`
	Description string `json:"description"`
	Image       string `json:"image"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
}

type CategoryUpdateRes struct{}

type CategoryDeleteReq struct {
	g.Meta `path:"/product/category/delete/:id" method:"delete" tags:"产品" summary:"删除产品分类"`
	Id     uint `json:"id" v:"required#分类ID不能为空"`
}

type CategoryDeleteRes struct{}

type CategoryWithProductsReq struct {
	g.Meta `path:"/product/category/with-products" method:"get" tags:"产品" summary:"获取分类及产品列表"`
	Status *int `json:"status" dc:"状态：1启用 0禁用"`
}

type CategoryWithProductsRes struct {
	List []CategoryWithProducts `json:"list"`
}

type CategoryWithProducts struct {
	Id          uint          `json:"id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	SortOrder   int           `json:"sortOrder"`
	Status      int           `json:"status"`
	Products    []ProductItem `json:"products"`
}

type ProductItem struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Subtitle string  `json:"subtitle"`
	Image    string  `json:"image"`
	Price    float64 `json:"price"`
	Weight   string  `json:"weight"`
	Status   int     `json:"status"`
}

type ListReq struct {
	g.Meta     `path:"/product/list" method:"get" tags:"产品" summary:"获取产品列表"`
	CategoryId *uint  `json:"categoryId" dc:"分类ID"`
	Keyword    string `json:"keyword" dc:"关键词搜索"`
	Status     *int   `json:"status" dc:"状态：1上架 0下架"`
	Page       int    `json:"page" d:"1" dc:"页码"`
	PageSize   int    `json:"pageSize" d:"10" dc:"每页数量"`
}

type ListRes struct {
	List  []ListItem `json:"list"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
}

type ListItem struct {
	Id           uint    `json:"id"`
	CategoryId   uint    `json:"categoryId"`
	CategoryIds  string  `json:"categoryIds"`
	Name         string  `json:"name"`
	Subtitle     string  `json:"subtitle"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Weight       string  `json:"weight"`
	PurchaseLink string  `json:"purchaseLink"`
	SortOrder    int     `json:"sortOrder"`
	Status       int     `json:"status"`
	ViewCount    int     `json:"viewCount"`
}

type DetailReq struct {
	g.Meta `path:"/product/detail/:id" method:"get" tags:"产品" summary:"获取产品详情"`
	Id     uint `json:"id" v:"required#产品ID不能为空"`
}

type DetailRes struct {
	Id            uint     `json:"id"`
	CategoryId    uint     `json:"categoryId"`
	CategoryIds   string   `json:"categoryIds"`
	Name          string   `json:"name"`
	Subtitle      string   `json:"subtitle"`
	Description   string   `json:"description"`
	Image         string   `json:"image"`
	Images        []string `json:"images"`
	Price         float64  `json:"price"`
	Stock         int      `json:"stock"`
	Weight        string   `json:"weight"`
	Ingredients   string   `json:"ingredients"`
	Nutrition     string   `json:"nutrition"`
	Usage         string   `json:"usage"`
	Features      []string `json:"features"`
	OrganicCert   string   `json:"organicCert"`
	RecyclingInfo string   `json:"recyclingInfo"`
	PurchaseLink  string   `json:"purchaseLink"`
	AllergenInfo  string   `json:"allergenInfo"`
	StorageInfo   string   `json:"storageInfo"`
	Origin        string   `json:"origin"`
	SortOrder     int      `json:"sortOrder"`
	Status        int      `json:"status"`
	ViewCount     int      `json:"viewCount"`
}

type CreateReq struct {
	g.Meta        `path:"/product/create" method:"post" tags:"产品" summary:"创建产品"`
	CategoryId    uint     `json:"categoryId" v:"required#分类ID不能为空"`
	CategoryIds   string   `json:"categoryIds"`
	Name          string   `json:"name" v:"required#产品名称不能为空"`
	Subtitle      string   `json:"subtitle"`
	Description   string   `json:"description"`
	Image         string   `json:"image" v:"required#产品图片不能为空"`
	Images        []string `json:"images"`
	Price         float64  `json:"price"`
	Stock         int      `json:"stock"`
	Weight        string   `json:"weight"`
	Ingredients   string   `json:"ingredients"`
	Nutrition     string   `json:"nutrition"`
	Usage         string   `json:"usage"`
	Features      []string `json:"features"`
	OrganicCert   string   `json:"organicCert"`
	RecyclingInfo string   `json:"recyclingInfo"`
	PurchaseLink  string   `json:"purchaseLink"`
	AllergenInfo  string   `json:"allergenInfo"`
	StorageInfo   string   `json:"storageInfo"`
	Origin        string   `json:"origin"`
	SortOrder     int      `json:"sortOrder"`
	Status        int      `json:"status" d:"1"`
}

type CreateRes struct {
	Id uint `json:"id"`
}

type UpdateReq struct {
	g.Meta        `path:"/product/update" method:"put" tags:"产品" summary:"更新产品"`
	Id            uint     `json:"id" v:"required#产品ID不能为空"`
	CategoryId    uint     `json:"categoryId" v:"required#分类ID不能为空"`
	CategoryIds   string   `json:"categoryIds"`
	Name          string   `json:"name" v:"required#产品名称不能为空"`
	Subtitle      string   `json:"subtitle"`
	Description   string   `json:"description"`
	Image         string   `json:"image" v:"required#产品图片不能为空"`
	Images        []string `json:"images"`
	Price         float64  `json:"price"`
	Stock         int      `json:"stock"`
	Weight        string   `json:"weight"`
	Ingredients   string   `json:"ingredients"`
	Nutrition     string   `json:"nutrition"`
	Usage         string   `json:"usage"`
	Features      []string `json:"features"`
	OrganicCert   string   `json:"organicCert"`
	RecyclingInfo string   `json:"recyclingInfo"`
	PurchaseLink  string   `json:"purchaseLink"`
	AllergenInfo  string   `json:"allergenInfo"`
	StorageInfo   string   `json:"storageInfo"`
	Origin        string   `json:"origin"`
	SortOrder     int      `json:"sortOrder"`
	Status        int      `json:"status"`
}

type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/product/delete/:id" method:"delete" tags:"产品" summary:"删除产品"`
	Id     uint `json:"id" v:"required#产品ID不能为空"`
}

type DeleteRes struct{}
