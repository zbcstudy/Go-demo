package main

import (
	"fmt"
	"reflect"
)

func ReflectTest(i interface{}) {
	// 通过反射获取变量的type reflect.Type
	rType := reflect.TypeOf(i)
	fmt.Println("rType:", rType)

	rVal := reflect.ValueOf(i)
	// rVal + 1 error
	i2 := rVal.Int() + 1
	fmt.Printf("rVal=%v,type=%T,i2=%d\n", rVal, rVal, i2)

	// 将rVal转成interface
	iv := rVal.Interface()
	num2 := iv.(int)
	fmt.Printf("num2=%d\n", num2)
}

func ReflectStruct(i interface{}) {
	// 通过反射获取变量的type reflect.Type
	rType := reflect.TypeOf(i)
	fmt.Println("rType:", rType)

	rVal := reflect.ValueOf(i)
	rv := rVal.Interface()
	// rv.Name 无法获取到对应的字段 类型能够获取到是在运行时 编译器无法识别 必须通过断言
	fmt.Printf("rVal=%v,type=%T\n", rVal, rv)
	student, ok := rv.(Student)
	if ok {
		fmt.Println("stu:", student)
	}
}

func main() {
	var name = 100
	ReflectTest(name)

	fmt.Println("***************")
	// 结构体反射
	stu := Student{
		Name: "zbc",
		Age:  89,
	}
	ReflectStruct(stu)
}

type Student struct {
	Name string
	Age  int
}
