package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取所有的命令行参数
	args := os.Args
	for i, v := range args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
