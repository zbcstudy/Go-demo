package data

import (
	"fmt"
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

func TestLength(t *testing.T) {
	str := "hello好"
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("字符为：%c\n", str[i])
	}
	sli := []rune(str)
	for i := 0; i < len(sli); i++ {
		fmt.Printf("字符为：%c\n", sli[i])
	}
}
