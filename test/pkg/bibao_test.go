package pkg

import (
	"strings"
	"testing"
)

// 闭包测试
func AddUpper() func(int) int {
	n := 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func TestClose(t *testing.T) {
	a := AddUpper()
	println(a(1)) // 11
	println(a(2)) // 13
	println(a(3)) // 16

	str1 := makeSuffix("**")
	println(str1("test"))
}
