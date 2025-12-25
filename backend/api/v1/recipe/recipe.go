package recipe

import "time"

// ListItem 食谱列表项
type ListItem struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	NameEn      string    `json:"nameEn"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CookingTime int       `json:"cookingTime"`
	Difficulty  int       `json:"difficulty"`
	Servings    int       `json:"servings"`
	ProductIds  string    `json:"productIds"`
	Tags        string    `json:"tags"`
	ViewCount   int       `json:"viewCount"`
	LikeCount   int       `json:"likeCount"`
	SortOrder   int       `json:"sortOrder"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

// DetailRes 食谱详情响应
type DetailRes struct {
	Id          uint                   `json:"id"`
	Name        string                 `json:"name"`
	NameEn      string                 `json:"nameEn"`
	Subtitle    string                 `json:"subtitle"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Images      []string               `json:"images"`
	Ingredients []map[string]string    `json:"ingredients"`
	Content     string                 `json:"content"`
	CookingTime int                    `json:"cookingTime"`
	Difficulty  int                    `json:"difficulty"`
	Servings    int                    `json:"servings"`
	ProductIds  string                 `json:"productIds"`
	Tags        string                 `json:"tags"`
	ViewCount   int                    `json:"viewCount"`
	LikeCount   int                    `json:"likeCount"`
	SortOrder   int                    `json:"sortOrder"`
	Status      int                    `json:"status"`
	CreateTime  time.Time              `json:"createTime"`
	UpdateTime  time.Time              `json:"updateTime"`
}

// CreateReq 创建食谱请求
type CreateReq struct {
	Name        string                 `json:"name" v:"required#食谱名称不能为空"`
	NameEn      string                 `json:"nameEn"`
	Subtitle    string                 `json:"subtitle"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Images      []string               `json:"images"`
	Ingredients []map[string]string    `json:"ingredients"`
	Content     string                 `json:"content"`
	CookingTime int                    `json:"cookingTime"`
	Difficulty  int                    `json:"difficulty" d:"1"`
	Servings    int                    `json:"servings" d:"1"`
	ProductIds  string                 `json:"productIds"`
	Tags        string                 `json:"tags"`
	SortOrder   int                    `json:"sortOrder"`
	Status      int                    `json:"status" d:"1"`
}

// UpdateReq 更新食谱请求
type UpdateReq struct {
	Id          uint                   `json:"id" v:"required#食谱ID不能为空"`
	Name        string                 `json:"name" v:"required#食谱名称不能为空"`
	NameEn      string                 `json:"nameEn"`
	Subtitle    string                 `json:"subtitle"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Images      []string               `json:"images"`
	Ingredients []map[string]string    `json:"ingredients"`
	Content     string                 `json:"content"`
	CookingTime int                    `json:"cookingTime"`
	Difficulty  int                    `json:"difficulty"`
	Servings    int                    `json:"servings"`
	ProductIds  string                 `json:"productIds"`
	Tags        string                 `json:"tags"`
	SortOrder   int                    `json:"sortOrder"`
	Status      int                    `json:"status"`
}
