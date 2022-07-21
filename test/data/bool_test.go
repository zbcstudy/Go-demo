package data

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestBool(t *testing.T) {
	var b = false
	fmt.Println("b=", b)
	b = true
	fmt.Println("bool 类型占用的空间为：", unsafe.Sizeof(b))

}
