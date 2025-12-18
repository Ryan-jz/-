// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// ConfigController 配置管理控制器
var Config = cConfig{}

type cConfig struct{}

// List 获取配置列表
func (c *cConfig) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现配置列表查询
	return nil, nil
}
