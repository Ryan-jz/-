// Package service 提供业务逻辑服务
package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ContextModel 自定义上下文模型
// 用于在请求生命周期中传递数据
type ContextModel struct {
	Data g.Map // 自定义数据存储
}

// IContext 上下文服务接口
type IContext interface {
	// Init 初始化上下文
	Init(r *ghttp.Request, customCtx *ContextModel)
	// Get 获取上下文
	Get(ctx context.Context) *ContextModel
	// SetUser 设置用户信息到上下文
	SetUser(ctx context.Context, user *ContextModel)
}

type contextImpl struct{}

const (
	// contextKey 上下文键名
	contextKey = "ContextModel"
)

// Context 获取上下文服务实例
func Context() IContext {
	return &contextImpl{}
}

// Init 初始化上下文
// 将自定义上下文绑定到请求中
func (s *contextImpl) Init(r *ghttp.Request, customCtx *ContextModel) {
	r.SetCtxVar(contextKey, customCtx)
}

// Get 获取上下文
// 从请求中获取自定义上下文
func (s *contextImpl) Get(ctx context.Context) *ContextModel {
	value := ctx.Value(contextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*ContextModel); ok {
		return localCtx
	}
	return nil
}

// SetUser 设置用户信息到上下文
func (s *contextImpl) SetUser(ctx context.Context, ctxModel *ContextModel) {
	// 实现用户信息设置逻辑
}
