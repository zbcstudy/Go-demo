package pkg

import (
	"fmt"
	"testing"
)

// 使用切片存储斐波那契数列
func fnb(count int) []uint64 {
	// 声明一个切片
	fnbSlice := make([]uint64, count)
	// 第一个和第二个数为1
	fnbSlice[0] = 1
	fnbSlice[1] = 1
	for i := 2; i < count; i++ {
		fnbSlice[i] = fnbSlice[i-1] + fnbSlice[i-2]
	}

	return fnbSlice
}

func TestFnb(t *testing.T) {
	fmt.Println(fnb(10))
}
