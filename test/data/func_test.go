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

func TestFuncAsParam2(t *testing.T) {
	a := sum
	fmt.Printf("a的类型：%T,sum的类型：%T \n", a, sum)

	res := a(1, 3)
	fmt.Println(res)
}

// 可变参数使用
func sumMulti(num int, args ...int) int {
	for i := 0; i < len(args); i++ {
		num = num + args[i]
	}
	return num
}

func TestFuncMultiParam(t *testing.T) {
	fmt.Println(sumMulti(1, 2, 3, 4, 5))
}
