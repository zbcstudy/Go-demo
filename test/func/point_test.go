package _func

import (
	"fmt"
	"testing"
)

// 指针测试
func test01(num int) {
	num = num + 10
	fmt.Println("test01 num=", num)
}

func test02(num *int) {
	*num = *num + 10
	fmt.Println("test02 num=", *num)
}

func TestPoint(t *testing.T) {
	// 方法值传递不会改变方法外的值
	num1 := 10
	test01(num1)
	fmt.Println("num1的值为num1=", num1)

	num2 := 10
	test02(&num2)
	fmt.Println("num2的值为num2=", num2)
}
