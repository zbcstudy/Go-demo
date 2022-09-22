package main

import (
	"fmt"
	"testing"
)

type A struct {
	num int
}

type B struct {
	num int
}

// 结构体之间进行转换 必须 结构体的字段名称、类型、个数完全一致才可以
func TestStructConvert(t *testing.T) {
	a := A{1}
	b := B(a)
	fmt.Println(a, b)

}
