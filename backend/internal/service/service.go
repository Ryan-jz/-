// Package service 提供业务逻辑服务
// 服务层负责处理业务逻辑，协调 DAO 层和控制器层
package service

import (
	"context"
)

var (
	localProduct IProduct
)

// Init 初始化服务
// 在应用启动时调用，用于初始化各种服务
func Init(ctx context.Context) {
	// 这里可以添加服务初始化逻辑
	// 例如：初始化缓存、连接池等
}

// Product 获取产品服务实例
func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

// RegisterProduct 注册产品服务实现
func RegisterProduct(i IProduct) {
	localProduct = i
}
