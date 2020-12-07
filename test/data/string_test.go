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

	// 字符串转byte
	bytes := []byte("example test")
	fmt.Printf("bytes=%v\n", bytes)

	//byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	//十进制转换
	str = strconv.FormatInt(123, 2)
	fmt.Println("十进制123对应的二进制是：", str)

	str = strconv.FormatInt(123, 16)
	fmt.Println("十进制123对应的16进制是：", str)
}
