package common

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("integer:", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("integer:", s)
		return
	}
	fmt.Println("unknown type")
}

func DoSomethingWithSwitch(p interface{}) {
	switch t := p.(type) {
	case int:
		fmt.Println("integer:", t)
	case string:
		fmt.Println("string:", t)
	default:
		fmt.Println("unknown type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("some")
	DoSomethingWithSwitch(11)
	DoSomethingWithSwitch("33")
	DoSomethingWithSwitch(5.5)
}
