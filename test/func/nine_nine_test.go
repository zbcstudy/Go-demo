package _func

import (
	"fmt"
	"testing"
)

// 99乘法表
func TestNineFunc(t *testing.T) {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println("")
	}
}
