package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取所有的命令行参数
	args := os.Args
	for i, v := range args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	arg0 := os.Args[0]
	base := filepath.Base(arg0)
	fmt.Println("base:", base)
}
