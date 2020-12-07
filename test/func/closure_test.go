package _func

import (
	"fmt"
	"testing"
)

//闭包测试
func AddUpper() func(int) int {
	var n = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func TestClosure(t *testing.T) {
	f := AddUpper()
	fmt.Println(f(2)) // 12
	fmt.Println(f(1)) // 13
	fmt.Println(f(3)) // 16

}
