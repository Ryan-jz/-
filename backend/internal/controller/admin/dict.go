// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// DictController 字典管理控制器
var Dict = cDict{}

type cDict struct{}

// List 获取字典列表
func (c *cDict) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现字典列表查询
	return nil, nil
}
