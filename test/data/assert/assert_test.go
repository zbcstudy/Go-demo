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
	var temp = Point{}
	fmt.Println("temp: ", temp)
	var point = Point{1, 2}

	a = point // ok

	var b Point
	// b = a 不能直接赋值
	b = a.(Point) // 可以 使用断言的方式
	fmt.Println(b)

	var x interface{}
	var z float32 = 1.1
	x = z
	y := x
	fmt.Printf("y的类型：%T,y的值为：%v\n", y, y)

	e, ok := x.(float64)
	if ok {
		fmt.Println("convert success")
		fmt.Printf("e的类型：%T,e的值为：%v\n", e, e)
	} else {
		fmt.Println("convert fail")
	}

	if f, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("f的类型：%T,f的值为：%v\n", f, f)
	} else {
		fmt.Println("convert fail")
	}
}
