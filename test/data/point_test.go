package data

import (
	"fmt"
	"testing"
)

// 指针
// 基本数据类型 变量存的就是值，也叫值类型
// 获取变量的地址 使用&
// 指针类型 变量存的是一个地址 这个地址指向的空间存的才是值
// 获取指针类型所指向的值，使用 *

// 值类型 ： int float bool string 数组 struct  数据在栈区
// 引用类型： 指针 slice切片 map 管道chan interface 数据在堆区
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, &a, aPtr)
	t.Log(a, &a, *aPtr)
}

func TestPoint2(t *testing.T) {
	var num int
	var ptr *int = &num // 指针变量
	fmt.Println("num的地址:", &num)
	fmt.Println("ptr本身的地址：", &ptr)
	// ptr存储的值是num的内存地址
	fmt.Println("ptr存储的值为：", ptr)
	t.Log(ptr, &ptr, *ptr)
}

func TestPoint3(t *testing.T) {
	var num int = 9
	fmt.Printf("num的address：%v\n", &num)

	var ptr *int = &num

	*ptr = 10 // 这里修改会关联到 num

	fmt.Println("num:", num)
}

func TestPoint4(t *testing.T) {
	num1 := 100
	fmt.Printf("num1的类型%T,num1的值%v,num1的地址%v\n", num1, num1, &num1)
	//num1的类型int,num1的值100,num1的地址0xc00000a368

	num2 := new(int)
	fmt.Printf("num2的类型%T,num2的值%v,num2的地址%v\n", num2, num2, &num2)
	//	num2的类型*int,num2的值0xc00000a370,num2的地址0xc000006038

	println("num2指针对应的值为: ", *num2)

	*num2 = 100
	println("num2指针对应的值为: ", *num2)

}
