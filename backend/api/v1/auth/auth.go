// Package auth 定义认证相关的 API 接口结构
package auth

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 登录请求参数
type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"认证" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

// LoginRes 登录响应数据
type LoginRes struct {
	Token    string `json:"token" dc:"访问令牌"`
	UserId   int64  `json:"userId" dc:"用户ID"`
	UserName string `json:"userName" dc:"用户名"`
}

// GetInfoReq 获取用户信息请求
type GetInfoReq struct {
	g.Meta `path:"/getInfo" method:"get" tags:"认证" summary:"获取用户信息"`
}

// GetInfoRes 获取用户信息响应
type GetInfoRes struct {
	UserId      int64  `json:"userId" dc:"用户ID"`
	UserName    string `json:"userName" dc:"用户名"`
	NickName    string `json:"nickName" dc:"昵称"`
	Email       string `json:"email" dc:"邮箱"`
	Phonenumber string `json:"phonenumber" dc:"手机号"`
	Sex         string `json:"sex" dc:"性别"`
	Avatar      string `json:"avatar" dc:"头像"`
}
