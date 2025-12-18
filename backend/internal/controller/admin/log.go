// Package admin 提供管理后台接口
package admin

import (
	"context"
)

// LogController 日志管理控制器
var Log = cLog{}

type cLog struct{}

// List 获取日志列表
func (c *cLog) List(ctx context.Context, req interface{}) (res interface{}, err error) {
	// TODO: 实现日志列表查询
	return nil, nil
}
