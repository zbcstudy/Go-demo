package assert

import (
	"fmt"
	"testing"
)

type Point struct {
	x int
	y int
}

// 如何使用类型断言
func TestAssert(t *testing.T) {
	var a interface{}
	var temp Point = Point{}
	fmt.Println("temp: ", temp)
	var point Point = Point{1, 2}

	a = point // ok

	var b Point
	// b = a 不能直接赋值
	b = a.(Point) // 可以 使用断言的方式
	fmt.Println(b)
}
