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
	var i int
	println(i)
	println("***************")
	println("hello go")
	fmt.Println(math.Pi)

	fmt.Println("姓名\t百度一下")
	fmt.Println("姓名\n百度一下")
}
