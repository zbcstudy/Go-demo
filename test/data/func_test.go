package data

import (
	"fmt"
	"testing"
)

// 函数本身就是一种数据类型 可以被当做形参 并且调用
func sum(num1 int, num2 int) int {
	return num1 + num2
}

func sumFunc(funcVar func(int, int) int, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

func TestFuncAsParam(t *testing.T) {
	res := sumFunc(sum, 20, 30)
	fmt.Println(res)
}
