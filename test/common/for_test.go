package common

import (
	"fmt"
	"testing"
)

func TestForFunc(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}
}

func TestForFunc2(t *testing.T) {
	i := 0
	for i < 10 {
		fmt.Println("hello world")
		i++
	}
}

func TestForString(t *testing.T) {
	var str string = "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c---%v\n", str[i], str[i])
	}
}

func TestForString2(t *testing.T) {
	str := "hello,world"
	for index, val := range str {
		fmt.Printf("index=%d,value=%c\n", index, val)
	}
}

func TestForString3(t *testing.T) {
	str := "hello,北京"
	newStr := []rune(str)
	for i := 0; i < len(newStr); i++ {
		fmt.Printf("%c---%v\n", newStr[i], newStr[i])
	}
}

// rang是按字节进行遍历 推荐使用
func TestForString4(t *testing.T) {
	str := "hello,上海"
	for index, val := range str {
		fmt.Printf("index=%d,value=%c\n", index, val)
	}
}
