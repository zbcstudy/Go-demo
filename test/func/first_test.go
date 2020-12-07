package _func

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Log("first log test")
}

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		fmt.Println(" ", b)
		temp := a
		a = b
		b = temp + a
	}

}

//指针测试
//指针变量存的是一个地址
func TestPointer(t *testing.T) {
	var i = 10
	//ptr是一个指针变量
	var ptr = &i
	fmt.Println(ptr)
	//获取指针变量指向的值
	fmt.Println(*ptr)
	fmt.Println(&i)
}
