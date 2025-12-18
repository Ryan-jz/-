// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// DeptController 部门管理控制器
var Dept = cDept{}

type cDept struct{}

// List 获取部门列表
func (c *cDept) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现部门列表查询
	return nil, nil
}
