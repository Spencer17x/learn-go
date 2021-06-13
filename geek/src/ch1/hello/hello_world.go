package main

import (
	"fmt"
	"os"
)

// 应用程序入口
// 1. 必须是 main 包：package main
// 2. 必须是 main 方法：func main()
// 3. 文件名不一定是 main.go

// 退出返回值
// 1. Go 中 main 函数不支持任何返回值
// 2. 通过 os.Exit 来返回状态

// 获取命令行参数
// 1. main 函数不支持传入参数
// 2. 在程序中直接通过 os.Args 获取命令行参数
func main() {
	if len(os.Args) > 1 {
		fmt.Println("os.Args[1]", os.Args[1])
	}
	fmt.Println("Hello World")
	os.Exit(0) // 正常退出
	// os.Exit(-1) // 异常退出
}
