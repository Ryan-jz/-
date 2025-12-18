// Package packed 用于打包静态资源
// 将配置文件等资源打包到二进制文件中
package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	// 如果需要打包资源，在这里添加
	// 使用 gf pack 命令生成资源文件
	if err := gres.Add(""); err != nil {
		panic(err)
	}
}
