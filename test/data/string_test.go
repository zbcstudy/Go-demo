package data

import (
	"strconv"
	"testing"
)

func TestConv(t *testing.T) {
	//数字转字符串
	s := strconv.Itoa(10)
	t.Log("str:", s)
	if num, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + num)
	}
}
