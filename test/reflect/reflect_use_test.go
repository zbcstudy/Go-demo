package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("------start-----")
	fmt.Println(m)
	fmt.Println("-------end------")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func Call(a interface{}) {
	// 获取type类型
	rType := reflect.TypeOf(a)

	rVal := reflect.ValueOf(a)

	kind := rVal.Kind()

	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取结构体中有几个字段
	num := rVal.NumField()
	fmt.Printf("struct has %d field\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("field-%d:的值为%v\n", i, rVal.Field(i))

		// 获取到struct标签 需要通过reflect.Type来获取tag标签的值
		tag := rType.Field(i).Tag.Get("json")
		if tag != "" {
			fmt.Printf("field-%d:tag为：%s\n", i, tag)
		}
	}

	// 方法操作
	numMethod := rVal.NumMethod() //获取方法个数
	fmt.Printf("struct has %d 个方法\n", numMethod)

	// 获取结构体的第2个方法 并进行调用 结构体的方法是按照方法名进行排序的
	rVal.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))

	res := rVal.Method(0).Call(params)
	fmt.Printf("method call res:%v\n", res[0].Int())
}

func TestStruct(t *testing.T) {

	value := reflect.ValueOf(10)
	fmt.Println("value:", value)

	var a = Monster{
		Name:  "黄鼠狼",
		Age:   100,
		Score: 20.3,
	}
	Call(a)
}
