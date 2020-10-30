package data

import (
	"fmt"
	"testing"
)

func recursion(num int) int {
	if num > 2 {
		num--
		recursion(num)
	}
	fmt.Println(num)
	return num
}

func TestRecursion(t *testing.T) {
	recursion(5)
}

func TestFuncType(t *testing.T) {
	a := recursion
	fmt.Printf("t的数据类型为：%T,recursion的数据类型为：%T\n", a, recursion)
	res := a(5)
	fmt.Println(res)

}
