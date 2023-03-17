package test_hard

import (
	"fmt"
	"testing"
)

// for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
// 所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
func TestRange(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for index, val := range slice {
		fmt.Println("val->", val)
		fmt.Println("&val->", &val)
		m[index] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", v)
		fmt.Println(k, "->", *v)
		fmt.Println(k, "->", &v)
		fmt.Println("--------------")
	}
}

func TestRange2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for index, val := range slice {
		value := val
		fmt.Println("&value->", value)
		m[index] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", v)
		fmt.Println(k, "->", *v)
		fmt.Println(k, "->", &v)
		fmt.Println("--------------")
	}
}
