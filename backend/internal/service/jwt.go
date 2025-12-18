// Package service 提供业务逻辑服务
package service

import (
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
)

// IJwt JWT服务接口
// 提供 Token 的生成和解析功能
type IJwt interface {
	// GenerateToken 生成 JWT Token
	GenerateToken(userId int64, userName string) (string, error)
	// ParseToken 解析 JWT Token
	ParseToken(tokenString string) (*Claims, error)
}

type jwtImpl struct{}

// Claims JWT 声明结构
// 包含用户信息和标准声明
type Claims struct {
	UserId   int64  `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	jwt.RegisteredClaims
}

// JWT 获取 JWT 服务实例
func JWT() IJwt {
	return &jwtImpl{}
}

// GenerateToken 生成 JWT Token
// 根据用户ID和用户名生成加密的 Token 字符串
func (s *jwtImpl) GenerateToken(userId int64, userName string) (string, error) {
	ctx := gctx.New()

	// 获取配置
	signingKey := g.Cfg().MustGet(ctx, "jwt.signingKey").String()
	expireTime := g.Cfg().MustGet(ctx, "jwt.expireTime").Int64()

	// 创建声明
	claims := Claims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireTime) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// ParseToken 解析 JWT Token
// 验证并解析 Token 字符串，返回用户声明信息
func (s *jwtImpl) ParseToken(tokenString string) (*Claims, error) {
	ctx := gctx.New()
	signingKey := g.Cfg().MustGet(ctx, "jwt.signingKey").String()

	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
