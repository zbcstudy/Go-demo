package data

import "testing"

// 指针
// 基本数据类型 变量存的就是值，也叫值类型
// 获取变量的地址 使用&
// 指针类型 变量存的是一个地址 这个地址指向的空间存的才是值
// 获取指针类型所指向的值，使用 *

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, &a, aPtr)
	t.Log(a, &a, *aPtr)
}

func TestPoint2(t *testing.T) {
	var num int
	var ptr *int = &num
	t.Log(ptr, &ptr, *ptr)

}
