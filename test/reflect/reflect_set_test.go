package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectSetValue(t *testing.T) {
	var num = 10
	ReflectNum(&num)
	fmt.Println("rVal:", num)
}

func ReflectNum(i interface{}) {
	rVal := reflect.ValueOf(i)

	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	// reflect.Value.SetInt using unaddressable value [recovered]
	// rVal.SetInt(11) // 直接报错

	rVal.Elem().SetInt(100) // 这种方式才能设置成功
}
