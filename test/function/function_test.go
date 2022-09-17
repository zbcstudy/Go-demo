package function

import (
	"fmt"
	"testing"
)

func sum(ops ...int) int {
	var res = 0
	for _, op := range ops {
		res += op
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("last clear resource")
}

// defer 关键字用来最后资源的释放
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	//panic("12error")
}

// 自定义类型myInt 在go中 虽然myInt和int都是 int类型 但是go认为是两个类型
func TestFuncTypeWithInit(t *testing.T) {
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	// num2 = num1 不能直接复制
	num2 = int(num1)
	fmt.Println("num1=", num1, "num2=", num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

// go 函数允许对函数返回值命名
func TestFuncWithReturn(t *testing.T) {
	a, b := getSumAndSub(100, 20)
	fmt.Println("a=", a, "b=", b)
}
