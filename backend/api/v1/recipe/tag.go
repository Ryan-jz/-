package recipe

import "time"

type TagListItem struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	SortOrder int       `json:"sortOrder"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTagReq struct {
	Name      string `json:"name" v:"required#标签名称不能为空"`
	Slug      string `json:"slug" v:"required#标签标识不能为空"`
	SortOrder int    `json:"sortOrder"`
	Status    int    `json:"status" d:"1"`
}

type UpdateTagReq struct {
	Id        uint   `json:"id" v:"required#标签ID不能为空"`
	Name      string `json:"name" v:"required#标签名称不能为空"`
	Slug      string `json:"slug" v:"required#标签标识不能为空"`
	SortOrder int    `json:"sortOrder"`
	Status    int    `json:"status"`
}
