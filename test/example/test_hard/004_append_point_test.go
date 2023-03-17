package test_hard

import (
	"fmt"
	"testing"
)

func TestAppendPoint(t *testing.T) {
	s := new([]int)
	// append不能对指针进行值添加
	//s = append(s, 1)
	fmt.Println(s)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// 追加切片的时候需要使用下面方式
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
