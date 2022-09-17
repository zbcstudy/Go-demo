package common

import (
	"testing"
)

// 斐波那契数列
func fnb(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		i := fnb(n-1) + fnb(n-2)
		return i
	}

}

func TestFnb(t *testing.T) {
	println(fnb(10))
}
