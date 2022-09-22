package main

import (
	"fmt"
	"testing"
)

func TestStructMethod(t *testing.T) {
	monster := Monster{
		Name: "zbc",
		Age:  10,
		R:    "r-t",
	}
	fmt.Printf("调用方法之前的地址%p,age=%v\n", &monster, monster.Age)
	monster.SetAge(20)
	fmt.Printf("调用方法之后的地址%p,age=%v\n", &monster, monster.Age)
}

func TestStructPointMethod(t *testing.T) {
	monster := &Monster{
		Name: "zbc",
		Age:  10,
		R:    "r-t",
	}
	fmt.Printf("调用方法之前的地址%p,age=%v\n", &monster, monster.Age)
	monster.setAgeWithPoint(20)
	fmt.Printf("调用方法之后的地址%p,age=%v\n", &monster, monster.Age)
}

func TestStructInvoke(t *testing.T) {
	c := Circle{4}
	res := c.area()
	fmt.Println(res)
}
