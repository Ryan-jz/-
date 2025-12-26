package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// 用于生成密码哈希的工具
// 使用方法: go run hack/hash_password.go <password>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run hack/hash_password.go <password>")
		fmt.Println("示例: go run hack/hash_password.go admin")
		os.Exit(1)
	}

	password := os.Args[1]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("生成密码哈希失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("加密后的密码: %s\n", string(hashedPassword))
	fmt.Println("\n请将加密后的密码更新到数据库中")
}
