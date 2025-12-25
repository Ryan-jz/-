package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// IAuth 认证服务接口
type IAuth interface {
	Login(ctx context.Context, username, password string) (token string, userId int64, userName string, err error)
	GetUserInfo(ctx context.Context, userId int64) (map[string]interface{}, error)
}

type authImpl struct{}

var localAuth IAuth

func init() {
	RegisterAuth(&authImpl{})
}

// Auth 获取认证服务实例
func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

// RegisterAuth 注册认证服务实现
func RegisterAuth(i IAuth) {
	localAuth = i
}

// Login 用户登录
func (s *authImpl) Login(ctx context.Context, username, password string) (token string, userId int64, userName string, err error) {
	// 查询用户
	var user struct {
		UserId   int64  `json:"userId"`
		UserName string `json:"userName"`
		Password string `json:"password"`
		Status   int    `json:"status"`
	}

	g.Log().Infof(ctx, "查询用户: username=%s", username)

	err = g.DB().Model("sys_user").
		Where("user_name", username).
		Where("del_flag", 0).
		Scan(&user)

	if err != nil {
		g.Log().Errorf(ctx, "数据库查询失败: %v", err)
		return "", 0, "", err
	}

	// 检查是否查询到数据
	if user.UserId == 0 {
		g.Log().Warningf(ctx, "用户不存在: username=%s", username)
		return "", 0, "", gerror.New("用户不存在")
	}

	// 验证密码（这里简化处理，实际应该使用加密）
	if user.Password != password {
		g.Log().Warningf(ctx, "密码错误: username=%s", username)
		g.Log().Infof(ctx, "密码错误: usernamePwd=%s, password=%s",  user.Password, password)
		return "", 0, "", gerror.New("密码错误")
	}

	// 检查用户状态
	if user.Status != 0 {
		g.Log().Warningf(ctx, "账号已被停用: username=%s", username)
		return "", 0, "", gerror.New("账号已被停用")
	}

	// 生成 Token
	token, err = JWT().GenerateToken(user.UserId, user.UserName)
	if err != nil {
		g.Log().Errorf(ctx, "生成Token失败: %v", err)
		return "", 0, "", err
	}

	return token, user.UserId, user.UserName, nil
}

// GetUserInfo 获取用户信息
func (s *authImpl) GetUserInfo(ctx context.Context, userId int64) (map[string]interface{}, error) {
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

	err := g.DB().Model("sys_user").
		Where("user_id", userId).
		Scan(&user)

	if err != nil {
		return nil, err
	}

	if user.UserId == 0 {
		return nil, gerror.New("用户不存在")
	}

	return g.Map{
		"userId":      user.UserId,
		"userName":    user.UserName,
		"nickName":    user.NickName,
		"email":       user.Email,
		"phonenumber": user.Phonenumber,
		"sex":         user.Sex,
		"avatar":      user.Avatar,
	}, nil
}
