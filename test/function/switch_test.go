package function

import (
	"fmt"
	"go/types"
	"testing"
)

func TestSwitchFunc(t *testing.T) {
	var value byte
	fmt.Println("请输入a,b,c,d,e,f")
	fmt.Scanf("%c", &value)

	switch value {
	case 'a':
		fmt.Println("aaaaaa")
	case 'b':
		fmt.Println("b")
	default:
		fmt.Println("输入错误")
	}
}

var x interface{}

func TestSwitchInterface(t *testing.T) {
	var y = 10.0
	x = y
	switch i := x.(type) {
	case types.Nil:
		fmt.Printf("x的类型为：%T", i)
	case float64:
		fmt.Printf("x的类型是float64")
	case int:
		fmt.Printf("x的类型是int")
	case func(int2 int) float64:
		fmt.Println("x 是func(int)类型")
	default:
		fmt.Println("未知性")
	}
}
