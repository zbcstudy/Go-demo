package _interface

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type Student struct {
	Name string
}

func (stu Student) TestC() {
	fmt.Println("test03")
}

func (stu Student) Test02() {
	fmt.Println("test02")
}

func (stu Student) Test01() {
	fmt.Println("test01")
}

func TestInterface(t *testing.T) {

	var stu Student
	//var ic InterfaceC = stu // 会报错
	spew.Dump(stu)
	stu.TestC()

	var t1 T = stu
	fmt.Println(t1)

	// t2可以接受任何值 空接口没有类型限制
	var t2 interface{} = stu
	fmt.Println(t2)
}
