package common

import (
	"fmt"
	"testing"
)

// 全局匿名函数
var (
	func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// 匿名函数测试
func TestNiMingFunc(t *testing.T) {
	// 匿名哈数直接调用
	rest := func(num1 int, num2 int) int {
		return num2 + num1
	}(10, 20)
	println(rest)

	result := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Printf("result的数据类型：%T \n", result)
	println(result(20, 10))

	println(func1(10, 20))
}
