// Package main 是应用程序的入口点
// 负责初始化和启动 GoFrame 应用服务器
package main

import (
	_ "gf-admin/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-admin/internal/cmd"
)

// main 主函数，应用程序入口
// 创建上下文并运行命令行应用
func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
