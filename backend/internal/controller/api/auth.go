// Package api 提供公开 API 接口
// 不需要认证即可访问的接口
package api

import (
	"context"

	"gf-admin/api/v1/auth"
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthController 认证控制器
// 处理用户登录、注册等认证相关操作
var Auth = cAuth{}

type cAuth struct{}

// Login 用户登录接口
// 验证用户名密码，返回 JWT Token
func (c *cAuth) Login(ctx context.Context, req *auth.LoginReq) (res *auth.LoginRes, err error) {
	// 查询用户
	var user struct {
		UserId   int64  `json:"userId"`
		UserName string `json:"userName"`
		Password string `json:"password"`
		Status   int    `json:"status"`
	}

	err = g.DB().Model("sys_user").
		Where("user_name", req.Username).
		Where("del_flag", 0).
		Scan(&user)

	if err != nil {
		return nil, err
	}

	// 检查是否查询到数据
	if user.UserId == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 验证密码
	encryptPassword, _ := gmd5.Encrypt(req.Password)
	if user.Password != encryptPassword {
		return nil, gerror.New("密码错误")
	}

	// 检查用户状态
	if user.Status != 0 {
		return nil, gerror.New("账号已被停用")
	}

	// 生成 Token
	token, err := service.JWT().GenerateToken(user.UserId, user.UserName)
	if err != nil {
		return nil, err
	}

	return &auth.LoginRes{
		Token:    token,
		UserId:   user.UserId,
		UserName: user.UserName,
	}, nil
}

// GetInfo 获取当前登录用户信息
func (c *cAuth) GetInfo(ctx context.Context, req *auth.GetInfoReq) (res *auth.GetInfoRes, err error) {
	// 从上下文获取用户信息
	customCtx := service.Context().Get(ctx)
	if customCtx == nil {
		return nil, gerror.New("获取用户信息失败")
	}

	claims := customCtx.Data["ContextKey"].(*service.Claims)

	// 查询用户详细信息
	var user struct {
		UserId      int64  `json:"userId"`
		UserName    string `json:"userName"`
		NickName    string `json:"nickName"`
		Email       string `json:"email"`
		Phonenumber string `json:"phonenumber"`
		Sex         string `json:"sex"`
		Avatar      string `json:"avatar"`
	}

	err = g.DB().Model("sys_user").
		Where("user_id", claims.UserId).
		Scan(&user)

	if err != nil {
		return nil, err
	}

	return &auth.GetInfoRes{
		UserId:      user.UserId,
		UserName:    user.UserName,
		NickName:    user.NickName,
		Email:       user.Email,
		Phonenumber: user.Phonenumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
	}, nil
}
