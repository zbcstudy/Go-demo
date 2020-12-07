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

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}

func TestNewFunc(t *testing.T) {
	num := new(int)
	*num = 100
	fmt.Printf("num的类型为：%T,num的内存地址为：%v,num的值为：%v", *num, &num, *num)
}
