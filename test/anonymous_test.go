package test

import (
	"fmt"
	"testing"
)

// 匿名函数测试
func TestAnonymous(t *testing.T) {
	res := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(res)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Printf("a的类型未%T\n", a)
	fmt.Println(a(10, 5))
}
