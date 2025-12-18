// Package admin 定义管理后台 API 接口结构
package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserListReq 用户列表请求
type UserListReq struct {
	g.Meta      `path:"/user/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
	UserName    string `json:"userName" dc:"用户名"`
	Phonenumber string `json:"phonenumber" dc:"手机号"`
	Status      int    `json:"status" dc:"状态"`
	Page        int    `json:"page" dc:"页码"`
	PageSize    int    `json:"pageSize" dc:"每页数量"`
}

// UserItem 用户列表项
type UserItem struct {
	UserId      int64       `json:"userId" dc:"用户ID"`
	UserName    string      `json:"userName" dc:"用户名"`
	NickName    string      `json:"nickName" dc:"昵称"`
	Email       string      `json:"email" dc:"邮箱"`
	Phonenumber string      `json:"phonenumber" dc:"手机号"`
	Sex         string      `json:"sex" dc:"性别"`
	Avatar      string      `json:"avatar" dc:"头像"`
	Status      int         `json:"status" dc:"状态"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}

// UserListRes 用户列表响应
type UserListRes struct {
	List     []UserItem `json:"list" dc:"用户列表"`
	Total    int        `json:"total" dc:"总数"`
	Page     int        `json:"page" dc:"页码"`
	PageSize int        `json:"pageSize" dc:"每页数量"`
}

// UserDetailReq 用户详情请求
type UserDetailReq struct {
	g.Meta `path:"/user/detail" method:"get" tags:"用户管理" summary:"获取用户详情"`
	UserId int64 `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
}

// UserDetailRes 用户详情响应
type UserDetailRes struct {
	User UserItem `json:"user" dc:"用户信息"`
}

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	g.Meta      `path:"/user/create" method:"post" tags:"用户管理" summary:"创建用户"`
	UserName    string `json:"userName" v:"required#用户名不能为空" dc:"用户名"`
	NickName    string `json:"nickName" v:"required#昵称不能为空" dc:"昵称"`
	Password    string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Email       string `json:"email" dc:"邮箱"`
	Phonenumber string `json:"phonenumber" dc:"手机号"`
	Sex         string `json:"sex" dc:"性别"`
	Status      int    `json:"status" dc:"状态"`
	DeptId      int64  `json:"deptId" dc:"部门ID"`
	Remark      string `json:"remark" dc:"备注"`
}

// UserCreateRes 创建用户响应
type UserCreateRes struct {
	UserId int64 `json:"userId" dc:"用户ID"`
}

// UserUpdateReq 更新用户请求
type UserUpdateReq struct {
	g.Meta      `path:"/user/update" method:"put" tags:"用户管理" summary:"更新用户"`
	UserId      int64  `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
	NickName    string `json:"nickName" dc:"昵称"`
	Email       string `json:"email" dc:"邮箱"`
	Phonenumber string `json:"phonenumber" dc:"手机号"`
	Sex         string `json:"sex" dc:"性别"`
	Status      int    `json:"status" dc:"状态"`
	DeptId      int64  `json:"deptId" dc:"部门ID"`
	Remark      string `json:"remark" dc:"备注"`
}

// UserUpdateRes 更新用户响应
type UserUpdateRes struct{}

// UserDeleteReq 删除用户请求
type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"用户管理" summary:"删除用户"`
	UserId int64 `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
}

// UserDeleteRes 删除用户响应
type UserDeleteRes struct{}
