package banner

import "github.com/gogf/gf/v2/frame/g"

// GetListReq 获取轮播图列表请求
type GetListReq struct {
	g.Meta   `path:"/banner/list" method:"get" tags:"轮播图" summary:"获取轮播图列表"`
	Position string `json:"position" dc:"位置"`
	Status   *int   `json:"status" dc:"状态"`
	Page     int    `json:"page" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" dc:"每页数量"`
}

type GetListRes struct {
	List  []BannerItem `json:"list"`
	Total int          `json:"total"`
}

type BannerItem struct {
	BannerId   int64  `json:"bannerId"`
	Title      string `json:"title"`
	Description string `json:"description"`
	MediaType  int    `json:"mediaType"`
	MediaUrl   string `json:"mediaUrl"`
	ButtonText string `json:"buttonText"`
	ButtonLink string `json:"buttonLink"`
	SortOrder  int    `json:"sortOrder"`
	Status     int    `json:"status"`
	Position   string `json:"position"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

// CreateReq 创建轮播图请求
type CreateReq struct {
	g.Meta      `path:"/banner/create" method:"post" tags:"轮播图" summary:"创建轮播图"`
	Title       string `json:"title" v:"required#标题不能为空"`
	Description string `json:"description"`
	MediaType   int    `json:"mediaType" v:"required|in:1,2#媒体类型不能为空|媒体类型必须是1或2"`
	MediaUrl    string `json:"mediaUrl" v:"required#媒体URL不能为空"`
	ButtonText  string `json:"buttonText"`
	ButtonLink  string `json:"buttonLink"`
	SortOrder   int    `json:"sortOrder" d:"0"`
	Status      int    `json:"status" d:"1" v:"in:0,1#状态必须是0或1"`
	Position    string `json:"position" d:"home"`
}

type CreateRes struct {
	BannerId int64 `json:"bannerId"`
}

// UpdateReq 更新轮播图请求
type UpdateReq struct {
	g.Meta      `path:"/banner/update" method:"put" tags:"轮播图" summary:"更新轮播图"`
	BannerId    int64  `json:"bannerId" v:"required#轮播图ID不能为空"`
	Title       string `json:"title" v:"required#标题不能为空"`
	Description string `json:"description"`
	MediaType   int    `json:"mediaType" v:"required|in:1,2#媒体类型不能为空|媒体类型必须是1或2"`
	MediaUrl    string `json:"mediaUrl" v:"required#媒体URL不能为空"`
	ButtonText  string `json:"buttonText"`
	ButtonLink  string `json:"buttonLink"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status" v:"in:0,1#状态必须是0或1"`
	Position    string `json:"position"`
}

type UpdateRes struct{}

// DeleteReq 删除轮播图请求
type DeleteReq struct {
	g.Meta   `path:"/banner/delete" method:"delete" tags:"轮播图" summary:"删除轮播图"`
	BannerId int64 `json:"bannerId" v:"required#轮播图ID不能为空"`
}

type DeleteRes struct{}

// GetDetailReq 获取轮播图详情请求
type GetDetailReq struct {
	g.Meta   `path:"/banner/detail" method:"get" tags:"轮播图" summary:"获取轮播图详情"`
	BannerId int64 `json:"bannerId" v:"required#轮播图ID不能为空"`
}

type GetDetailRes struct {
	BannerItem
}
