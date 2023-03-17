package data

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

// string 类型转换
func TestStringFmt(t *testing.T) {
	c1 := 10
	var num2 = 23.345
	var val3 bool
	var val4 byte = 'h'

	var str string

	str = fmt.Sprintf("%d", c1)
	fmt.Printf("str type %T,str = %v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T,str = %q\n", str, str)

	str = fmt.Sprintf("%t", val3)
	fmt.Printf("str type %T,str = %q\n", str, str)

	str = fmt.Sprintf("%c", val4)
	fmt.Printf("str type %T,str = %q\n", str, str)

	fmt.Println("************************************")

}

func TestConv(t *testing.T) {
	//数字转字符串
	s := strconv.Itoa(10)
	t.Log("str:", s)
	if num, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + num)
	}

	parseInt, err := strconv.ParseInt("hello", 10, 0)
	t.Log(err)
	t.Log("parseInt:", parseInt)

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

// 字符串字面量 该变量的内存地址 底层字节切片
func TestStringStruct(t *testing.T) {
	a := "hello"
	fmt.Println(a, &a, (*reflect.StringHeader)(unsafe.Pointer(&a)))

	a = "world"
	fmt.Println(a, &a, (*reflect.StringHeader)(unsafe.Pointer(&a)))

	b := "hello"
	fmt.Println(b, &b, (*reflect.StringHeader)(unsafe.Pointer(&b)))

	fmt.Println(unsafe.Pointer(&b))
}
