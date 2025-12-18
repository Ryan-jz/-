// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// RoleController 角色管理控制器
var Role = cRole{}

type cRole struct{}

// List 获取角色列表
func (c *cRole) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现角色列表查询
	return nil, nil
}
