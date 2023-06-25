package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 结构体的定义方式
func TestStructDefine(t *testing.T) {
	// 1
	var person Person
	marshal, _ := json.Marshal(person)
	fmt.Println(string(marshal))
	person.Name = "name"

	// 2
	person2 := Person{
		Name: "name",
		Age:  18,
	}
	fmt.Println(person2)

	// 3
	var person3 = new(Person)
	// 下面两种赋值方式等价
	person3.Name = "1000"
	(*person3).Age = 38
	fmt.Printf("person3的类型%T\n", person3)
	fmt.Println(person3.Name)

	// 4 指针类型
	var person4 = &Person{
		Name: "new person",
		Age:  19,
	}
	fmt.Println(person4.Age)
}
