package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net"
	"time"
)

type Person struct {
	Name  string
	Age   int
	score [5]float64
	ptr   *int  // 指针
	slice []int // 切片
	info  map[string]string
}

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	R    string `json:"r"`
}

func (m Monster) SetAge(age int) {
	m.Age = age
	fmt.Println("修改后的值：", m.Age)
}

func (m *Monster) setAgeWithPoint(age int) {
	m.Age = age
	fmt.Println("修改后的值：", m.Age)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// 默认值
	defaultVal := Person{
		Name:  "",
		Age:   0,
		score: [5]float64{},
		ptr:   nil,
		slice: nil,
		info:  nil,
	}
	spew.Dump(defaultVal) // 打印默认值
	var person Person
	println("***************")
	person.slice = make([]int, 5) // 使用之前必须make
	fmt.Println(person.slice)

	person.Name = "name"
	person.Age = 18

	// 结构体默认采用值copy
	person2 := person
	person2.Name = "new name"

	fmt.Println("person：", person)
	fmt.Println("person2：", person2)

	var config net.ListenConfig
	config.KeepAlive = time.Hour * 2
	configByte, _ := json.Marshal(config)
	fmt.Println(configByte)

	listenConfig := net.ListenConfig{
		KeepAlive: 1,
	}
	lcb, _ := json.Marshal(listenConfig)
	fmt.Println(lcb)
}
