package api

import (
	"gf-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Auth = cAuth{}

type cAuth struct{}

func (c *cAuth) Login(r *ghttp.Request) {
	var req struct {
		Username string `json:"username" v:"required#用户名不能为空"`
		Password string `json:"password" v:"required#密码不能为空"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	token, userId, userName, err := service.Auth().Login(r.Context(), req.Username, req.Password)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data": g.Map{
			"token":    token,
			"userId":   userId,
			"userName": userName,
		},
	})
}

func (c *cAuth) GetInfo(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteJson(g.Map{"code": -1, "message": "未登录"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	g.Log().Infof(r.Context(), "GetInfo token: %s", token)

	claims, err := service.JWT().ParseToken(token)
	if err != nil {
		g.Log().Errorf(r.Context(), "ParseToken error: %v", err)
		r.Response.WriteJson(g.Map{"code": -1, "message": "token无效"})
		return
	}

	userInfo, err := service.Auth().GetUserInfo(r.Context(), claims.UserId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": -1, "message": err.Error()})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "success",
		"data":    userInfo,
	})
}
