package main

import (
	"fmt"
	"math"
)

// 可直接运行
// main包下的main方法 会直接运行
// go run hello.go
// go build hello.go 会得到一个可执行文件
func main() {
	println("hello go")
	fmt.Println(math.Pi)
}
