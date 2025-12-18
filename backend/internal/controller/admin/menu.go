// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// MenuController 菜单管理控制器
var Menu = cMenu{}

type cMenu struct{}

// List 获取菜单列表
func (c *cMenu) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现菜单列表查询
	return nil, nil
}
