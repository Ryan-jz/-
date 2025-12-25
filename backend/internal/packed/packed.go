// Package packed 用于打包静态资源
// 将配置文件等资源打包到二进制文件中
package packed

import (
	_ "github.com/gogf/gf/v2/os/gres"
)

func init() {
	// 如果需要打包资源，使用 gf pack 命令生成资源文件
	// 然后在这里使用 gres.Add() 添加打包的资源
	// 目前不需要打包资源，保持空实现
}
