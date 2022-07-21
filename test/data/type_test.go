package data

import (
	"fmt"
	"math"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
	t.Log(math.MaxInt8)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}

func TestNewFunc(t *testing.T) {
	num := new(int)
	println(num)
	*num = 100
	println(&num)
	fmt.Printf("num的类型为：%T,num的内存地址为：%v,num的值为：%v", *num, &num, *num)
}
