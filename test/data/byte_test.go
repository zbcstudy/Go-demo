package data

import (
	"fmt"
	"testing"
)

func TestByteValue(t *testing.T) {
	var a byte = 'a'
	var b byte = '0'
	// 当我们直接输出byte 打印的是对应的码值
	fmt.Println("a=", a)
	fmt.Println("0=", b)
	fmt.Printf("c=%c", a)
	fmt.Println("****************************")

	var c3 int = '北'
	fmt.Printf("c3=%c,c3对应的码值为%d", c3, c3)

}
